package main

import (
	"testing"
)

var validQueries = []string{
	"SELECT FROM beer-sample",
	"SELECT FROM beer-sample WHERE doc.abv > 5",
	"SELECT FROM beer-sample WHERE doc.abv > 5 && doc.type == \"beer\"",
	"SELECT FROM beer-sample WHERE doc.abv > 5 && doc.type == \"beer\" && doc.ibu < 30",
	"SELECT FROM beer-sample ORDER BY doc.abv",
    "SELECT FROM beer-sample WHERE doc.abv > 5 ORDER BY doc.abv",
    "SELECT FROM beer-sample WHERE doc.abv > 5 && doc.type == \"beer\" ORDER BY doc.abv",
    "SELECT FROM beer-sample WHERE doc.abv > 5 && doc.type == \"beer\" && doc.ibu < 30 ORDER BY doc.abv",
    "SELECT FROM beer-sample ORDER BY doc.abv LIMIT 5",
    "SELECT FROM beer-sample WHERE doc.abv > 5 ORDER BY doc.abv LIMIT 5",
    "SELECT FROM beer-sample WHERE doc.abv > 5 && doc.type == \"beer\" ORDER BY doc.abv LIMIT 5",
    "SELECT FROM beer-sample WHERE doc.abv > 5 && doc.type == \"beer\" && doc.ibu < 30 ORDER BY doc.abv LIMIT 5",
    "SELECT FROM beer-sample ORDER BY doc.abv LIMIT 5 OFFSET 2",
    "SELECT FROM beer-sample WHERE doc.abv > 5 ORDER BY doc.abv LIMIT 5 OFFSET 2",
    "SELECT FROM beer-sample WHERE doc.abv > 5 && doc.type == \"beer\" ORDER BY doc.abv LIMIT 5 OFFSET 2",
    "SELECT FROM beer-sample WHERE doc.abv > 5 && doc.type == \"beer\" && doc.ibu < 30 ORDER BY doc.abv LIMIT 5 OFFSET 2",
	"SELECT doc.abv FROM beer-sample WHERE doc.type == \"beer\" && doc.abv > 9",
	"SELECT {\"name\":doc.name,\"abv\":doc.abv} FROM beer-sample WHERE doc.type == \"beer\" && doc.abv > 9",
	"SELECT {\"name\":doc.name,\"abv\":[doc.abv]} FROM beer-sample WHERE doc.type == \"beer\" && doc.abv > 9",
	"SELECT {\"name\":doc.name,\"abv\":{\"abv\":doc.abv}} FROM beer-sample WHERE doc.type == \"beer\" && doc.abv > 9",
	"SELECT {\"name\":doc.name,\"literal_int\":7} FROM beer-sample WHERE doc.type == \"beer\" && doc.abv > 9",
	"SELECT {\"name\":doc.name,\"literal_bool\":true} FROM beer-sample WHERE doc.type == \"beer\" && doc.abv > 9",
	"SELECT {\"name\":doc.name,\"literal_string\":\"string\"} FROM beer-sample WHERE doc.type == \"beer\" && doc.abv > 9",
	"SELECT FROM beer-sample WHERE doc.abv > 5 GROUP BY doc.type",
	"SELECT FROM beer-sample WHERE _underscore_identifier > 4",
}

var invalidQueries = []string{
	"SELECT WHERE",
	"SELECT FROM beer-sample WHERE",
	"SELECT FROM beer-sample WHERE doc.abv > 5 &&",
	"SELECT FROM beer-sample WHERE doc.abv > 5 && doc.type ==\"beer\" &&",
	"SELECT FROM beer-sample WHERE doc.abv > 5 ORDER",
	"SELECT FROM beer-sample WHERE doc.abv > 5 ORDER BY",
	"SELECT FROM beer-sample WHERE doc.abv > 5 ORDER BY LIMIT",
	"SELECT FROM beer-sample WHERE doc.abv > 5 ORDER BY LIMIT GROUP",
	"SELECT doc.abv, doc.name FROM beer-sample",
}

func TestParser(t *testing.T) {

	for _, v := range validQueries {
		_, err := processNextLine(v)
		if err != nil {
			t.Errorf("Valid Query Parse Failed: %v", v)
		}
	}

	for _, v := range invalidQueries {
		_, err := processNextLine(v)
		if err == nil {
			t.Errorf("Invalid Query Parsed Successfully: %v", v)
		}
	}

}
