package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mschoch/elastigo/core"
	"log"
	"strings"
)

const FunctionPrefix = "__func__"

type CouchbaseElasticSearchDoc struct {
	Meta CouchbaseDocMeta `json:"meta"`
}

type CouchbaseDocMeta struct {
	Id         string `json:"id"`
	Rev        string `json:"rev"`
	Expiration int    `json:"expiration"`
	Flags      int    `json:"flags"`
}

func ESDocMetaMatchingQuery(datasource string, query map[string]interface{}) ([]CouchbaseDocMeta, error) {

	if *debugElasticSearch {
		queryBytes, _ := json.Marshal(query)
		queryString := string(queryBytes)
		log.Printf("ES Query JSON is: %v", queryString)
	}

	result := make([]CouchbaseDocMeta, 0)

	resultsSeen := 0
	searchresponse, err := core.Search(true, datasource, "", query, "1m")
	if err != nil {
		return nil, err
	}

	if *debugElasticSearch {
		log.Printf("Response %#v", searchresponse)
	}

	srmeta, err := DocMetaFromSearchResultHits(searchresponse)
	if err != nil {
		return nil, err
	}
	result = concatMetaSlices(result, srmeta)

	resultsSeen += len(searchresponse.Hits.Hits)
	scrollId := searchresponse.ScrollId
	total := searchresponse.Hits.Total

	offset_int := 0
	offset, has_offset := query["from"]
	if has_offset {
		offset_int = offset.(int)
		total = total - offset_int
	}

	limit, has_limit := query["size"]
	if has_limit {
		limit_int := limit.(int)
		if limit_int < total {
			total = limit_int
		}
	}

	for resultsSeen < total {
		searchresponse, err := core.Scroll(true, scrollId, "1m")
		if err != nil {
			return nil, err
		}

		if *debugElasticSearch {
			log.Printf("Response %#v", searchresponse)
		}

		srmeta, err := DocMetaFromSearchResultHits(searchresponse)
		if err != nil {
			return nil, err
		}
		result = concatMetaSlices(result, srmeta)

		resultsSeen += len(searchresponse.Hits.Hits)
		scrollId = searchresponse.ScrollId

	}

	if len(result) != total {
		return nil, errors.New("Length of result set does not match total")
	}

	return result, nil
}

func DocMetaFromSearchResultHits(sr core.SearchResult) ([]CouchbaseDocMeta, error) {
	result := make([]CouchbaseDocMeta, 0)

	for _, val := range sr.Hits.Hits {
		var source CouchbaseElasticSearchDoc
		// marshall into json
		jsonErr := json.Unmarshal(val.Source, &source)
		if jsonErr != nil {
			return nil, jsonErr
		}
		result = append(result, source.Meta)
	}

	return result, nil
}

func concatMetaSlices(old1, old2 []CouchbaseDocMeta) []CouchbaseDocMeta {
	newslice := make([]CouchbaseDocMeta, len(old1)+len(old2))
	copy(newslice, old1)
	copy(newslice[len(old1):], old2)
	return newslice
}

func ESAggregateQuery(datasource string, query map[string]interface{}, group_by string) ([]interface{}, error) {

	if *debugElasticSearch {
		queryBytes, _ := json.Marshal(query)
		queryString := string(queryBytes)
		log.Printf("ES Query JSON is: %v", queryString)
	}

	result := make([]interface{}, 0)

	searchresponse, err := core.Search(true, datasource, "", query, "1m")
	if err != nil {
		return nil, err
	}

	if *debugElasticSearch {
		log.Printf("Response %#v", searchresponse)
	}

	result, err = RowsFromSearchFacets(searchresponse, group_by)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func RowsFromSearchFacets(sr core.SearchResult, group_by string) ([]interface{}, error) {
	result := make([]interface{}, 0)

	var facetResults map[string]interface{}

	// marshall into json
	jsonErr := json.Unmarshal(sr.Facets, &facetResults)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if group_by == "" {
		// if there was no group by, then there is only 1 row

		row := make(map[string]interface{})

		for stat_facet_name, stat_facet_details := range facetResults {

			if strings.HasPrefix(stat_facet_name, StatsPrefix) {
				stat_field := stat_facet_name[len(StatsPrefix):]
				stat_facet_details_map := stat_facet_details.(map[string]interface{})

				stat_term_count := map[string]interface{}{FunctionPrefix + "count": BuildContextRecursive(stat_field, stat_facet_details_map["count"])}
				MergeContext(row, stat_term_count)
				stat_term_min := map[string]interface{}{FunctionPrefix + "min": BuildContextRecursive(stat_field, stat_facet_details_map["min"])}
				MergeContext(row, stat_term_min)
				stat_term_max := map[string]interface{}{FunctionPrefix + "max": BuildContextRecursive(stat_field, stat_facet_details_map["max"])}
				MergeContext(row, stat_term_max)
				stat_term_avg := map[string]interface{}{FunctionPrefix + "avg": BuildContextRecursive(stat_field, stat_facet_details_map["mean"])}
				MergeContext(row, stat_term_avg)
				stat_term_sum := map[string]interface{}{FunctionPrefix + "sum": BuildContextRecursive(stat_field, stat_facet_details_map["total"])}
                MergeContext(row, stat_term_sum)

			}

		}

		// add this row to the result set
		result = append(result, row)
	} else {

		facet_name := group_by
		facet_details := facetResults[facet_name]

		facet_details_map := facet_details.(map[string]interface{})

		//      missing := facet_details_map["missing"]
		//      total := facet_details_map["total"]
		other := facet_details_map["other"]

		if other.(float64) != 0 {
			return nil, fmt.Errorf("Facet results reported %#v \"other\" rows, increate your esMaxAggregate value and try again", other)
		}

		facet_type := facet_details_map["_type"]

		if facet_type == "terms" {
			terms := facet_details_map["terms"]

			for _, term := range terms.([]interface{}) {
				term_map := term.(map[string]interface{})
				//              row := make(map[string]interface{})
				//              row[facet_name] = term_map["term"]
				row := BuildContextRecursive(facet_name, term_map["term"])
				term_count := map[string]interface{}{FunctionPrefix + "count": BuildContextRecursive(facet_name, term_map["count"])}
				MergeContext(row, term_count)

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
								stat_term_count := map[string]interface{}{FunctionPrefix + "count": BuildContextRecursive(stat_field, stat_term_map["count"])}
								MergeContext(row, stat_term_count)
								stat_term_min := map[string]interface{}{FunctionPrefix + "min": BuildContextRecursive(stat_field, stat_term_map["min"])}
								MergeContext(row, stat_term_min)
								stat_term_max := map[string]interface{}{FunctionPrefix + "max": BuildContextRecursive(stat_field, stat_term_map["max"])}
								MergeContext(row, stat_term_max)
								stat_term_avg := map[string]interface{}{FunctionPrefix + "avg": BuildContextRecursive(stat_field, stat_term_map["mean"])}
								MergeContext(row, stat_term_avg)
								stat_term_sum := map[string]interface{}{FunctionPrefix + "sum": BuildContextRecursive(stat_field, stat_term_map["total"])}
                                MergeContext(row, stat_term_sum)
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

func MergeContext(source map[string]interface{}, add map[string]interface{}) {
	for k, v := range add {
		val, ok := source[k]
		if ok {
			// source map already exists for this level
			// recurse deeper
			MergeContext(val.(map[string]interface{}), v.(map[string]interface{}))
		} else {
			// source does not have anything here
			source[k] = v
		}
	}
}

func BuildContextRecursive(property string, value interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	dotIndex := strings.Index(property, ".")
	if dotIndex < 0 {
		result[property] = value
	} else {
		this_property := property[0:dotIndex]
		rest_property := property[dotIndex+1:]
		result[this_property] = BuildContextRecursive(rest_property, value)
	}

	return result
}
