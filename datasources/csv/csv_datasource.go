package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/mschoch/tuq/datasources"
	"github.com/mschoch/tuq/parser"
	"github.com/mschoch/tuq/planner"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func init() {
	datasources.RegisterDataSourceImpl("csv", NewCSVDataSource)
}

type CSVDataSource struct {
	Name          string
	As            string
	OutputChannel planner.DocumentChannel
	cancelled     bool
	filename      string
}

func NewCSVDataSource(config map[string]interface{}) planner.DataSource {
	return &CSVDataSource{
		filename:      config["path"].(string),
		OutputChannel: make(planner.DocumentChannel),
		cancelled:     false}
}

// PlanPipelineComponent interface

func (ds *CSVDataSource) SetSource(source planner.PlanPipelineComponent) {
	log.Fatalf("SetSource called on DataSource")
}

func (ds *CSVDataSource) GetSource() planner.PlanPipelineComponent {
	return nil
}

func (ds *CSVDataSource) GetDocumentChannel() planner.DocumentChannel {
	return ds.OutputChannel
}

func (ds *CSVDataSource) Run() {
	defer close(ds.OutputChannel)

	var sr io.Reader
	if strings.HasPrefix(ds.filename, "http") {
		resp, err := http.Get(ds.filename)
		if err != nil {
			log.Printf("Error:", err)
			return
		}
		sr = resp.Body
		defer resp.Body.Close()
	} else {
		log.Printf("opening file %v", ds.filename)
		file, err := os.Open(ds.filename)
		if err != nil {
			log.Printf("Error:", err)
			return
		}
		sr = file
		defer file.Close()
	}

	reader := csv.NewReader(sr)
	var columnHeaders []string
	row := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if row == 0 {
			columnHeaders = record
		} else {
			doc := make(planner.Document)
			for i, v := range record {
				// try to do some guesswork here about the datatype
				// we're not trying to be perfect, but we would like to support
				// string, int, float, bool as best we can
				if v == "true" {
					doc[columnHeaders[i]] = true
				} else if v == "false" {
					doc[columnHeaders[i]] = false
				} else {

					// now try int
					v_i, err := strconv.ParseInt(v, 0, 64)
					if err != nil {
						// now try float
						v_f, err := strconv.ParseFloat(v, 64)
						if err != nil {
							// leave it as string
							doc[columnHeaders[i]] = v
						} else {
							doc[columnHeaders[i]] = v_f
						}
					} else {
						doc[columnHeaders[i]] = v_i
					}
				}
			}
			if ds.As != "" {
				doccopy := doc
				doc = make(planner.Document)
				doc[ds.As] = doccopy
			}
			ds.OutputChannel <- doc
		}
		row += 1
	}
}

func (ds *CSVDataSource) Explain() {
	defer close(ds.OutputChannel)

	thisStep := map[string]interface{}{
		"_type":    "FROM",
		"impl":     "CSV File",
		"filename": ds.filename,
		"name":     ds.Name,
		"as":       ds.As}

	ds.OutputChannel <- thisStep
}

func (ds *CSVDataSource) Cancel() {
	ds.cancelled = true
}

// DataSource Interface

func (ds *CSVDataSource) SetName(name string) {
	ds.Name = name
}

func (ds *CSVDataSource) SetAs(as string) {
	ds.As = as
}

func (ds *CSVDataSource) GetAs() string {
	return ds.As
}

func (ds *CSVDataSource) SetFilter(filter parser.Expression) error {
	return fmt.Errorf("CSV DataSource does not support filter")
}

func (ds *CSVDataSource) GetFilter() parser.Expression {
	return nil
}

func (ds *CSVDataSource) SetOrderBy(sortlist parser.SortList) error {
	return fmt.Errorf("CSV DataSource does not support order by")
}

func (ds *CSVDataSource) GetOrderBy() parser.SortList {
	return nil
}

func (ds *CSVDataSource) SetLimit(e parser.Expression) error {
	return fmt.Errorf("CSV DataSource does not support limit")
}

func (ds *CSVDataSource) SetOffset(e parser.Expression) error {
	return fmt.Errorf("CSV DataSource does not support offset")
}

func (ds *CSVDataSource) SetGroupByWithStatsFields(groupby parser.ExpressionList, stats_fields []string) error {
	return fmt.Errorf("CSV DataSource does not support group by")
}

func (ds *CSVDataSource) SetHaving(having parser.Expression) error {
	return fmt.Errorf("CSV DataSource does not support having")
}

func (ds *CSVDataSource) DocsFromIds(docIds []string) ([]planner.Document, error) {
	panic("Unexpected call to CSV DataSource to get DocsFromIds")
}
