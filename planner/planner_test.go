package planner

import (
	//"encoding/json"
	"github.com/mschoch/tuq/parser"
	//"github.com/robertkrimen/otto"
	//"log"
	"testing"
)

type StubDataSource struct {
	Name          string
	As            string
	OutputChannel DocumentChannel
	cancelled     bool
}

func (ds *StubDataSource) Cancel() {
	return
}

func (ds *StubDataSource) Cost() float64 {
	return 10.0
}

func (ds *StubDataSource) TotalCost() float64 {
	return ds.Cost()
}

func (ds *StubDataSource) GetDocumentChannel() DocumentChannel {
	return ds.OutputChannel
}

func (ds *StubDataSource) GetSource() PlanPipelineComponent {
	return nil
}

func (ds *StubDataSource) SetSource(source PlanPipelineComponent) {
	return
}

func (ds *StubDataSource) Explain() {
	defer close(ds.OutputChannel)

	thisStep := map[string]interface{}{
		"_type":     "FROM",
		"impl":      "Stub",
		"name":      ds.Name,
		"as":        ds.As,
		"cost":      ds.Cost(),
		"totalCost": ds.TotalCost()}

	ds.OutputChannel <- thisStep
}

func (ds *StubDataSource) Run() {
	defer close(ds.OutputChannel)

	docs := []Document{
		Document{
			"name":  "Marty",
			"age":   34,
			"type":  "employee",
			"roles": []string{"it", "qa"}},
		Document{
			"name":  "Tom",
			"age":   29,
			"type":  "employee",
			"roles": []string{"it", "sales"}},
		Document{
			"name":  "Rajiv",
			"age":   44,
			"type":  "employee",
			"roles": []string{"it", "engineering"}},
		Document{
			"name":  "Chris",
			"age":   38,
			"type":  "employee",
			"roles": []string{"it", "engineering"}},
		Document{
			"name":  "Amit",
			"age":   22,
			"type":  "contract",
			"roles": []string{"marketing", "sales"}},
		Document{
			"name":  "Divya",
			"age":   30,
			"type":  "contract",
			"roles": []string{"it", "sales"}},
	}

	for _, doc := range docs {
		ds.OutputChannel <- Document{ds.As: doc}
	}

}

func NewStubDataSource(name, as string) *StubDataSource {
	return &StubDataSource{
		OutputChannel: make(DocumentChannel),
		cancelled:     false,
		Name:          name,
		As:            as}
}

func TestFilter(t *testing.T) {

	ds := NewStubDataSource("stub", "stub")

	f := NewOttoFilter()
	f.SetSource(ds)
	f.SetFilter(parser.NewGreaterThanExpression(parser.NewProperty("stub.age"), parser.NewIntegerLiteral(30)))

	go f.Run()

	output := f.GetDocumentChannel()

	for doc := range output {
		if doc["stub"].(Document)["age"].(int) <= 30 {
			t.Errorf("Row does not match expected filter output")
		}

	}
}

func TestGrouper(t *testing.T) {
	ds := NewStubDataSource("stub", "stub")

	g := NewOttoGrouper()
	g.SetSource(ds)
	g.SetGroupByWithStatsFields(parser.ExpressionList{parser.NewProperty("stub.type")}, []string{})

	go g.Run()

	output := g.GetDocumentChannel()

	for doc := range output {
		group := doc["stub"].(Document)["type"].(string)
		count := doc["__func__"].(Document)["count"].(Document)["stub"].(Document)["type"].(int)
		if group == "contract" && count != 2 {
			t.Errorf("Expected 2 contract workers, got %v", count)
		}
		if group == "employee" && count != 4 {
			t.Errorf("Expected 4 contract workers, got %v", count)
		}
	}
}

func TestLimit(t *testing.T) {

	ds := NewStubDataSource("stub", "stub")

	l := NewOttoLimitter()
	l.SetSource(ds)
	l.SetLimit(parser.NewIntegerLiteral(3))

	go l.Run()

	output := l.GetDocumentChannel()

	rows := 0
	for _ = range output {
		rows += 1
	}

	if rows != 3 {
		t.Errorf("Expected only 3 rows, got %v", rows)
	}
}

func TestOffset(t *testing.T) {

	ds := NewStubDataSource("stub", "stub")

	l := NewOttoOffsetter()
	l.SetSource(ds)
	l.SetOffset(parser.NewIntegerLiteral(3))

	go l.Run()

	output := l.GetDocumentChannel()

	rows := 0
	for _ = range output {
		rows += 1
	}

	if rows != 3 {
		t.Errorf("Expected only 3 rows, got %v", rows)
	}
}

func TestOrder(t *testing.T) {

	ds := NewStubDataSource("stub", "stub")

	o := NewOttoOrderer()
	o.SetSource(ds)
	o.SetOrderBy(parser.SortList{parser.SortItem{Sort: parser.NewProperty("stub.age"), Ascending: true}})

	go o.Run()

	output := o.GetDocumentChannel()

	lastAge := 0
	for doc := range output {
		age := doc["stub"].(Document)["age"].(int)
		if !(age > lastAge) {
			t.Errorf("Results not ordered correctly")
		}
		lastAge = age
	}
}

func TestOver(t *testing.T) {

	ds := NewStubDataSource("stub", "stub")

	o := NewOttoOver()
	o.SetSource(ds)
	o.SetPath(parser.NewProperty("stub.roles"))
	o.SetAs("role")

	go o.Run()

	output := o.GetDocumentChannel()

	rows := 0
	for doc := range output {
		rows += 1
		_, ok := doc["role"].(string)
		if !ok {
			t.Errorf("Row should have string role")
		}
	}

	if rows != 12 {
		t.Errorf("Expected 12 rows, got %v", rows)
	}
}

func TestOttoSelect(t *testing.T) {

	ds := NewStubDataSource("stub", "stub")

	s := NewOttoSelecter()
	s.SetSource(ds)
	s.SetSelect(parser.NewProperty("stub.age"))

	go s.Run()

	output := s.GetRowChannel()

	for row := range output {
		_, ok := row.(float64)
		if !ok {
			t.Errorf("Row should be float64")
		}
	}
}

func TestDefaultSelect(t *testing.T) {

	ds := NewStubDataSource("stub", "stub")

	s := NewDefaultSelecter()
	s.SetSource(ds)

	go s.Run()

	output := s.GetRowChannel()

	for doc := range output {
		_, ok := doc.(Document)["stub"]
		if !ok {
			t.Errorf("Default select should be map containing alias stub")
		}
	}
}

func TestCartesianProductJoiner(t *testing.T) {

	ds1 := NewStubDataSource("stub", "stub1")
	ds2 := NewStubDataSource("stub", "stub2")

	j := NewOttoCartesianProductJoiner()
	j.SetLeftSource(ds1)
	j.SetRightSource(ds2)
	j.SetCondition(parser.NewBoolLiteral(true))

	go j.Run()

	output := j.GetDocumentChannel()

	rows := 0
	for doc := range output {
		rows += 1
		_, lok := doc["stub1"]
		_, rok := doc["stub2"]
		if !lok || !rok {
			t.Errorf("Expected row to contain stub1 and stub2")
		}
	}

	if rows != 36 {
		t.Errorf("Expected 36 rows from full cartesian product, got %v", rows)
	}
}