package planner

import (
	"fmt"
	"github.com/mschoch/tuq/parser"
	"github.com/robertkrimen/otto"
	//"log"
)

type OttoFilter struct {
	Source        PlanPipelineComponent
	OutputChannel DocumentChannel
	expr          parser.Expression
	Otto          *otto.Otto
}

func NewOttoFilter() *OttoFilter {
	return &OttoFilter{
		OutputChannel: make(DocumentChannel),
		Otto:          otto.New()}
}

// PlanPipelineComponent interface

func (of *OttoFilter) SetSource(s PlanPipelineComponent) {
	of.Source = s
}

func (of *OttoFilter) GetSource() PlanPipelineComponent {
	return of.Source
}

func (of *OttoFilter) GetDocumentChannel() DocumentChannel {
	return of.OutputChannel
}

func (of *OttoFilter) Run() {

	defer close(of.OutputChannel)

	// tell our source to start
	go of.Source.Run()

	// get our sources channel
	sourceChannel := of.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {
		putDocumentIntoEnvironment(of.Otto, doc)
		expr_result := evaluateExpressionInEnvironmentAsBoolean(of.Otto, of.expr)
		if expr_result {
			of.OutputChannel <- doc
		}
		cleanupDocumentFromEnvironment(of.Otto, doc)

	}

}

func (of *OttoFilter) Explain() {

	defer close(of.OutputChannel)

	// tell our source to start
	go of.Source.Explain()

	// get our sources channel
	sourceChannel := of.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {

		thisStep := map[string]interface{}{
			"_type":      "WHERE",
			"impl":       "Otto",
			"expression": fmt.Sprintf("%v", of.expr),
			"source":     doc}

		of.OutputChannel <- thisStep
	}

}

func (of *OttoFilter) Cancel() {

}

// Filter Interface

func (of *OttoFilter) SetFilter(e parser.Expression) {
	of.expr = e
}

func (of *OttoFilter) GetFilter() parser.Expression {
	return of.expr
}
