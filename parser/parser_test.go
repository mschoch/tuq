package parser

import (
	"testing"
)

var validQueries = []string{
	"SELECT FROM beer-sample",
	"SELECT DISTINCT FROM beer-sample",
	"SELECT FROM beer-sample WHERE doc.abv > 5",
	"SELECT FROM beer-sample WHERE doc.abv > -5",
	"SELECT FROM beer-sample WHERE doc.abv > 5 && doc.type == \"beer\"",
	"SELECT FROM beer-sample WHERE doc.abv > 5 && doc.type == \"beer\" && doc.ibu < 30",
	"SELECT FROM beer-sample ORDER BY doc.abv",
	"SELECT FROM beer-sample WHERE doc.abv > 5 ORDER BY doc.abv",
	"SELECT FROM beer-sample WHERE doc.abv > 5 ORDER BY doc.abv ASC",
	"SELECT FROM beer-sample WHERE doc.abv > 5 ORDER BY doc.abv DESC",
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
	"SELECT FROM orders OVER orders.items AS item",
	"SELECT FROM orders AS o OVER o.items AS item",
	"SELECT 1+1 FROM orders",
	"SELECT 1-1 FROM orders",
	"SELECT 1*1 FROM orders",
	"SELECT 1/1 FROM orders",
	"SELECT MAX(doc.abv) FROM beer-sample",
	"SELECT MIN(doc.abv) FROM beer-sample",
	"SELECT COUNT(doc.abv) FROM beer-sample",
	"SELECT AVG(doc.abv) FROM beer-sample",
	"SELECT SUM(doc.abv) FROM beer-sample",
	"SELECT FROM beer-sample WHERE doc.abv >= 15",
	"SELECT FROM beer-sample WHERE doc.abv <= 15",
	"SELECT FROM beer-sample WHERE doc.abv != 15",
	"SELECT FROM beer-sample WHERE doc.abv == 15.3",
	"SELECT FROM beer-sample WHERE doc.abv == -15.3",
	"SELECT FROM beer-sample WHERE doc.abv != null",
	"SELECT FROM beer-sample WHERE doc.abv > 5 || doc.ibu < 3",
	"SELECT FROM beer-sample WHERE !(doc.abv < 4)",
	"SELECT doc.abv > 0 ? doc.abv : null FROM beer-sample",
	"SELECT FROM beer-sample WHERE doc.abv > 5 GROUP BY doc.type HAVING doc.type != \"bob\"",
	"SELECT FROM beer-sample WHERE doc.type == 'beer'",
	"SELECT FROM beer-sample WHERE doc.type != 100 % 4",
	"SELECT FROM beer-sample WHERE doc.type <> 100 % 4",
	"SELECT FROM beer-sample AS bs UNION SELECT FROM beer-sample AS bs2",
	"SELECT FROM beer-sample AS bs UNION ALL SELECT FROM beer-sample AS bs2",
	"SELECT FROM beer-sample AS bs INTERSECT SELECT FROM beer-sample AS bs2",
	"SELECT FROM beer-sample AS bs EXCEPT SELECT FROM beer-sample AS bs2",
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
	"@abc",
}

func TestParser(t *testing.T) {

	unqlParser := NewUnqlParser(false, false, true)

	for _, v := range validQueries {
		pq, err := unqlParser.Parse(v)
		if err != nil {
			t.Errorf("Valid Query Parse Failed: %v - %v", v, err)
		}
		if !pq.WasParsedSuccessfully() {
			t.Errorf("Valid Query was not parsed successfully: %v - %v", v, err)
		}
		if pq.IsExplainOnly() != false {
			t.Errorf("Explain only should be false")
		}

		err = pq.Validate()
		if err != nil {
			t.Errorf("Error validating query: %v", err)
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

var testExpressions = []Expression{
	NewIntegerLiteral(7),
	NewFloatLiteral(5.5),
	NewNull(),
	NewStringLiteral("name"),
	NewBoolLiteral(true),
	NewProperty("bob"),
	NewArrayLiteral(ExpressionList{NewProperty("bob"), NewProperty("jay")}),
	NewNotExpression(NewProperty("bob")),
	NewBracketMemberExpression(NewProperty("bob"), NewIntegerLiteral(0)),
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
	NewTernaryExpression(NewProperty("bob"), NewProperty("jay"), NewProperty("cat")),
}

var testExpressionStrings = []string{
	"7",
	"5.500000",
	"null",
	"\"name\"",
	"true",
	"bob",
	"[bob,jay]",
	"!(bob)",
	"bob[0]",
	"bob + jay",
	"bob - jay",
	"bob * jay",
	"bob / jay",
	"bob || jay",
	"bob && jay",
	"bob < jay",
	"bob <= jay",
	"bob > jay",
	"bob >= jay",
	"bob == jay",
	"bob != jay",
	"bob ? jay : cat",
}

func TestExpressionsToString(t *testing.T) {

	for i, expr := range testExpressions {
		exprToString := expr.String()
		if exprToString != testExpressionStrings[i] {
			t.Errorf("Expected expression to evalute to %v, got: %v", testExpressionStrings[i], exprToString)
		}
	}

	// manually test function call, couldn't be set up as literal
	f := NewFunction("sum")
	f.AddArguments(ExpressionList{NewProperty("bob"), NewProperty("jay")})
	fStr := f.String()
	if fStr != "__func__.sum.bob,jay" {
		t.Errorf("Expected expression to evalute to __func__.sum.bob,jay, got: %v", fStr)
	}

	// manually test object literally separately, can be either of 2 strings (order not guaranteed)
	ol := NewObjectLiteral(Object{"field": NewProperty("bob"), "field2": NewProperty("jay")})
	olStr := ol.String()
	if olStr != "{\"field\": bob, \"field2\": jay}" && olStr != "{\"field2\": jay, \"field\": bob}" {
		t.Errorf("Expected expression to evalute to {\"field\": bob, \"field2\": jay} OR {\"field2\": jay, \"field\": bob}, got: %v", olStr)
	}

}

var sortList = SortList{
	SortItem{Sort: NewProperty("abv"),
		Ascending: true},
	SortItem{Sort: NewProperty("ibu"),
		Ascending: false},
}

var sortListString = "abv ASC, ibu DESC"

func TestSortListsToString(t *testing.T) {
	slString := sortList.String()
	if slString != sortListString {
		t.Errorf("Expected sort list to evalute to %v, got: %v", sortListString, slString)
	}
}

var exprList = ExpressionList{
	NewProperty("bob"),
	NewProperty("jay"),
}

var exprListString = "bob, jay"

func TestExpressionListsToString(t *testing.T) {
	elString := exprList.String()
	if elString != exprListString {
		t.Errorf("Expected expression list to evalute to %v, got: %v", exprListString, elString)
	}
}

func TestSimpleProperties(t *testing.T) {
	sp := NewProperty("a")
	head := sp.Head()
	if head != "a" {
		t.Errorf("Expected property head to be a, got %v", head)
	}

	tail := sp.Tail()
	if tail != nil {
		t.Errorf("Expected property tail to nil got %v", tail)
	}

	if sp.HasSubProperty() != false {
		t.Errorf("Expected property a.b.c to not have a sub property")
	}

	if sp.IsReferencingDataSource("b") != false {
		t.Errorf("Expected property a to not refer to datasource b")
	}
}

func TestComplexProperties(t *testing.T) {
	cp := NewProperty("a.b.c")
	head := cp.Head()
	if head != "a" {
		t.Errorf("Expected property head to be a, got %v", head)
	}

	tail := cp.Tail()
	if tail.String() != "b.c" {
		t.Errorf("Expected property tail to be b.c, got %v", tail)
	}

	if cp.HasSubProperty() != true {
		t.Errorf("Expected property a.b.c to have a sub property")
	}

	if cp.IsReferencingDataSource("a") != true {
		t.Errorf("Expected property a.b.c to refer to datasource a")
	}
}

var pragmaQueries = []string{
	"PRAGMA debugTokens=true",
	"PRAGMA debugGrammar=true",
	"PRAGMA debugTokens=7",
	"PRAGMA debugGrammar=7",
	"PRAGMA debugTokens=false",
	"PRAGMA debugGrammar=false",
	"PRAGMA 7=debug",
}

func TestPragma(t *testing.T) {

	unqlParser := NewUnqlParser(false, false, true)

	for _, v := range pragmaQueries {
		_, err := unqlParser.Parse(v)
		if err != nil {
			t.Errorf("Valid Query Parse Failed: %v - %v", v, err)
		}
	}

}

var aggQueries = []string{
	"SELECT FROM beer-sample WHERE doc.abv > 5 GROUP BY doc.type",
}

func TestAggregates(t *testing.T) {

	unqlParser := NewUnqlParser(false, false, true)

	for _, v := range aggQueries {
		pq, err := unqlParser.Parse(v)
		if err != nil {
			t.Errorf("Valid Query Parse Failed: %v - %v", v, err)
		}
		if !pq.IsAggregateQuery() {
			t.Errorf("Expected query to be recognized as an aggregate")
		}
	}

}

var explainQueries = []string{
	"EXPLAIN SELECT FROM beer-sample WHERE doc.abv > 5 GROUP BY doc.type",
}

func TestExplainQueries(t *testing.T) {

	unqlParser := NewUnqlParser(false, false, true)

	for _, v := range explainQueries {
		pq, err := unqlParser.Parse(v)
		if err != nil {
			t.Errorf("Valid Query Parse Failed: %v - %v", v, err)
		}
		if !pq.IsExplainOnly() {
			t.Errorf("Expected query to be recognized as explain only")
		}
	}

}

var parsableButInvalidQueries = []string {
	"SELECT xyz.cat FROM beer-sample AS beer, beer-sample as brewery", //xyz is not valid datasource
	"SELECT FROM beer-sample AS beer, beer-sample as brewery WHERE xyz.abc > 5", //xyz is not valid datasource
	"SELECT FROM beer-sample AS beer, beer-sample as brewery GROUP BY xyz.cat", //xyz is not valid datasource
	"SELECT FROM beer-sample AS beer, beer-sample as brewery GROUP BY beer HAVING xyz.cat > 7", //xyz is not valid datasource
	"SELECT FROM beer-sample AS beer, beer-sample as brewery LIMIT xyz.cat", //xyz is not valid datasource
	"SELECT FROM beer-sample AS beer, beer-sample as brewery LIMIT 1 OFFSET xyz.cat", //xyz is not valid datasource
	"SELECT FROM beer-sample AS beer, beer-sample as brewery ORDER BY xyz.cat", //xyz is not valid datasource
	"SELECT WHERE abc > 5", // 0 from clauses (allowed by UNQL grammar)
	"SELECT FROM beer-sample, beer-sample", // ambiguous datasource names
	"SELECT FROM beer-sample AS bs OVER bs.beers AS bs", // ambiguous datasource in OVER
	//"SELECT xyz.cat FROM beer-sample AS beer, beer-sample AS breweries", // xyz invalid reference
}

func TestParsableButInvalidQueries(t *testing.T) {

	unqlParser := NewUnqlParser(false, false, true)

	for _, v := range parsableButInvalidQueries {
		pq, err := unqlParser.Parse(v)
		if err != nil {
			t.Errorf("Valid Query Parse Failed: %v - %v", v, err)
			continue
		}

		err = pq.Validate()
		if err == nil {
			t.Errorf("Expected validation error, got none")
		}
	}

}