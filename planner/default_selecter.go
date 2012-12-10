package planner

import (
	"github.com/mschoch/tuq/parser"
	//"log"
)

type DefaultSelecter struct {
	Source        PlanPipelineComponent
	OutputChannel RowChannel
	expr          parser.Expression
}

func NewDefaultSelecter() *DefaultSelecter {
	return &DefaultSelecter{
		OutputChannel: make(RowChannel)}
}

// Selecter interface

func (ds *DefaultSelecter) SetSource(s PlanPipelineComponent) {
	ds.Source = s
}

func (ds *DefaultSelecter) GetSource() PlanPipelineComponent {
	return ds.Source
}

func (ds *DefaultSelecter) GetRowChannel() RowChannel {
	return ds.OutputChannel
}

func (ds *DefaultSelecter) Run() {

	defer close(ds.OutputChannel)

	// tell our source to start
	go ds.Source.Run()

	// get our sources channel
	sourceChannel := ds.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {

		ds.OutputChannel <- doc

	}
}

func (ds *DefaultSelecter) Explain() {

	defer close(ds.OutputChannel)

	// tell our source to start explaining
	go ds.Source.Explain()

	// get our sources channel
	sourceChannel := ds.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {

		thisStep := map[string]interface{}{
			"_type":      "SELECT",
			"impl":       "Default",
			"expression": "FULL DOCUMENTS",
			"source":     doc,
			"cost":       ds.Cost(),
			"totalCost":  ds.TotalCost()}

		ds.OutputChannel <- thisStep
	}
}

func (ds *DefaultSelecter) Cancel() {

}

func (ds *DefaultSelecter) SetSelect(e parser.Expression) {

}

func (ds *DefaultSelecter) Cost() float64 {
	return 1000
}

func (ds *DefaultSelecter) TotalCost() float64 {
	return ds.Cost() + ds.Source.TotalCost()
}
