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

var zeroSymbolExpressions = []Expression{
	NewIntegerLiteral(7),
	NewFloatLiteral(5.5),
	NewNull(),
	NewStringLiteral("name"),
	NewBoolLiteral(true),
}

var oneSymbolExpressions = []Expression{
	NewProperty("bob"),
	NewArrayLiteral(ExpressionList{NewProperty("bob")}),
	NewObjectLiteral(Object{"field": NewProperty("bob")}),
	NewNotExpression(NewProperty("bob")),
	NewBracketMemberExpression(NewProperty("bob"), NewIntegerLiteral(0)),
}

var twoSymbolExpressions = []Expression{
	NewPlusExpression(NewProperty("bob"), NewProperty("jay")),
	NewMinusExpression(NewProperty("bob"), NewProperty("jay")),
	NewMultiplyExpression(NewProperty("bob"), NewProperty("jay")),
	NewDivideExpression(NewProperty("bob"), NewProperty("jay")),
	NewOrExpression(NewProperty("bob"), NewProperty("jay")),
	NewAndExpression(NewProperty("bob"), NewProperty("jay")),
	NewLessThanExpression(NewProperty("bob"), NewProperty("jay")),
	NewLessThanOrEqualExpression(NewProperty("bob"), NewProperty("jay")),
	NewGreaterThanExpression(NewProperty("bob"), NewProperty("jay")),
	NewGreaterThanOrEqualExpression(NewProperty("bob"), NewProperty("jay")),
	NewEqualsExpression(NewProperty("bob"), NewProperty("jay")),
	NewNotEqualsExpression(NewProperty("bob"), NewProperty("jay")),
}

var threeSymbolExpressions = []Expression{
	NewTernaryExpression(NewProperty("bob"), NewProperty("jay"), NewProperty("cat")),
}

func TestSymbolsReferenced(t *testing.T) {

	for _, v := range zeroSymbolExpressions {
		symbols := v.SymbolsReferenced()
		if len(symbols) != 0 {
			t.Errorf("Expected 0 symbols and found: %v", symbols)
		}
	}

	for _, v := range oneSymbolExpressions {
		symbols := v.SymbolsReferenced()
		if len(symbols) != 1 {
			t.Errorf("Expected 1 symbols and found: %v", symbols)
		}
		if symbols[0] != "bob" {
			t.Errorf("Expected symbol to be bob, got: %v", symbols[0])
		}
	}

	for _, v := range twoSymbolExpressions {
		symbols := v.SymbolsReferenced()
		if len(symbols) != 2 {
			t.Errorf("Expected 2 symbols and found: %v", symbols)
		}
		if symbols[0] != "bob" {
			t.Errorf("Expected symbol 1 to be bob, got: %v", symbols[0])
		}
		if symbols[1] != "jay" {
			t.Errorf("Expected symbol 2 to be jay, got: %v", symbols[0])
		}
	}

	for _, v := range threeSymbolExpressions {
		symbols := v.SymbolsReferenced()
		if len(symbols) != 3 {
			t.Errorf("Expected 3 symbols and found: %v", symbols)
		}
		if symbols[0] != "bob" {
			t.Errorf("Expected symbol 1 to be bob, got: %v", symbols[0])
		}
		if symbols[1] != "jay" {
			t.Errorf("Expected symbol 2 to be jay, got: %v", symbols[0])
		}
		if symbols[2] != "cat" {
			t.Errorf("Expected symbol 3 to be cat, got: %v", symbols[0])
		}
	}

	// manually test function call, couldn't be set up as literal
	f := NewFunction("sum")
	f.AddArguments(ExpressionList{NewProperty("bob")})
	symbols := f.SymbolsReferenced()
	if len(symbols) != 1 {
		t.Errorf("Expected 1 symbols and found: %v", symbols)
	}
	if symbols[0] != "bob" {
		t.Errorf("Expected symbol to be bob, got: %v", symbols[0])
	}
}

func TestPrefixSymbols(t *testing.T) {

	for _, v := range zeroSymbolExpressions {
		v.PrefixSymbols("ds.")
		symbols := v.SymbolsReferenced()
		if len(symbols) != 0 {
			t.Errorf("Expected 0 symbols and found: %v", symbols)
		}
	}

	for _, v := range oneSymbolExpressions {
		v.PrefixSymbols("ds.")
		symbols := v.SymbolsReferenced()
		if len(symbols) != 1 {
			t.Errorf("Expected 1 symbols and found: %v", symbols)
		}
		if symbols[0] != "ds.bob" {
			t.Errorf("Expected symbol to be ds.bob, got: %v", symbols[0])
		}
	}

	for _, v := range twoSymbolExpressions {
		v.PrefixSymbols("ds.")
		symbols := v.SymbolsReferenced()
		if len(symbols) != 2 {
			t.Errorf("Expected 2 symbols and found: %v", symbols)
		}
		if symbols[0] != "ds.bob" {
			t.Errorf("Expected symbol 1 to be ds.bob, got: %v", symbols[0])
		}
		if symbols[1] != "ds.jay" {
			t.Errorf("Expected symbol 2 to be ds.jay, got: %v", symbols[0])
		}
	}

	for _, v := range threeSymbolExpressions {
		v.PrefixSymbols("ds.")
		symbols := v.SymbolsReferenced()
		if len(symbols) != 3 {
			t.Errorf("Expected 3 symbols and found: %v", symbols)
		}
		if symbols[0] != "ds.bob" {
			t.Errorf("Expected symbol 1 to be ds.bob, got: %v", symbols[0])
		}
		if symbols[1] != "ds.jay" {
			t.Errorf("Expected symbol 2 to be ds.jay, got: %v", symbols[0])
		}
		if symbols[2] != "ds.cat" {
			t.Errorf("Expected symbol 3 to be ds.cat, got: %v", symbols[0])
		}
	}

	// manually test function call, couldn't be set up as literal
	f := NewFunction("sum")
	f.AddArguments(ExpressionList{NewProperty("bob")})
	f.PrefixSymbols("ds.")
	symbols := f.SymbolsReferenced()
	if len(symbols) != 1 {
		t.Errorf("Expected 1 symbols and found: %v", symbols)
	}
	if symbols[0] != "ds.bob" {
		t.Errorf("Expected symbol to be ds.bob, got: %v", symbols[0])
	}
}