package elasticsearch

import (
	"encoding/json"
	"fmt"
	"github.com/mschoch/elastigo/api"
	"github.com/mschoch/elastigo/core"
	"github.com/mschoch/tuq/datasources"
	"github.com/mschoch/tuq/parser"
	"github.com/mschoch/tuq/planner"
	"log"
	"math"
	"net/url"
	"strings"
)

const StatsPrefix = "__stats__"
const defaultBatchSize = 10000
const defaultMaxAggregate = 1000000
const defaultDebug = false

type ElasticSearchDataSource struct {
	Name              string
	As                string
	OutputChannel     planner.DocumentChannel
	cancelled         bool
	query             *QueryDocument
	filter            *Filter
	facets            *Facets
	limit             int
	offset            int
	sort              []map[string]interface{}
	sortlist          parser.SortList
	docBodyDataSource planner.DataSource
	indexName         string
	filterExpression  parser.Expression
	batchSize         int
	maxAggregate      int
	debug             bool
	defaultFilter     *Filter
}

func init() {
	datasources.RegisterDataSourceImpl("elasticsearch", NewElasticSearchDataSource)
}

func NewElasticSearchDataSource(config map[string]interface{}) planner.DataSource {

	result := &ElasticSearchDataSource{
		OutputChannel: make(planner.DocumentChannel),
		cancelled:     false,
		limit:         -1,
		offset:        0,
		indexName:     config["index"].(string)}

	result.SetURL(config["elasticsearch"].(string))

	if config["doc_body"] != nil {
		result.docBodyDataSource = datasources.NewDataSourceWithName(config["doc_body"].(string))
	}

	if config["batch_size"] != nil {
		result.batchSize = config["batch_size"].(int)
	} else {
		result.batchSize = defaultBatchSize
	}

	if config["max_aggregate"] != nil {
		result.maxAggregate = config["max_aggregate"].(int)
	} else {
		result.maxAggregate = defaultMaxAggregate
	}

	if config["debug"] != nil {
		result.debug = config["debug"].(bool)
	} else {
		result.debug = defaultDebug
	}

	if config["default_filter"] != nil {
		result.defaultFilter = ConvertToFilterRecursive(config["default_filter"].(map[string]interface{}))
	}

	return result
}

// PlanPipelineComponent interface

func (ds *ElasticSearchDataSource) SetSource(source planner.PlanPipelineComponent) {
	log.Fatalf("SetSource called on DataSource")
}

func (ds *ElasticSearchDataSource) GetSource() planner.PlanPipelineComponent {
	return nil
}

func (ds *ElasticSearchDataSource) GetDocumentChannel() planner.DocumentChannel {
	return ds.OutputChannel
}

func (ds *ElasticSearchDataSource) Run() {
	defer close(ds.OutputChannel)

	// build the query
	ds.buildQuery()

	if ds.facets != nil {
		ds.runAggregateQuery()
	} else {
		ds.runQuery()
	}

}

func (ds *ElasticSearchDataSource) runAggregateQuery() {

	// send the query (first batch)
	searchresponse, err := core.Search(true, ds.indexName, "", ds.query, "1m")
	if err != nil {
		log.Printf("Error running ES aggregate query: %v", err)
		return
	}

	// extract results from the response
	docs, err := RowsFromFacetSearchResults(searchresponse, ds.As)
	if err != nil {
		log.Printf("Error parsing ES response: %v", err)
		return
	}

	// send the results
	for _, doc := range docs {
		ds.OutputChannel <- doc
	}

}

func (ds *ElasticSearchDataSource) runQuery() {
	// initial state
	resultsSeen := 0

	// send the query (first batch)
	searchresponse, err := core.Search(true, ds.indexName, "", ds.query, "1m")
	if err != nil {
		log.Printf("Error running ES query: %v", err)
		return
	}

	// extract results from the response
	docs, err := ds.DocsFromSearchResults(searchresponse)
	if err != nil {
		log.Printf("Error parsing ES response: %v", err)
		return
	}

	// send this batch of results
	for _, doc := range docs {
		ds.OutputChannel <- doc
	}

	// updated state after initial response
	resultsSeen += len(searchresponse.Hits.Hits)
	scrollId := searchresponse.ScrollId
	// total results adjusted for any offset
	total := searchresponse.Hits.Total - ds.offset
	if ds.limit != -1 && total > ds.limit {
		// if user specified a limit that is lower than the total
		// we must stop at that limit
		total = ds.limit
	}

	for resultsSeen < total {
		// request the next batch
		searchresponse, err := core.Scroll(true, scrollId, "1m")
		if err != nil {
			log.Printf("Error running ES scroll: %v", err)
			return
		}

		// extract results from the response
		docs, err := ds.DocsFromSearchResults(searchresponse)
		if err != nil {
			log.Printf("Error parsing ES response: %v", err)
			return
		}

		// send this batch of results
		for _, doc := range docs {
			if resultsSeen < total {
				ds.OutputChannel <- doc
			}
			resultsSeen += 1
		}

		scrollId = searchresponse.ScrollId
	}
}

func (ds *ElasticSearchDataSource) Explain() {
	defer close(ds.OutputChannel)

	// build the query
	ds.buildQuery()

	thisStep := map[string]interface{}{
		"_type":     "FROM",
		"impl":      "ElasticSearch",
		"index":     ds.indexName,
		"name":      ds.Name,
		"as":        ds.As,
		"query":     ds.query,
		"limit":     ds.limit,
		"offset":    ds.offset,
		"cost":      ds.Cost(),
		"totalCost": ds.TotalCost()}

	ds.OutputChannel <- thisStep
}

func (ds *ElasticSearchDataSource) Cancel() {
	ds.cancelled = true
}

// DataSource Interface

func (ds *ElasticSearchDataSource) SetName(name string) {
	ds.Name = name
}

func (ds *ElasticSearchDataSource) SetAs(as string) {
	ds.As = as
}

func (ds *ElasticSearchDataSource) GetAs() string {
	return ds.As
}

func (ds *ElasticSearchDataSource) SetFilter(filter parser.Expression) error {
	f, err := ds.BuildElasticSearchFilterRecursive(filter)
	if err != nil {
		return err
	} else {
		ds.filter = f
		ds.filterExpression = filter
	}
	return nil
}

func (ds *ElasticSearchDataSource) GetFilter() parser.Expression {
	return ds.filterExpression
}

func (ds *ElasticSearchDataSource) SetOrderBy(sortlist parser.SortList) error {
	if ds.facets != nil {
		return fmt.Errorf("ElasticSearch DataSource does not support ordering with aggregate query")
	}
	d, err := ds.BuildElasticSearchOrderBy(sortlist)
	if err != nil {
		return err
	} else {
		ds.sortlist = sortlist
		ds.sort = d
	}
	return nil
}

func (ds *ElasticSearchDataSource) GetOrderBy() parser.SortList {
	return ds.sortlist
}

func (ds *ElasticSearchDataSource) SetLimit(expr parser.Expression) error {
	if ds.facets != nil {
		// this sort of works, but its usefulness is questionable
		// normally offset/limit only make sense with a sort order
		// and our ability to sort aggregates in ES is limited
		// so im disabling it for now
		return fmt.Errorf("ElasticSearch DataSource does not support limit with aggregate query")
	}
	switch expr := expr.(type) {
	case *parser.IntegerLiteral:
		ds.limit = expr.Val
	default:
		return fmt.Errorf("ElasticSearch DataSource only supports limiting by integer literals")
	}
	return nil
}

func (ds *ElasticSearchDataSource) SetOffset(expr parser.Expression) error {
	if ds.facets != nil {
		return fmt.Errorf("ElasticSearch DataSource does not support offset with aggregate query")
	}
	switch expr := expr.(type) {
	case *parser.IntegerLiteral:
		ds.offset = expr.Val
	default:
		return fmt.Errorf("ElasticSearch DataSource only supports offsetting by integer literals")
	}
	return nil
}

func (ds *ElasticSearchDataSource) SetGroupByWithStatsFields(groupby parser.ExpressionList, stats_fields []string) error {
	if len(groupby) != 1 {
		return fmt.Errorf("ElasticSearch DataSource can only group by a single expression")
	}

	ds.facets = EmptyFacets()

	if len(groupby) == 1 {

		switch val := groupby[0].(type) {
		case *parser.Property:

			// check the property
			if !val.IsReferencingDataSource(ds.As) {
				return fmt.Errorf("ElasticSearch can only group results by property in this datasource")
			} else {
				// strip off the data sourc portion of the property
				val = val.Tail()
			}

			facet := NewTermsFacet(val.Symbol, ds.maxAggregate)
			(*ds.facets)[val.Symbol] = facet

			for _, stat_field := range stats_fields {
				if !strings.HasPrefix(stat_field, ds.As+".") {
					return fmt.Errorf("ElasticSearch can only gather stats on properties in this datasource")
				} else {
					stat_field = stat_field[len(ds.As)+1:]
				}
				if stat_field != val.Symbol {
					facet := NewTermsStatsFacet(val.Symbol, stat_field, ds.maxAggregate)
					(*ds.facets)[StatsPrefix+stat_field] = facet
				}
			}
		case *parser.BoolLiteral:
			if val.Val == true {
				// group by true means, put everything in one group
				for _, stat_field := range stats_fields {
					if !strings.HasPrefix(stat_field, ds.As+".") {
						return fmt.Errorf("ElasticSearch can only gather stats on properties in this datasource")
					} else {
						stat_field = stat_field[len(ds.As)+1:]
					}
					facet := NewStatisticalFacet(stat_field, ds.maxAggregate)
					(*ds.facets)[StatsPrefix+stat_field] = facet
				}
			} else {
				fmt.Errorf("ElasticSearch only supports boolean expression true")
			}
		default:
			return fmt.Errorf("ElasticSearch only supporter property expressions in group by")
		}
	}

	return nil
}

func (ds *ElasticSearchDataSource) SetHaving(having parser.Expression) error {
	return fmt.Errorf("ElasticSearch DataSource does not yet support having")
}

// ES Specific

func (ds *ElasticSearchDataSource) SetURL(urlString string) error {
	esURL, err := url.Parse(urlString)
	if err != nil {
		return err
	} else {
		api.Protocol = esURL.Scheme
		colonIndex := strings.Index(esURL.Host, ":")
		if colonIndex < 0 {
			api.Domain = esURL.Host
			api.Port = "9200"
		} else {
			api.Domain = esURL.Host[0:colonIndex]
			api.Port = esURL.Host[colonIndex+1:]
		}

	}
	return nil
}

// Support

func (ds *ElasticSearchDataSource) buildQuery() {
	ds.query = NewDefaultQuery(ds.batchSize)

	//combine filter with default filter if there was one
	if ds.defaultFilter != nil {
		if ds.filter != nil {
			ds.filter = NewAndFilter(ds.defaultFilter, ds.filter)
		} else {
			ds.filter = ds.defaultFilter
		}
	}

	// add filter to non-aggregate query
	if ds.facets == nil {
		if ds.filter != nil {
			(*ds.query)["filter"] = ds.filter
		}
		if ds.sort != nil {
			(*ds.query)["sort"] = ds.sort
		}
		if ds.limit != -1 && ds.limit < ds.batchSize {
			(*ds.query)["size"] = ds.limit
		}
		if ds.offset != 0 {
			(*ds.query)["from"] = ds.offset
		}
	} else {

		if ds.filter != nil {
			// walk through all the facets and set the filter as a facet filter
			for _, facet := range *ds.facets {
				(*facet.(*Facet)).SetFacetFilter(ds.filter)
			}
		}
		if ds.limit != -1 {
			// walk through all the facets and set the size
			for _, facet := range *ds.facets {
				(*facet.(*Facet)).SetFacetSize(ds.limit)
			}
		}

		(*ds.query)["facets"] = ds.facets
		(*ds.query)["size"] = 0
	}
}

func (ds *ElasticSearchDataSource) DocsFromSearchResults(sr core.SearchResult) ([]planner.Document, error) {

	var docBodies []planner.Document
	var err error
	if ds.docBodyDataSource != nil {

		docIds := make([]string, 0, len(sr.Hits.Hits))
		for _, val := range sr.Hits.Hits {
			docIds = append(docIds, val.Id)
		}

		docBodies, err = ds.docBodyDataSource.DocsFromIds(docIds)
		if err != nil {
			return nil, err
		}

	}

	result := make([]planner.Document, 0, len(sr.Hits.Hits))
	for i, val := range sr.Hits.Hits {
		var source planner.Document
		// marshall into json
		jsonErr := json.Unmarshal(val.Source, &source)
		if jsonErr != nil {
			return nil, jsonErr
		}

		// add bodies if present
		if docBodies != nil {
			source["doc"] = docBodies[i]
		}

		source = datasources.WrapDocWithDatasourceAs(ds.As, source)

		result = append(result, source)
	}

	return result, nil
}

func RowsFromFacetSearchResults(sr core.SearchResult, as string) ([]planner.Document, error) {
	result := make([]planner.Document, 0)

	var facetResults map[string]interface{}

	// unmarshall from json
	jsonErr := json.Unmarshal(sr.Facets, &facetResults)
	if jsonErr != nil {
		return nil, jsonErr
	}

	// look for the group by field (if there was one)
	group_by := ""
	for stat_facet_name, _ := range facetResults {
		if !strings.HasPrefix(stat_facet_name, StatsPrefix) {
			group_by = stat_facet_name
			break
		}
	}

	if group_by == "" {
		// if there was no group by, then there is only 1 row

		row := make(map[string]interface{})

		for stat_facet_name, stat_facet_details := range facetResults {

			if strings.HasPrefix(stat_facet_name, StatsPrefix) {
				stat_field := stat_facet_name[len(StatsPrefix):]
				stat_facet_details_map := stat_facet_details.(map[string]interface{})
				planner.SetDocumentProperty(row, parser.NewProperty(fmt.Sprintf("__func__.count.%v.%v", as, stat_field)), stat_facet_details_map["count"])
				planner.SetDocumentProperty(row, parser.NewProperty(fmt.Sprintf("__func__.min.%v.%v", as, stat_field)), stat_facet_details_map["min"])
				planner.SetDocumentProperty(row, parser.NewProperty(fmt.Sprintf("__func__.max.%v.%v", as, stat_field)), stat_facet_details_map["max"])
				planner.SetDocumentProperty(row, parser.NewProperty(fmt.Sprintf("__func__.avg.%v.%v", as, stat_field)), stat_facet_details_map["mean"])
				planner.SetDocumentProperty(row, parser.NewProperty(fmt.Sprintf("__func__.sum.%v.%v", as, stat_field)), stat_facet_details_map["total"])
			}

		}

		// add this row to the result set
		result = append(result, row)
	} else {

		facet_name := group_by
		facet_details := facetResults[facet_name]

		facet_details_map := facet_details.(map[string]interface{})

		other := facet_details_map["other"]
		if other.(float64) != 0 {
			return nil, fmt.Errorf("Facet results reported %#v \"other\" rows, increate your esMaxAggregate value and try again", other)
		}

		facet_type := facet_details_map["_type"]

		if facet_type == "terms" {
			terms := facet_details_map["terms"]

			for _, term := range terms.([]interface{}) {
				term_map := term.(map[string]interface{})
				row := make(map[string]interface{})
				planner.SetDocumentProperty(row, parser.NewProperty(fmt.Sprintf("%v.%v", as, facet_name)), term_map["term"])
				planner.SetDocumentProperty(row, parser.NewProperty(fmt.Sprintf("__func__.count.%v.%v", as, facet_name)), term_map["count"])

				// now look for any other stats facets with the same term
				// and add those results to this row

				for stat_facet_name, stat_facet_details := range facetResults {

					if strings.HasPrefix(stat_facet_name, StatsPrefix) {
						stat_field := stat_facet_name[len(StatsPrefix):]
						stat_facet_details_map := stat_facet_details.(map[string]interface{})
						stat_terms := stat_facet_details_map["terms"]
						for _, stat_term := range stat_terms.([]interface{}) {
							stat_term_map := stat_term.(map[string]interface{})
							if term_map["term"] == stat_term_map["term"] {
								//this is the term we're looking for
								planner.SetDocumentProperty(row, parser.NewProperty(fmt.Sprintf("__func__.count.%v.%v", as, stat_field)), stat_term_map["count"])
								planner.SetDocumentProperty(row, parser.NewProperty(fmt.Sprintf("__func__.min.%v.%v", as, stat_field)), stat_term_map["min"])
								planner.SetDocumentProperty(row, parser.NewProperty(fmt.Sprintf("__func__.max.%v.%v", as, stat_field)), stat_term_map["max"])
								planner.SetDocumentProperty(row, parser.NewProperty(fmt.Sprintf("__func__.avg.%v.%v", as, stat_field)), stat_term_map["mean"])
								planner.SetDocumentProperty(row, parser.NewProperty(fmt.Sprintf("__func__.sum.%v.%v", as, stat_field)), stat_term_map["total"])
								//MergeContext(row, stat_term_sum)
								// once we've found what we're looking for
								// break out of the inner loop
								break
							}
						}

					}

				}

				// add this row to the result set
				result = append(result, row)
			}
		}
	}

	return result, nil
}

func ESSupportedLiteralValue(expr parser.Expression) (interface{}, error) {
	switch expr := expr.(type) {
	case *parser.BoolLiteral:
		return expr.Val, nil
	case *parser.IntegerLiteral:
		return expr.Val, nil
	case *parser.FloatLiteral:
		return expr.Val, nil
	case *parser.StringLiteral:
		return expr.Val, nil
	}
	return nil, fmt.Errorf("ElasticSearch does not support comparison with this value")
}

func (ds *ElasticSearchDataSource) ESSupportedProperty(expr parser.Expression) (string, error) {
	prop, ok := expr.(*parser.Property)
	if !ok {
		return "", fmt.Errorf("ElasticSearch LHS must be property")
	}
	if strings.HasPrefix(prop.Symbol, ds.As+".") {
		return prop.Symbol[len(ds.As)+1:], nil
	}
	return "", fmt.Errorf("Property refers to another datasource, not supported")
}

func (ds *ElasticSearchDataSource) BuildElasticSearchFilterRecursive(filter parser.Expression) (*Filter, error) {
	switch expr := filter.(type) {
	case *parser.EqualsExpression:
		lhsprop, err := ds.ESSupportedProperty(expr.Left)
		if err != nil {
			return nil, err
		}
		rhsval, err := ESSupportedLiteralValue(expr.Right)
		if err != nil {
			// not a literal value, maybe its a property in this ds?
			rhsprop, err := ds.ESSupportedProperty(expr.Right)
			if err != nil {
				return nil, err
			}
			return NewScriptFilter(lhsprop, rhsprop, "=="), nil
		}
		return NewTermFilter(lhsprop, rhsval), nil
	case *parser.NotEqualsExpression:
		lhsprop, err := ds.ESSupportedProperty(expr.Left)
		if err != nil {
			return nil, err
		}
		rhsval, err := ESSupportedLiteralValue(expr.Right)
		if err != nil {
			// not a literal value, maybe its a property in this ds?
			rhsprop, err := ds.ESSupportedProperty(expr.Right)
			if err != nil {
				return nil, err
			}
			return NewScriptFilter(lhsprop, rhsprop, "!="), nil
		}
		termFilter := NewTermFilter(lhsprop, rhsval)
		return NewNotFilter(termFilter), nil
	case *parser.LessThanExpression:
		lhsprop, err := ds.ESSupportedProperty(expr.Left)
		if err != nil {
			return nil, err
		}
		rhsval, err := ESSupportedLiteralValue(expr.Right)
		if err != nil {
			// not a literal value, maybe its a property in this ds?
			rhsprop, err := ds.ESSupportedProperty(expr.Right)
			if err != nil {
				return nil, err
			}
			return NewScriptFilter(lhsprop, rhsprop, "<"), nil
		}
		return NewRangeFilter(lhsprop, rhsval, "lt"), nil
	case *parser.GreaterThanExpression:
		lhsprop, err := ds.ESSupportedProperty(expr.Left)
		if err != nil {
			return nil, err
		}
		rhsval, err := ESSupportedLiteralValue(expr.Right)
		if err != nil {
			// not a literal value, maybe its a property in this ds?
			rhsprop, err := ds.ESSupportedProperty(expr.Right)
			if err != nil {
				return nil, err
			}
			return NewScriptFilter(lhsprop, rhsprop, ">"), nil
		}
		return NewRangeFilter(lhsprop, rhsval, "gt"), nil
	case *parser.LessThanOrEqualExpression:
		lhsprop, err := ds.ESSupportedProperty(expr.Left)
		if err != nil {
			return nil, err
		}
		rhsval, err := ESSupportedLiteralValue(expr.Right)
		if err != nil {
			// not a literal value, maybe its a property in this ds?
			rhsprop, err := ds.ESSupportedProperty(expr.Right)
			if err != nil {
				return nil, err
			}
			return NewScriptFilter(lhsprop, rhsprop, "<="), nil
		}
		return NewRangeFilter(lhsprop, rhsval, "lte"), nil
	case *parser.GreaterThanOrEqualExpression:
		lhsprop, err := ds.ESSupportedProperty(expr.Left)
		if err != nil {
			return nil, err
		}
		rhsval, err := ESSupportedLiteralValue(expr.Right)
		if err != nil {
			// not a literal value, maybe its a property in this ds?
			rhsprop, err := ds.ESSupportedProperty(expr.Right)
			if err != nil {
				return nil, err
			}
			return NewScriptFilter(lhsprop, rhsprop, ">="), nil
		}
		return NewRangeFilter(lhsprop, rhsval, "gte"), nil
	case *parser.AndExpression:
		lhs, err := ds.BuildElasticSearchFilterRecursive(expr.Left)
		if err != nil {
			return nil, err
		}
		rhs, err := ds.BuildElasticSearchFilterRecursive(expr.Right)
		if err != nil {
			return nil, err
		}
		return NewAndFilter(lhs, rhs), nil
	case *parser.OrExpression:
		lhs, err := ds.BuildElasticSearchFilterRecursive(expr.Left)
		if err != nil {
			return nil, err
		}
		rhs, err := ds.BuildElasticSearchFilterRecursive(expr.Right)
		if err != nil {
			return nil, err
		}
		return NewOrFilter(lhs, rhs), nil
	case *parser.Property:
		return nil, fmt.Errorf("ElasticSearch does not support filter of bare literal")
	case *parser.StringLiteral:
		return nil, fmt.Errorf("ElasticSearch does not support filter of bare literal")
	case *parser.BoolLiteral:
		return nil, fmt.Errorf("ElasticSearch does not support filter of bare literal")
	case *parser.IntegerLiteral:
		return nil, fmt.Errorf("ElasticSearch does not support filter of bare literal")
	case *parser.FloatLiteral:
		return nil, fmt.Errorf("ElasticSearch does not support filter of bare literal")
	}

	return nil, fmt.Errorf("ElasticSearch does not support filter of Unknown type: %T", filter)
}

func (ds *ElasticSearchDataSource) BuildElasticSearchOrderBy(sortlist parser.SortList) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)
	for _, sortItem := range sortlist {
		sortProperty, ok := sortItem.Sort.(*parser.Property)
		if !ok {
			return nil, fmt.Errorf("ElasticSearch DataSource only support sorting on properties")
		}
		if !sortProperty.IsReferencingDataSource(ds.As) {
			return nil, fmt.Errorf("Property refers to wrong data source")
		}
		sortProperty = sortProperty.Tail()
		if sortItem.Ascending {
			sort := map[string]interface{}{sortProperty.Symbol: map[string]interface{}{"order": "asc"}}
			result = append(result, sort)
		} else {
			sort := map[string]interface{}{sortProperty.Symbol: map[string]interface{}{"order": "desc"}}
			result = append(result, sort)
		}
	}
	return result, nil
}

func (ds *ElasticSearchDataSource) SetDocBodyDataSource(datasource planner.DataSource) {
	ds.docBodyDataSource = datasource
}

func (ds *ElasticSearchDataSource) DocsFromIds(docIds []string) ([]planner.Document, error) {
	panic("Unexpected call to CSV DataSource to get DocsFromIds")
}

func (ds *ElasticSearchDataSource) Cost() float64 {
	return math.Inf(1)
}

func (ds *ElasticSearchDataSource) TotalCost() float64 {
	return ds.Cost()
}
