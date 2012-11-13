package planner

import (
	//"encoding/json"
	"github.com/mschoch/go-unql-couchbase/parser"
	"github.com/robertkrimen/otto"
	"log"
	"testing"
)

func TestCSVReadEmployees(t *testing.T) {

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	plan := Plan{}
	plan.Root = employeesCSV

	documentChannel := plan.Run()
	rows := 0
	for _ = range documentChannel {
		rows += 1
	}

	if rows != 10 {
		t.Errorf("Expected 10 rows, got %d", rows)
	}
}

func TestLimit3Employees(t *testing.T) {

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	limit3 := NewOttoLimitter()
	limit3.SetLimit(parser.NewIntegerLiteral(3))
	limit3.SetSource(employeesCSV)

	plan := Plan{}
	plan.Root = limit3

	documentChannel := plan.Run()
	rows := 0
	for _ = range documentChannel {
		rows += 1
	}

	if rows != 3 {
		t.Errorf("Expected 3 rows, got %d", rows)
	}
}

func TestOffset3Employees(t *testing.T) {

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	offset3 := NewOttoOffsetter()
	offset3.SetOffset(parser.NewIntegerLiteral(3))
	offset3.SetSource(employeesCSV)

	plan := Plan{}
	plan.Root = offset3

	documentChannel := plan.Run()
	rows := 0
	for _ = range documentChannel {
		rows += 1
	}

	if rows != 7 {
		t.Errorf("Expected 7 rows, got %d", rows)
	}
}

func TestOffsetAndLimitEmployees(t *testing.T) {

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	offset3 := NewOttoOffsetter()
	offset3.SetOffset(parser.NewIntegerLiteral(3))
	offset3.SetSource(employeesCSV)

	limit1 := NewOttoLimitter()
	limit1.SetLimit(parser.NewIntegerLiteral(1))
	limit1.SetSource(offset3)

	plan := Plan{}
	plan.Root = limit1

	documentChannel := plan.Run()
	rows := 0
	for _ = range documentChannel {
		rows += 1
	}

	if rows != 1 {
		t.Errorf("Expected 1 rows, got %d", rows)
	}
}

func TestOttoFilterEmployeesNamedMarty(t *testing.T) {

	name := parser.NewProperty("name")
	marty := parser.NewStringLiteral("Marty")
	eq := parser.NewEqualsExpression(name, marty)

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	ottoFilter := NewOttoFilter()
	ottoFilter.SetSource(employeesCSV)
	ottoFilter.SetFilter(eq)

	plan := Plan{}
	plan.Root = ottoFilter

	documentChannel := plan.Run()

	rows := 0
	for doc := range documentChannel {
		rows += 1
		if doc["name"] != "Marty" {
			t.Errorf("Expected name Marty, got %v", doc["name"])
		}
	}

	if rows != 1 {
		t.Errorf("Expected 1 rows, got %d", rows)
	}

}

func TestOttoFilterEmployeesOver30(t *testing.T) {

	name := parser.NewProperty("age")
	thirty := parser.NewIntegerLiteral(30)
	gt := parser.NewGreaterThanExpression(name, thirty)

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	ottoFilter := NewOttoFilter()
	ottoFilter.SetSource(employeesCSV)
	ottoFilter.SetFilter(gt)

	plan := Plan{}
	plan.Root = ottoFilter

	documentChannel := plan.Run()

	rows := 0
	for _ = range documentChannel {
		rows += 1
	}

	if rows != 8 {
		t.Errorf("Expected 8 rows, got %d", rows)
	}

}

func TestOttoOrdererEmployeesByName(t *testing.T) {

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	sortExpression := parser.NewProperty("name")
	sortItem := parser.NewSortItem(sortExpression, true)
	sortList := parser.SortList{*sortItem}

	ottoOrderer := NewOttoOrderer()
	ottoOrderer.SetSource(employeesCSV)
	ottoOrderer.SetOrderBy(sortList)

	plan := Plan{}
	plan.Root = ottoOrderer

	documentChannel := plan.Run()

	row := 0
	for doc := range documentChannel {
		if row == 0 && doc["name"] != "Adam" {
			t.Errorf("Expected row 0 name Adam, got %v", doc["name"])
		} else if row == 4 && doc["name"] != "Jamie" {
			t.Errorf("Expected row 4 name Jamie, got %v", doc["name"])
		}
		row += 1
	}
}

func TestOttoOrdererEmployeesByAgeDescending(t *testing.T) {

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	sortExpression := parser.NewProperty("age")
	sortItem := parser.NewSortItem(sortExpression, false)
	sortList := parser.SortList{*sortItem}

	ottoOrderer := NewOttoOrderer()
	ottoOrderer.SetSource(employeesCSV)
	ottoOrderer.SetOrderBy(sortList)

	plan := Plan{}
	plan.Root = ottoOrderer

	documentChannel := plan.Run()

	row := 0
	for doc := range documentChannel {
		if row == 0 && doc["age"].(int64) != 64 {
			t.Errorf("Expected row 0 age 64, got %v", doc["age"])
		} else if row == 9 && doc["age"].(int64) != 2 {
			t.Errorf("Expected row 9 age 2, got %v", doc["age"])
		}
		row += 1
	}
}

func TestOttoOrdererEmployeesByAgeAscAndNameDesc(t *testing.T) {

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	sortExpression := parser.NewProperty("age")
	sortItem := parser.NewSortItem(sortExpression, true)

	sortExpression2 := parser.NewProperty("name")
	sortItem2 := parser.NewSortItem(sortExpression2, false)

	sortList := parser.SortList{*sortItem, *sortItem2}

	ottoOrderer := NewOttoOrderer()
	ottoOrderer.SetSource(employeesCSV)
	ottoOrderer.SetOrderBy(sortList)

	plan := Plan{}
	plan.Root = ottoOrderer

	documentChannel := plan.Run()

	row := 0
	for doc := range documentChannel {
		if row == 3 && doc["name"] != "Marty" {
			t.Errorf("Expected row 3 name Marty, got %v", doc["name"])
		} else if row == 4 && doc["name"] != "Dan" {
			t.Errorf("Expected row 4 name Dan, got %v", doc["name"])
		}
		row += 1
	}
}

func TestOttoOrdererEmployeesByAgeAscAndNameAsc(t *testing.T) {

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	sortExpression := parser.NewProperty("age")
	sortItem := parser.NewSortItem(sortExpression, true)

	sortExpression2 := parser.NewProperty("name")
	sortItem2 := parser.NewSortItem(sortExpression2, true)

	sortList := parser.SortList{*sortItem, *sortItem2}

	ottoOrderer := NewOttoOrderer()
	ottoOrderer.SetSource(employeesCSV)
	ottoOrderer.SetOrderBy(sortList)

	plan := Plan{}
	plan.Root = ottoOrderer

	documentChannel := plan.Run()

	row := 0
	for doc := range documentChannel {
		if row == 3 && doc["name"] != "Marty" {
			t.Errorf("Expected row 3 name Marty, got %v", doc["name"])
		} else if row == 2 && doc["name"] != "Dan" {
			t.Errorf("Expected row 2 name Dan, got %v", doc["name"])
		}
		row += 1
	}
}

func TestOttoGrouperByAge(t *testing.T) {

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	groupExpression := parser.NewProperty("age")
	//groupExpression2 := parser.NewProperty("name")
	//expList := parser.ExpressionList{groupExpression, groupExpression2}
	expList := parser.ExpressionList{groupExpression}

	ottoGrouper := NewOttoGrouper()
	ottoGrouper.SetGroupBy(expList)
	ottoGrouper.SetSource(employeesCSV)

	plan := Plan{}
	plan.Root = ottoGrouper

	documentChannel := plan.Run()

	for doc := range documentChannel {
		if doc["age"].(float64) == 34 {
			stats := doc["stats"].(ExpressionStatsMap)
			age := stats["age"]
			if age.Count != 3 {
				t.Errorf("Expected 3 people aged 34, got %d", age.Count)
			}
		}
	}

}

func TestOttoGrouperByCity(t *testing.T) {

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	groupExpression := parser.NewProperty("city")
	//groupExpression2 := parser.NewProperty("name")
	//expList := parser.ExpressionList{groupExpression, groupExpression2}
	expList := parser.ExpressionList{groupExpression}

	ottoGrouper := NewOttoGrouper()
	ottoGrouper.SetGroupBy(expList)
	ottoGrouper.SetSource(employeesCSV)

	plan := Plan{}
	plan.Root = ottoGrouper

	documentChannel := plan.Run()

	for doc := range documentChannel {
		if doc["city"] == "San Francisco" {
			stats := doc["stats"].(ExpressionStatsMap)
			city := stats["city"]
			if city.Count != 4 {
				t.Errorf("Expected 4 people from San Francisco, got %d", city.Count)
			}
		}
	}

}

func TestOttoGrouperByCityAndAge(t *testing.T) {

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")

	groupExpression := parser.NewProperty("city")
	groupExpression2 := parser.NewProperty("age")
	expList := parser.ExpressionList{groupExpression, groupExpression2}

	ottoGrouper := NewOttoGrouper()
	ottoGrouper.SetGroupBy(expList)
	ottoGrouper.SetSource(employeesCSV)

	plan := Plan{}
	plan.Root = ottoGrouper

	documentChannel := plan.Run()

	for doc := range documentChannel {
		if doc["city"] == "San Francisco" && doc["age"].(float64) == 39 {
			stats := doc["stats"].(ExpressionStatsMap)
			city := stats["city"]
			if city.Count != 2 {
				t.Errorf("Expected 2 people from San Francisco, got %d", city.Count)
			}
		}
	}

}

//func TestCPJoinerEmployeesAndDepartments(t *testing.T) {
//
//	employeesCSV := NewCSVDataSource()
//	employeesCSV.SetName("test_csv_datasources/employees.csv")
//	employeesCSV.SetAs("emp")
//
//	departmentsCSV := NewCSVDataSource()
//	departmentsCSV.SetName("test_csv_datasources/departments.csv")
//	departmentsCSV.SetAs("dept")
//
//	joiner := NewCartesianProductJoiner()
//	joiner.SetLeftSource(employeesCSV)
//	joiner.SetRightSource(departmentsCSV)
//
//	plan := Plan{}
//	plan.Root = joiner
//
//	documentChannel := plan.Run()
//
//	row := 0
//	for _ = range documentChannel {
//		row += 1
//	}
//
//	if row != 50 {
//		t.Errorf("Expected cartesian product to contain 50 rows, got %d", row)
//	}
//}

func TestCPJoinerEmployeesAndDepartmentsWithWhereClause(t *testing.T) {
	log.Printf("testing join")
	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")
	employeesCSV.SetAs("emp")

	departmentsCSV := NewCSVDataSource()
	departmentsCSV.SetName("test_csv_datasources/departments.csv")
	departmentsCSV.SetAs("dept")

	joiner := NewCartesianProductJoiner()
	joiner.SetLeftSource(employeesCSV)
	joiner.SetRightSource(departmentsCSV)

	name := parser.NewProperty("emp.name")
	marty := parser.NewStringLiteral("Marty")
	eq := parser.NewEqualsExpression(name, marty)

	ottoFilter := NewOttoFilter()
	ottoFilter.SetSource(joiner)
	ottoFilter.SetFilter(eq)

	plan := Plan{}
	plan.Root = ottoFilter

	documentChannel := plan.Run()

	row := 0
	for _ = range documentChannel {
		row += 1
	}

	if row != 5 {
		t.Errorf("Expected cartesian product to contain 5 rows, got %d", row)
	}
}

func TestCPJoinerEmployeesAndDepartmentsWithWhereDepartmentIdMatch(t *testing.T) {
	log.Printf("testing join")

	employeesCSV := NewCSVDataSource()
	employeesCSV.SetName("test_csv_datasources/employees.csv")
	employeesCSV.SetAs("emp")

	departmentsCSV := NewCSVDataSource()
	departmentsCSV.SetName("test_csv_datasources/departments.csv")
	departmentsCSV.SetAs("dept")

	joiner := NewCartesianProductJoiner()
	joiner.SetLeftSource(employeesCSV)
	joiner.SetRightSource(departmentsCSV)

	name := parser.NewProperty("emp.department_id")
	marty := parser.NewProperty("dept.department_id")
	eq := parser.NewEqualsExpression(name, marty)

	ottoFilter := NewOttoFilter()
	ottoFilter.SetSource(joiner)
	ottoFilter.SetFilter(eq)

	plan := Plan{}
	plan.Root = ottoFilter

	documentChannel := plan.Run()

	row := 0
	for doc := range documentChannel {
		log.Printf("%v", doc)
		row += 1
	}

	if row != 10 {
		t.Errorf("Expected cartesian product to contain 10 rows, got %d", row)
	}
}

func TestOttoExport(t *testing.T) {
	o := otto.New()

	//o.Run("x = {'bob': 'cat'}")
	//x, err := o.Get("x")
	_, err := o.Run("y = {'string': 'cat', 'number': 7.4, 'bool': true}")
	if err != nil {
		log.Printf("run y error %v", err)
	}

	x, err := o.Run("x = {'string': 'cat', 'number': 7.4, 'bool': true, 'object': y}")
	if err != nil {
		log.Printf("run x error %v", err)
	}

	asmap, err := x.Export()
	if err != nil {
		log.Printf("export error %v", err)
	}

	log.Printf("%#v", asmap)

//	err = o.Set("z", asmap)
//	if err != nil {
//		log.Printf("Set error %v", err)
//	}
}
