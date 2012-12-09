package mongodb

import (
	"fmt"
	"github.com/mschoch/tuq/datasources"
	"github.com/mschoch/tuq/parser"
	"github.com/mschoch/tuq/planner"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"strings"
)

const StatsPrefix = "__stats__"
const defaultBatchSize = 10000
const defaultDebug = false

type MongoDBDataSource struct {
	Url               string
	Name              string
	As                string
	OutputChannel     planner.DocumentChannel
	cancelled         bool
	limit             int
	offset            int
	sort              bson.D
	sortlist          parser.SortList
	docBodyDataSource planner.DataSource
	indexName         string
	filterExpression  parser.Expression
	batchSize         int
	debug             bool
	database          string
	collection        string
	filter            *Filter
}

func init() {
	datasources.RegisterDataSourceImpl("mongodb", NewMongoDBDataSource)
}

func NewMongoDBDataSource(config map[string]interface{}) planner.DataSource {

	result := &MongoDBDataSource{
		OutputChannel: make(planner.DocumentChannel),
		cancelled:     false,
		limit:         -1,
		offset:        0,
		Url:           config["mongodb"].(string),
		database:      config["database"].(string),
		collection:    config["collection"].(string),
		filter:        EmptyFilter()}

	if config["batch_size"] != nil {
		result.batchSize = config["batch_size"].(int)
	} else {
		result.batchSize = defaultBatchSize
	}

	if config["debug"] != nil {
		result.debug = config["debug"].(bool)
	} else {
		result.debug = defaultDebug
	}

	return result
}

// PlanPipelineComponent interface

func (ds *MongoDBDataSource) SetSource(source planner.PlanPipelineComponent) {
	log.Fatalf("SetSource called on DataSource")
}

func (ds *MongoDBDataSource) GetSource() planner.PlanPipelineComponent {
	return nil
}

func (ds *MongoDBDataSource) GetDocumentChannel() planner.DocumentChannel {
	return ds.OutputChannel
}

func (ds *MongoDBDataSource) Run() {
	defer close(ds.OutputChannel)

	session, err := mgo.Dial(ds.Url)
	if err != nil {
		log.Printf("Error connecting to MongoDB")
		return
	}
	c := session.DB(ds.database).C(ds.collection)

	theq := map[string]interface{}{"$query": ds.filter,
		"$orderby": ds.sort}
	q := c.Find(theq)

	q.Batch(ds.batchSize)
	if ds.limit > 0 {
		q.Limit(ds.limit)
	}
	if ds.offset > 0 {
		q.Skip(ds.offset)
	}

	iter := q.Iter()
	result := make(map[string]interface{})
	for iter.Next(&result) {
		// add as if necessary
		if ds.As != "" {
			doccopy := result
			result = make(map[string]interface{})
			result[ds.As] = doccopy
		}
		ds.OutputChannel <- result
		result = make(map[string]interface{})
	}
	if iter.Err() != nil {
		log.Printf("got error %v", iter.Err())
	}

}

func (ds *MongoDBDataSource) Explain() {
	defer close(ds.OutputChannel)

	thisStep := map[string]interface{}{
		"_type":      "FROM",
		"impl":       "MongoDB",
		"database":   ds.database,
		"collection": ds.collection,
		"name":       ds.Name,
		"as":         ds.As,
		"filter":     ds.filter,
		"limit":      ds.limit,
		"offset":     ds.offset}

	ds.OutputChannel <- thisStep
}

func (ds *MongoDBDataSource) Cancel() {
	ds.cancelled = true
}

// DataSource Interface

func (ds *MongoDBDataSource) SetName(name string) {
	ds.Name = name
}

func (ds *MongoDBDataSource) SetAs(as string) {
	ds.As = as
}

func (ds *MongoDBDataSource) GetAs() string {
	return ds.As
}

func (ds *MongoDBDataSource) SetFilter(filter parser.Expression) error {
	f, err := ds.BuildMongoDBFilterRecursive(filter)
	if err != nil {
		return err
	} else {
		ds.filter = f
		ds.filterExpression = filter
	}
	return nil
}

func (ds *MongoDBDataSource) GetFilter() parser.Expression {
	return ds.filterExpression
}

func (ds *MongoDBDataSource) SetOrderBy(sortlist parser.SortList) error {
	s, err := ds.BuildMongoDBOrderBy(sortlist)
	if err != nil {
		return err
	}
	ds.sort = s
	return nil
}

func (ds *MongoDBDataSource) GetOrderBy() parser.SortList {
	return ds.sortlist
}

func (ds *MongoDBDataSource) SetLimit(expr parser.Expression) error {
	switch expr := expr.(type) {
	case *parser.IntegerLiteral:
		ds.limit = expr.Val
	default:
		return fmt.Errorf("MongoDB DataSource only supports limiting by integer literals")
	}
	return nil
}

func (ds *MongoDBDataSource) SetOffset(expr parser.Expression) error {
	switch expr := expr.(type) {
	case *parser.IntegerLiteral:
		ds.offset = expr.Val
	default:
		return fmt.Errorf("MongoDB DataSource only supports offsetting by integer literals")
	}
	return nil
}

func (ds *MongoDBDataSource) SetGroupByWithStatsFields(groupby parser.ExpressionList, stats_fields []string) error {
	return fmt.Errorf("MongoDB DataSource does not yet support group by")
}

func (ds *MongoDBDataSource) SetHaving(having parser.Expression) error {
	return fmt.Errorf("MongoDB DataSource does not yet support having")
}

func MongoDBSupportedLiteralValue(expr parser.Expression) (interface{}, error) {
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
	return nil, fmt.Errorf("MongoDB does not support comparison with this value")
}

func (ds *MongoDBDataSource) MongoDBSupportedProperty(expr parser.Expression) (string, error) {
	prop, ok := expr.(*parser.Property)
	if !ok {
		return "", fmt.Errorf("MongoDB LHS must be property")
	}
	if strings.HasPrefix(prop.Symbol, ds.As+".") {
		return prop.Symbol[len(ds.As)+1:], nil
	}
	return "", fmt.Errorf("Property refers to another datasource, not supported")
}

func (ds *MongoDBDataSource) BuildMongoDBFilterRecursive(filter parser.Expression) (*Filter, error) {
	switch expr := filter.(type) {
	case *parser.EqualsExpression:
		lhsprop, err := ds.MongoDBSupportedProperty(expr.Left)
		if err != nil {
			return nil, err
		}
		rhsval, err := MongoDBSupportedLiteralValue(expr.Right)
		if err != nil {
			return nil, err
		}
		return NewEqualsFilter(lhsprop, rhsval), nil
	case *parser.NotEqualsExpression:
		lhsprop, err := ds.MongoDBSupportedProperty(expr.Left)
		if err != nil {
			return nil, err
		}
		rhsval, err := MongoDBSupportedLiteralValue(expr.Right)
		if err != nil {
			return nil, err
		}
		eqFilter := NewEqualsFilter(lhsprop, rhsval)
		return NewNotFilter(eqFilter), nil
	case *parser.LessThanExpression:
		lhsprop, err := ds.MongoDBSupportedProperty(expr.Left)
		if err != nil {
			return nil, err
		}
		rhsval, err := MongoDBSupportedLiteralValue(expr.Right)
		if err != nil {
			return nil, err
		}
		return NewRangeFilter(lhsprop, rhsval, "$lt"), nil
	case *parser.GreaterThanExpression:
		lhsprop, err := ds.MongoDBSupportedProperty(expr.Left)
		if err != nil {
			return nil, err
		}
		rhsval, err := MongoDBSupportedLiteralValue(expr.Right)
		if err != nil {
			return nil, err
		}
		return NewRangeFilter(lhsprop, rhsval, "$gt"), nil
	case *parser.LessThanOrEqualExpression:
		lhsprop, err := ds.MongoDBSupportedProperty(expr.Left)
		if err != nil {
			return nil, err
		}
		rhsval, err := MongoDBSupportedLiteralValue(expr.Right)
		if err != nil {
			return nil, err
		}
		return NewRangeFilter(lhsprop, rhsval, "$lte"), nil
	case *parser.GreaterThanOrEqualExpression:
		lhsprop, err := ds.MongoDBSupportedProperty(expr.Left)
		if err != nil {
			return nil, err
		}
		rhsval, err := MongoDBSupportedLiteralValue(expr.Right)
		if err != nil {
			return nil, err
		}
		return NewRangeFilter(lhsprop, rhsval, "$gte"), nil
	case *parser.AndExpression:
		lhs, err := ds.BuildMongoDBFilterRecursive(expr.Left)
		if err != nil {
			return nil, err
		}
		rhs, err := ds.BuildMongoDBFilterRecursive(expr.Right)
		if err != nil {
			return nil, err
		}
		return NewAndFilter(lhs, rhs), nil
	case *parser.OrExpression:
		lhs, err := ds.BuildMongoDBFilterRecursive(expr.Left)
		if err != nil {
			return nil, err
		}
		rhs, err := ds.BuildMongoDBFilterRecursive(expr.Right)
		if err != nil {
			return nil, err
		}
		return NewOrFilter(lhs, rhs), nil
	case *parser.Property:
		return nil, fmt.Errorf("MongoDB does not support filter of bare literal")
	case *parser.StringLiteral:
		return nil, fmt.Errorf("MongoDB does not support filter of bare literal")
	case *parser.BoolLiteral:
		return nil, fmt.Errorf("MongoDB does not support filter of bare literal")
	case *parser.IntegerLiteral:
		return nil, fmt.Errorf("MongoDB does not support filter of bare literal")
	case *parser.FloatLiteral:
		return nil, fmt.Errorf("MongoDB does not support filter of bare literal")
	}

	return nil, fmt.Errorf("MongoDB does not support filter of Unknown type: %T", filter)
}

func (ds *MongoDBDataSource) SetDocBodyDataSource(datasource planner.DataSource) {
	ds.docBodyDataSource = datasource
}

func (ds *MongoDBDataSource) DocsFromIds(docIds []string) ([]planner.Document, error) {
	panic("Unexpected call to MongoDB DataSource to get DocsFromIds")
}

func (ds *MongoDBDataSource) BuildMongoDBOrderBy(sortlist parser.SortList) (bson.D, error) {
	result := make(bson.D, 0)
	for _, sortItem := range sortlist {
		sortProperty, ok := sortItem.Sort.(*parser.Property)
		if !ok {
			return nil, fmt.Errorf("MongoDB DataSource only support sorting on properties")
		}
		if !sortProperty.IsReferencingDataSource(ds.As) {
			return nil, fmt.Errorf("Property refers to wrong data source")
		}
		sortProperty = sortProperty.Tail()
		if sortItem.Ascending {
			sort := bson.DocElem{Name: sortProperty.Symbol, Value: 1}
			result = append(result, sort)
		} else {
			sort := bson.DocElem{Name: sortProperty.Symbol, Value: -1}
			result = append(result, sort)
		}
	}
	return result, nil
}
