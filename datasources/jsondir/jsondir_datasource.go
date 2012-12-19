package jsondir

import (
	"encoding/json"
	"fmt"
	"github.com/mschoch/tuq/datasources"
	"github.com/mschoch/tuq/parser"
	"github.com/mschoch/tuq/planner"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func init() {
	datasources.RegisterDataSourceImpl("jsondir", NewJsonDirDataSource)
}

type JsonDirDataSource struct {
	Name          string
	As            string
	OutputChannel planner.DocumentChannel
	cancelled     bool
	dirname       string
}

func NewJsonDirDataSource(config map[string]interface{}) planner.DataSource {
	return &JsonDirDataSource{
		dirname:       config["path"].(string),
		OutputChannel: make(planner.DocumentChannel),
		cancelled:     false}
}

// PlanPipelineComponent interface

func (ds *JsonDirDataSource) SetSource(source planner.PlanPipelineComponent) {
	log.Fatalf("SetSource called on DataSource")
}

func (ds *JsonDirDataSource) GetSource() planner.PlanPipelineComponent {
	return nil
}

func (ds *JsonDirDataSource) GetDocumentChannel() planner.DocumentChannel {
	return ds.OutputChannel
}

func (ds *JsonDirDataSource) Run() {
	defer close(ds.OutputChannel)

	files, err := ioutil.ReadDir(ds.dirname)
	if err != nil {
		log.Printf("Error reading files in directory: %v", err)
		return
	}

	for _, file := range files {

		data, err := ioutil.ReadFile(ds.dirname + "/" + file.Name())
		if err != nil {
			log.Printf("Error reading file: %v", err)
			continue
		}

		var parsedJson map[string]interface{}
		err = json.Unmarshal(data, &parsedJson)
		if err != nil {
			log.Printf("Error parsing json: %v", err)
			continue
		}

		id := file.Name()
		if strings.HasSuffix(id, ".json") {
			dotjsonindex := strings.LastIndex(id, ".json")
			id = id[0:dotjsonindex]
		}

		result := planner.Document{
			"doc": parsedJson,
			"meta": map[string]interface{}{
				"id": id}}

		result = datasources.WrapDocWithDatasourceAs(ds.As, result)
		ds.OutputChannel <- result

	}
}

func (ds *JsonDirDataSource) Explain() {
	defer close(ds.OutputChannel)

	thisStep := map[string]interface{}{
		"_type":     "FROM",
		"impl":      "JSON Directory",
		"direcotry": ds.dirname,
		"name":      ds.Name,
		"as":        ds.As,
		"cost":      ds.Cost(),
		"totalCost": ds.TotalCost()}

	ds.OutputChannel <- thisStep
}

func (ds *JsonDirDataSource) Cancel() {
	ds.cancelled = true
}

// DataSource Interface

func (ds *JsonDirDataSource) SetName(name string) {
	ds.Name = name
}

func (ds *JsonDirDataSource) SetAs(as string) {
	ds.As = as
}

func (ds *JsonDirDataSource) GetAs() string {
	return ds.As
}

func (ds *JsonDirDataSource) SetFilter(filter parser.Expression) error {
	return fmt.Errorf("JsonDir DataSource does not support filter")
}

func (ds *JsonDirDataSource) GetFilter() parser.Expression {
	return nil
}

func (ds *JsonDirDataSource) SetOrderBy(sortlist parser.SortList) error {
	return fmt.Errorf("JsonDir DataSource does not support order by")
}

func (ds *JsonDirDataSource) GetOrderBy() parser.SortList {
	return nil
}

func (ds *JsonDirDataSource) SetLimit(e parser.Expression) error {
	return fmt.Errorf("JsonDir DataSource does not support limit")
}

func (ds *JsonDirDataSource) SetOffset(e parser.Expression) error {
	return fmt.Errorf("JsonDir DataSource does not support offset")
}

func (ds *JsonDirDataSource) SetGroupByWithStatsFields(groupby parser.ExpressionList, stats_fields []string) error {
	return fmt.Errorf("JsonDir DataSource does not support group by")
}

func (ds *JsonDirDataSource) SetHaving(having parser.Expression) error {
	return fmt.Errorf("JsonDir DataSource does not support having")
}

func (ds *JsonDirDataSource) DocsFromIds(docIds []string) ([]planner.Document, error) {
	panic("Unexpected call to JsonDir DataSource to get DocsFromIds")
}

func (ds *JsonDirDataSource) Cost() float64 {
	return math.Inf(1)
}

func (ds *JsonDirDataSource) TotalCost() float64 {
	return ds.Cost()
}
