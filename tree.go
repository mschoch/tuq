package main

import (
	"errors"
)

const StatsPrefix = "__stats__"

type Select struct {
	left               *Select
	right              *Select
	mode               string
	distinct           bool
	sel                map[string]interface{}
	from               []DataSource
	where              map[string]interface{}
	groupby            []interface{}
	having             map[string]interface{}
	orderby            []map[string]interface{}
	esQuery            map[string]interface{}
	es_group_by_field  string
	limit              map[string]interface{}
	offset             map[string]interface{}
	parsedSuccessfully bool
	isAggregateQuery   bool
}

type DataSource struct {
	def string
	as  string
}

func NewSelect() *Select {
	return &Select{
		left:               nil,
		right:              nil,
		mode:               "simple",
		distinct:           false,
		sel:                nil,
		from:               make([]DataSource, 0),
		where:              nil,
		groupby:            nil,
		orderby:            make([]map[string]interface{}, 0),
		having:             nil,
		esQuery:            nil,
		es_group_by_field:  "",
		limit:              nil,
		offset:             nil,
		parsedSuccessfully: false,
		isAggregateQuery:   false}
}

func (s *Select) AddDataSource(ds *DataSource) {
	s.from = append(s.from, *ds)
}

func (s *Select) Execute() (interface{}, error) {

	// Make sure we support this query
	err := s.Validate()
	if err != nil {
		return nil, err
	}

	var results interface{}

	if s.isAggregateQuery {

		// Build and ES Query based on our understanding
		s.esQuery = s.BuildESAggregateQuery()

		// Run the query and get back the document meta
		// of each matching document
		// FIXME passing in the group by field is cheating, but helps me find facet
		rows, err := ESAggregateQuery(s.from[0].def, s.esQuery, s.es_group_by_field)

		if err != nil {
			return nil, err
		}

		results = rows

		// Apply the select clause, if any, to muate each row
		if s.sel != nil {
			results, err = ResultsFromDocsAndSelectClause(rows, s.sel)
		}

	} else {

		// Build and ES Query based on our understanding
		s.esQuery = s.BuildESQuery()

		// Run the query and get back the document meta
		// of each matching document
		docMeta, err := ESDocMetaMatchingQuery(s.from[0].def, s.esQuery)

		if err != nil {
			return nil, err
		}

		// Convert the document meta to full wrapped documents (doc and meta section)
		docs, err := DocsFromMeta(s.from[0].def, docMeta)

		if err != nil {
			return nil, err
		}

		results = docs

		// Apply the select clause, if any, to muate each row
		if s.sel != nil {
			results, err = ResultsFromDocsAndSelectClause(docs, s.sel)
		}

	}

	return results, nil
}

func ResultsFromDocsAndSelectClause(docs []interface{}, sel map[string]interface{}) ([]interface{}, error) {

	result := make([]interface{}, 0)
	for _, doc := range docs {

		val, err := EvaluateExpressionInContext(sel["expression"].(map[string]interface{}), doc.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}

func (s *Select) Validate() error {
	if len(s.from) != 1 {
		return errors.New("Unsupported Query: At this time only 1 datasource is supported.")
	}

	if len(s.groupby) > 1 {
		return errors.New("Unsupported Query: At this time only 1 group by expression is supported.")
	}

	// FIXME identify subqueries and warn they are not supported

	return nil
}

func (s *Select) BuildESQuery() map[string]interface{} {
	queryDoc := make(map[string]interface{})
	querySection := make(map[string]interface{})
	filteredSection := make(map[string]interface{})
	innerQuery := make(map[string]interface{})
	matchAll := make(map[string]interface{})

	innerQuery["match_all"] = matchAll
	filteredSection["query"] = innerQuery
	querySection["filtered"] = filteredSection
	queryDoc["query"] = querySection
	queryDoc["size"] = *esBatchSize

	filterSection := BuildESDefaultFilter()
	whereFilter := BuildESQueryRecursive(s.where)
	if whereFilter != nil {
		// combine this filter with the default
		and_section := make(map[string]interface{})
		and_section["and"] = []interface{}{filterSection, whereFilter}
		filterSection = and_section
	}

	queryDoc["filter"] = filterSection

	orderbySection := BuildESOrderBy(s.orderby)
	queryDoc["sort"] = orderbySection

	if s.limit != nil {
		limitSection := BuildESLimit(s.limit)
		queryDoc["size"] = limitSection
	}

	if s.offset != nil {
		offsetSection := BuildESOffset(s.offset)
		queryDoc["from"] = offsetSection
	}

	return queryDoc
}

func (s *Select) BuildESAggregateQuery() map[string]interface{} {
	queryDoc := make(map[string]interface{})
	querySection := make(map[string]interface{})
	filteredSection := make(map[string]interface{})
	innerQuery := make(map[string]interface{})
	matchAll := make(map[string]interface{})

	innerQuery["match_all"] = matchAll
	filteredSection["query"] = innerQuery
	querySection["filtered"] = filteredSection
	queryDoc["query"] = querySection
	queryDoc["size"] = 0 //aggregate queries do not need document matches

	// peek at the select clause
	agg_stats_fields := []string{}
	if s.sel != nil {
		agg_stats_fields = FindPropertiesAndIdentifiers(s.sel["expression"].(map[string]interface{}))
	}

	filterSection := BuildESDefaultFilter()
	whereFilter := BuildESQueryRecursive(s.where)
	if whereFilter != nil {
		// combine this filter with the default
		and_section := make(map[string]interface{})
		and_section["and"] = []interface{}{filterSection, whereFilter}
		filterSection = and_section
	}

	facetSection, group_by := BuildFacets(s.groupby, filterSection, agg_stats_fields)
	s.es_group_by_field = group_by

	queryDoc["facets"] = facetSection

	// FIXME add support for order by and limit to aggregate queries

	return queryDoc
}

func BuildESDefaultFilter() interface{} {
	filter_section := make(map[string]interface{})
	not_section := make(map[string]interface{})
	type_section := make(map[string]interface{})

	type_section["value"] = *esDefaultExcludeType
	not_section["type"] = type_section
	filter_section["not"] = not_section

	return filter_section
}

func BuildFacets(groupby []interface{}, filter interface{}, stats_fields []string) (interface{}, string) {
	facetSection := make(map[string]interface{})
	group_by := ""

	if len(groupby) == 1 {
		gba := make(map[string]interface{})

		gby := groupby[0].(map[string]interface{})

		isIdentifier := false
		val, isProperty := gby["property"]
		if !isProperty {
			val, isIdentifier = gby["identifier"]
		}

		group_by = val.(string)

		facetSection[val.(string)] = gba
		terms := make(map[string]interface{})
		gba["terms"] = terms
		terms["field"] = val
		terms["size"] = *esMaxAggregate
		gba["facet_filter"] = filter

		for _, stat_field := range stats_fields {
			gba := make(map[string]interface{})

			if isProperty || isIdentifier {

				if stat_field != val.(string) {

					facetSection[StatsPrefix+stat_field] = gba
					terms := make(map[string]interface{})
					gba["terms_stats"] = terms
					terms["key_field"] = val
					terms["value_field"] = stat_field
					terms["size"] = *esMaxAggregate
				}
			}

			gba["facet_filter"] = filter

		}
	}

	return facetSection, group_by
}

func BuildESLimit(limit map[string]interface{}) int {
	result := BuildESQueryRecursive(limit["expression"].(map[string]interface{}))
	return result.(int)
}

func BuildESOffset(offset map[string]interface{}) int {
	result := BuildESQueryRecursive(offset["expression"].(map[string]interface{}))
	return result.(int)
}

func BuildESOrderBy(orderby []map[string]interface{}) interface{} {

	result := make([]interface{}, 0)

	for _, order := range orderby {
		thisSection := make(map[string]interface{})
		val, ok := order["ascending"]
		if ok && val.(bool) {
			thisSection[BuildESQueryRecursive(order["expression"].(map[string]interface{})).(string)] = map[string]interface{}{"order": "asc"}
		} else {
			thisSection[BuildESQueryRecursive(order["expression"].(map[string]interface{})).(string)] = map[string]interface{}{"order": "desc"}
		}
		result = append(result, thisSection)
	}

	return result
}

func BuildESQueryRecursive(where map[string]interface{}) interface{} {
	_, isOperation := where["op"]

	if isOperation {

		if where["op"] == "==" {
			thisSection := make(map[string]interface{})
			nextSection := make(map[string]interface{})
			thisSection["term"] = nextSection
			lhs := BuildESQueryRecursive(where["left"].(map[string]interface{}))
			rhs := BuildESQueryRecursive(where["right"].(map[string]interface{}))
			nextSection[lhs.(string)] = rhs
			return thisSection
		} else if where["op"] == "<" {
			thisSection := make(map[string]interface{})
			nextSection := make(map[string]interface{})
			rangeSection := make(map[string]interface{})
			thisSection["range"] = nextSection
			lhs := BuildESQueryRecursive(where["left"].(map[string]interface{}))
			rhs := BuildESQueryRecursive(where["right"].(map[string]interface{}))
			rangeSection["lt"] = rhs
			nextSection[lhs.(string)] = rangeSection
			return thisSection
		} else if where["op"] == ">" {
			thisSection := make(map[string]interface{})
			nextSection := make(map[string]interface{})
			rangeSection := make(map[string]interface{})
			thisSection["range"] = nextSection
			lhs := BuildESQueryRecursive(where["left"].(map[string]interface{}))
			rhs := BuildESQueryRecursive(where["right"].(map[string]interface{}))
			rangeSection["gt"] = rhs
			nextSection[lhs.(string)] = rangeSection
			return thisSection
		} else if where["op"] == "<=" {
			thisSection := make(map[string]interface{})
			nextSection := make(map[string]interface{})
			rangeSection := make(map[string]interface{})
			thisSection["range"] = nextSection
			lhs := BuildESQueryRecursive(where["left"].(map[string]interface{}))
			rhs := BuildESQueryRecursive(where["right"].(map[string]interface{}))
			rangeSection["lte"] = rhs
			nextSection[lhs.(string)] = rangeSection
			return thisSection
		} else if where["op"] == ">=" {
			thisSection := make(map[string]interface{})
			nextSection := make(map[string]interface{})
			rangeSection := make(map[string]interface{})
			thisSection["range"] = nextSection
			lhs := BuildESQueryRecursive(where["left"].(map[string]interface{}))
			rhs := BuildESQueryRecursive(where["right"].(map[string]interface{}))
			rangeSection["gte"] = rhs
			nextSection[lhs.(string)] = rangeSection
			return thisSection
		} else if where["op"] == "&&" {
			thisSection := make(map[string]interface{})
			nextSection := make([]interface{}, 0)
			lhs := BuildESQueryRecursive(where["left"].(map[string]interface{}))
			rhs := BuildESQueryRecursive(where["right"].(map[string]interface{}))
			nextSection = append(nextSection, lhs)
			nextSection = append(nextSection, rhs)
			thisSection["and"] = nextSection
			return thisSection
		} else if where["op"] == "||" {
			thisSection := make(map[string]interface{})
			nextSection := make([]interface{}, 0)
			lhs := BuildESQueryRecursive(where["left"].(map[string]interface{}))
			rhs := BuildESQueryRecursive(where["right"].(map[string]interface{}))
			nextSection = append(nextSection, lhs)
			nextSection = append(nextSection, rhs)
			thisSection["or"] = nextSection
			return thisSection
		}

	} else {
		val, isString := where["string"]
		if isString {
			return val
		}
		val, isProperty := where["property"]
		if isProperty {
			return val
		}
		val, isIdentifier := where["identifier"]
		if isIdentifier {
			return val
		}
		val, isInt := where["int"]
		if isInt {
			return val
		}
		val, isReal := where["real"]
		if isReal {
			return val
		}
		val, isBool := where["bool"]
		if isBool {
			return val
		}
	}
	return nil
}

func NewDataSource(d string) *DataSource {
	return &DataSource{
		def: d,
		as:  d}
}

func NewDataSourceWithAs(d, a string) *DataSource {
	return &DataSource{
		def: d,
		as:  a}
}
