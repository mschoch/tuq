package main

import (
	"errors"
	"github.com/mschoch/go-unql-couchbasev2/parser"
	"log"
)

// this entire file is temporary
// the idea is to migrate the existing backend to the new frontend
// then once its stable again, rewrite the backend

const StatsPrefix = "__stats__"

type QueryExecutor struct {
}

func NewQueryExecutor() *QueryExecutor {
	return &QueryExecutor{}
}

func (qe *QueryExecutor) Execute(query parser.Select) (interface{}, error) {

	// Make sure we support this query
	err := Validate(query)
	if err != nil {
		return nil, err
	}

	var results interface{}

	esQuery := map[string]interface{}{}
	es_group_by_field := ""
	if query.IsAggregateQuery() {

		// Build and ES Query based on our understanding
		esQuery, es_group_by_field = BuildESAggregateQuery(query)

		// Run the query and get back the document meta
		// of each matching document
		// FIXME passing in the group by field is cheating, but helps me find facet
		rows, err := ESAggregateQuery(query.From[0].Def, esQuery, es_group_by_field)

		if err != nil {
			return nil, err
		}

		results = rows

		// Apply the select clause, if any, to muate each row
		if query.Sel != nil {
			results, err = ResultsFromDocsAndSelectClause(rows, query.Sel)
		}

	} else {

		// Build and ES Query based on our understanding
		esQuery = BuildESQuery(query)

		// Run the query and get back the document meta
		// of each matching document
		docMeta, err := ESDocMetaMatchingQuery(query.From[0].Def, esQuery)

		if err != nil {
			return nil, err
		}

		// Convert the document meta to full wrapped documents (doc and meta section)
		docs, err := DocsFromMeta(query.From[0].Def, docMeta)

		if err != nil {
			return nil, err
		}

		results = docs

		// Apply the select clause, if any, to muate each row
		if query.Sel != nil {
			results, err = ResultsFromDocsAndSelectClause(docs, query.Sel)
		}

	}

	return results, nil
}

func Validate(query parser.Select) error {
	if len(query.From) != 1 {
		return errors.New("Unsupported Query: At this time only 1 datasource is supported.")
	}

	if len(query.Groupby) > 1 {
		return errors.New("Unsupported Query: At this time only 1 group by expression is supported.")
	}

	// FIXME identify subqueries and warn they are not supported

	return nil
}

func BuildESAggregateQuery(query parser.Select) (map[string]interface{}, string) {
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
	if query.Sel != nil {
		agg_stats_fields = FindPropertiesAndIdentifiers(query.Sel)
	}

	filterSection := BuildESDefaultFilter()
	whereFilter := BuildESQueryRecursive(query.Where)
	if whereFilter != nil {
		// combine this filter with the default
		and_section := make(map[string]interface{})
		and_section["and"] = []interface{}{filterSection, whereFilter}
		filterSection = and_section
	}

	facetSection, group_by := BuildFacets(query.Groupby, filterSection, agg_stats_fields)

	queryDoc["facets"] = facetSection

	// FIXME add support for order by and limit to aggregate queries

	return queryDoc, group_by
}

func BuildESQuery(query parser.Select) map[string]interface{} {
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
	whereFilter := BuildESQueryRecursive(query.Where)
	if whereFilter != nil {
		// combine this filter with the default
		and_section := make(map[string]interface{})
		and_section["and"] = []interface{}{filterSection, whereFilter}
		filterSection = and_section
	}

	queryDoc["filter"] = filterSection

	orderbySection := BuildESOrderBy(query.Orderby)
	queryDoc["sort"] = orderbySection

	if query.Limit != nil {
		limitSection := BuildESLimit(query.Limit)
		queryDoc["size"] = limitSection
	}

	if query.Offset != nil {
		offsetSection := BuildESOffset(query.Offset)
		queryDoc["from"] = offsetSection
	}

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

func BuildFacets(groupby parser.ExpressionList, filter interface{}, stats_fields []string) (interface{}, string) {
	facetSection := make(map[string]interface{})
	group_by := ""

	if len(groupby) == 0 {
		// if there is no group by, we just add for the stats_fields
		// and use stats type directly instead of terms_stats
		for _, stat_field := range stats_fields {
			gba := make(map[string]interface{})

			facetSection[StatsPrefix+stat_field] = gba
			stats := make(map[string]interface{})
			gba["statistical"] = stats
			stats["field"] = stat_field
			stats["size"] = *esMaxAggregate

			gba["facet_filter"] = filter

		}

	} else if len(groupby) == 1 {
		gba := make(map[string]interface{})

		gby := groupby[0]

		switch val := gby.(type) {
		case parser.Property:
			group_by = val.Symbol

			facetSection[group_by] = gba
			terms := make(map[string]interface{})
			gba["terms"] = terms
			terms["field"] = val.Symbol
			terms["size"] = *esMaxAggregate
			gba["facet_filter"] = filter

			for _, stat_field := range stats_fields {
				gba := make(map[string]interface{})

				if stat_field != val.Symbol {

					facetSection[StatsPrefix+stat_field] = gba
					terms := make(map[string]interface{})
					gba["terms_stats"] = terms
					terms["key_field"] = val.Symbol
					terms["value_field"] = stat_field
					terms["size"] = *esMaxAggregate
				}

				gba["facet_filter"] = filter

			}
		}
	}

	return facetSection, group_by
}

func BuildESLimit(limit parser.Expression) int {
	result := BuildESQueryRecursive(limit)
	return result.(int)
}

func BuildESOffset(offset parser.Expression) int {
	result := BuildESQueryRecursive(offset)
	return result.(int)
}

func BuildESOrderBy(orderby parser.SortList) interface{} {

	result := make([]interface{}, 0)

	for _, order := range orderby {
		thisSection := make(map[string]interface{})
		if order.Ascending {
			thisSection[BuildESQueryRecursive(order.Sort).(string)] = map[string]interface{}{"order": "asc"}
		} else {
			thisSection[BuildESQueryRecursive(order.Sort).(string)] = map[string]interface{}{"order": "desc"}
		}
		result = append(result, thisSection)
	}

	return result
}

func BuildESQueryRecursive(where parser.Expression) interface{} {

	switch exp := where.(type) {
	case parser.EqualsExpression:
		thisSection := make(map[string]interface{})
		nextSection := make(map[string]interface{})
		thisSection["term"] = nextSection
		lhs := BuildESQueryRecursive(exp.Left)
		rhs := BuildESQueryRecursive(exp.Right)
		nextSection[lhs.(string)] = rhs
		return thisSection
	case parser.NotEqualsExpression:
		notSection := make(map[string]interface{})
		thisSection := make(map[string]interface{})
		nextSection := make(map[string]interface{})
		notSection["not"] = thisSection
		thisSection["term"] = nextSection
		lhs := BuildESQueryRecursive(exp.Left)
		rhs := BuildESQueryRecursive(exp.Right)
		nextSection[lhs.(string)] = rhs
		return notSection
	case parser.LessThanExpression:
		thisSection := make(map[string]interface{})
		nextSection := make(map[string]interface{})
		rangeSection := make(map[string]interface{})
		thisSection["range"] = nextSection
		lhs := BuildESQueryRecursive(exp.Left)
		rhs := BuildESQueryRecursive(exp.Right)
		rangeSection["lt"] = rhs
		nextSection[lhs.(string)] = rangeSection
		return thisSection
	case parser.GreaterThanExpression:
		thisSection := make(map[string]interface{})
		nextSection := make(map[string]interface{})
		rangeSection := make(map[string]interface{})
		thisSection["range"] = nextSection
		lhs := BuildESQueryRecursive(exp.Left)
		rhs := BuildESQueryRecursive(exp.Right)
		rangeSection["gt"] = rhs
		nextSection[lhs.(string)] = rangeSection
		return thisSection
	case parser.LessThanOrEqualExpression:
		thisSection := make(map[string]interface{})
		nextSection := make(map[string]interface{})
		rangeSection := make(map[string]interface{})
		thisSection["range"] = nextSection
		lhs := BuildESQueryRecursive(exp.Left)
		rhs := BuildESQueryRecursive(exp.Right)
		rangeSection["lte"] = rhs
		nextSection[lhs.(string)] = rangeSection
		return thisSection
	case parser.GreaterThanOrEqualExpression:
		thisSection := make(map[string]interface{})
		nextSection := make(map[string]interface{})
		rangeSection := make(map[string]interface{})
		thisSection["range"] = nextSection
		lhs := BuildESQueryRecursive(exp.Left)
		rhs := BuildESQueryRecursive(exp.Right)
		rangeSection["gte"] = rhs
		nextSection[lhs.(string)] = rangeSection
		return thisSection
	case parser.AndExpression:
		thisSection := make(map[string]interface{})
		nextSection := make([]interface{}, 0)
		lhs := BuildESQueryRecursive(exp.Left)
		rhs := BuildESQueryRecursive(exp.Right)
		nextSection = append(nextSection, lhs)
		nextSection = append(nextSection, rhs)
		thisSection["and"] = nextSection
		return thisSection
	case parser.OrExpression:
		thisSection := make(map[string]interface{})
		nextSection := make([]interface{}, 0)
		lhs := BuildESQueryRecursive(exp.Left)
		rhs := BuildESQueryRecursive(exp.Right)
		nextSection = append(nextSection, lhs)
		nextSection = append(nextSection, rhs)
		thisSection["or"] = nextSection
		return thisSection
	case parser.Property:
		return exp.Symbol
	case parser.StringLiteral:
		return exp.Val
	case parser.BoolLiteral:
		return exp.Val
	case parser.IntegerLiteral:
		return exp.Val
	case parser.FloatLiteral:
		return exp.Val
	}

	return nil
}

func ResultsFromDocsAndSelectClause(docs []interface{}, sel parser.Expression) ([]interface{}, error) {

	result := make([]interface{}, 0)
	for _, doc := range docs {
		val, err := EvaluateExpressionInContext(sel, doc.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}
