package planner

import (
	"fmt"
	"github.com/mschoch/go-unql-couchbase/parser"
	"github.com/robertkrimen/otto"
)

type OttoLimitter struct {
	Source        PlanPipelineComponent
	OutputChannel DocumentChannel
	Limit         parser.Expression
	Otto          *otto.Otto
}

func NewOttoLimitter() *OttoLimitter {
	return &OttoLimitter{
		OutputChannel: make(DocumentChannel),
		Otto:          otto.New()}
}

// PlanPipelineComponent interface

func (ol *OttoLimitter) SetSource(s PlanPipelineComponent) {
	ol.Source = s
}

func (ol *OttoLimitter) GetSource() PlanPipelineComponent {
	return ol.Source
}

func (ol *OttoLimitter) GetDocumentChannel() DocumentChannel {
	return ol.OutputChannel
}

func (ol *OttoLimitter) Run() {

	limit := evaluateExpressionInEnvironmentAsInteger(ol.Otto, ol.Limit)

	// tell our source to start
	go ol.Source.Run()

	// get our sources channel
	sourceChannel := ol.Source.GetDocumentChannel()

	rowsWritten := int64(1)
	// read from the channel until its closed
	for doc := range sourceChannel {
		// pass through until we reach the Limit
		if rowsWritten < limit {
			ol.OutputChannel <- doc
			rowsWritten += 1
		} else if rowsWritten == limit {
			ol.OutputChannel <- doc
			close(ol.OutputChannel)
			rowsWritten += 1
			ol.Source.Cancel()
		}
	}

}

func (ol *OttoLimitter) Explain() {

	defer close(ol.OutputChannel)

	// tell our source to explain
	go ol.Source.Explain()

	// get our sources channel
	sourceChannel := ol.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {

		thisStep := map[string]interface{}{
			"_type":      "LIMIT",
			"impl":       "Otto",
			"expression": fmt.Sprintf("%v", ol.Limit),
			"source":     doc}

		ol.OutputChannel <- thisStep
	}

}

func (ol *OttoLimitter) Cancel() {

}

// Limitter interface

func (ol *OttoLimitter) SetLimit(l parser.Expression) {
	ol.Limit = l
}

func (ol *OttoLimitter) GetLimit() parser.Expression {
	return ol.Limit
}
