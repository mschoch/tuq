package csv

import (
	"fmt"
	"github.com/mschoch/tuq/datasources"
	"github.com/mschoch/tuq/parser"
	"github.com/mschoch/tuq/planner"
	"log"
)

func init() {
	datasources.RegisterDataSourceImpl("datasources", NewDataSourcesDataSource)
}

type DataSourcesDataSource struct {
	Name          string
	As            string
	OutputChannel planner.DocumentChannel
	cancelled     bool
}

func NewDataSourcesDataSource(config map[string]interface{}) planner.DataSource {
	return &DataSourcesDataSource{
		OutputChannel: make(planner.DocumentChannel),
		cancelled:     false}
}

// PlanPipelineComponent interface

func (ds *DataSourcesDataSource) SetSource(source planner.PlanPipelineComponent) {
	log.Fatalf("SetSource called on DataSource")
}

func (ds *DataSourcesDataSource) GetSource() planner.PlanPipelineComponent {
	return nil
}

func (ds *DataSourcesDataSource) GetDocumentChannel() planner.DocumentChannel {
	return ds.OutputChannel
}

func (ds *DataSourcesDataSource) Run() {
	defer close(ds.OutputChannel)

	for k, v := range datasources.DataSources {
	    doc := planner.Document{"name": k, "definition": v}
	    doc = datasources.WrapDocWithDatasourceAs(ds.As, doc)
		ds.OutputChannel <- doc
	}
}

func (ds *DataSourcesDataSource) Explain() {
	defer close(ds.OutputChannel)

	thisStep := map[string]interface{}{
		"_type": "FROM",
		"impl":  "DataSources",
		"name":  ds.Name,
		"as":    ds.As}

	ds.OutputChannel <- thisStep
}

func (ds *DataSourcesDataSource) Cancel() {
	ds.cancelled = true
}

// DataSource Interface

func (ds *DataSourcesDataSource) SetName(name string) {
	ds.Name = name
}

func (ds *DataSourcesDataSource) SetAs(as string) {
	ds.As = as
}

func (ds *DataSourcesDataSource) GetAs() string {
	return ds.As
}

func (ds *DataSourcesDataSource) SetFilter(filter parser.Expression) error {
	return fmt.Errorf("DataSources DataSource does not support filter")
}

func (ds *DataSourcesDataSource) GetFilter() parser.Expression {
	return nil
}

func (ds *DataSourcesDataSource) SetOrderBy(sortlist parser.SortList) error {
	return fmt.Errorf("DataSources DataSource does not support order by")
}

func (ds *DataSourcesDataSource) GetOrderBy() parser.SortList {
	return nil
}

func (ds *DataSourcesDataSource) SetLimit(e parser.Expression) error {
	return fmt.Errorf("DataSources DataSource does not support limit")
}

func (ds *DataSourcesDataSource) SetOffset(e parser.Expression) error {
	return fmt.Errorf("DataSources DataSource does not support offset")
}

func (ds *DataSourcesDataSource) SetGroupByWithStatsFields(groupby parser.ExpressionList, stats_fields []string) error {
	return fmt.Errorf("DataSources DataSource does not support group by")
}

func (ds *DataSourcesDataSource) SetHaving(having parser.Expression) error {
	return fmt.Errorf("DataSources DataSource does not support having")
}

func (ds *DataSourcesDataSource) DocsFromIds(docIds []string) ([]planner.Document, error) {
	panic("Unexpected call to DataSources DataSource to get DocsFromIds")
}
