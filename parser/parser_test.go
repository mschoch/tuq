package parser

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
	"SELECT {\"name\":doc.name,\"abv\":[doc.abv, doc.ibu]} FROM beer-sample WHERE doc.type == \"beer\" && doc.abv > 9",
	"SELECT {\"name\":doc.name,\"abv\":{\"abv\":doc.abv}} FROM beer-sample WHERE doc.type == \"beer\" && doc.abv > 9",
	"SELECT {\"name\":doc.name,\"literal_int\":7} FROM beer-sample WHERE doc.type == \"beer\" && doc.abv > 9",
	"SELECT {\"name\":doc.name,\"literal_bool\":true} FROM beer-sample WHERE doc.type == \"beer\" && doc.abv > 9",
	"SELECT {\"name\":doc.name,\"literal_string\":\"string\"} FROM beer-sample WHERE doc.type == \"beer\" && doc.abv > 9",
	"SELECT FROM beer-sample WHERE doc.abv > 5 GROUP BY doc.type",
	"SELECT FROM beer-sample WHERE _underscore_identifier > 4",
	"SELECT FROM beer-sample WHERE doc[\"abv\"] > 7",
	"SELECT FROM orders AS o OVER o.items AS item",
	"SELECT 1+1 FROM orders",
	"SELECT 1-1 FROM orders",
	"SELECT 1*1 FROM orders",
	"SELECT 1/1 FROM orders",
	"SELECT MAX(doc.abv) FROM beer-sample",
	"SELECT FROM beer-sample WHERE doc.abv >= 15",
	"SELECT FROM beer-sample WHERE doc.abv <= 15",
	"SELECT FROM beer-sample WHERE doc.abv != 15",
	"SELECT FROM beer-sample WHERE doc.abv == 15.3",
	"SELECT FROM beer-sample WHERE doc.abv != null",
	"SELECT FROM beer-sample WHERE doc.abv > 5 || doc.ibu < 3",
	"SELECT FROM beer-sample WHERE !(doc.abv < 4)",
	"SELECT doc.abv > 0 ? doc.abv : null FROM beer-sample",
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

	unqlParser := NewUnqlParser(false, false, true)

	for _, v := range validQueries {
		_, err := unqlParser.Parse(v)
		if err != nil {
			t.Errorf("Valid Query Parse Failed: %v - %v", v, err)
		}
	}

	for _, v := range invalidQueries {
		_, err := unqlParser.Parse(v)
		if err == nil {
			t.Errorf("Invalid Query Parsed Successfully: %v - %v", v, err)
		}
	}

}
