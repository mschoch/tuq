package planner

import (
	"fmt"
	"github.com/mschoch/tuq/parser"
	"github.com/robertkrimen/otto"
	"log"
)

type OttoSelecter struct {
	Source        PlanPipelineComponent
	OutputChannel RowChannel
	expr          parser.Expression
	Otto          *otto.Otto
}

func NewOttoSelecter() *OttoSelecter {
	return &OttoSelecter{
		OutputChannel: make(RowChannel),
		Otto:          otto.New()}
}

// Selecter interface

func (os *OttoSelecter) SetSource(s PlanPipelineComponent) {
	os.Source = s
}

func (os *OttoSelecter) GetSource() PlanPipelineComponent {
	return os.Source
}

func (os *OttoSelecter) GetRowChannel() RowChannel {
	return os.OutputChannel
}

func (os *OttoSelecter) Run() {

	defer close(os.OutputChannel)

	// tell our source to start
	go os.Source.Run()

	// get our sources channel
	sourceChannel := os.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {
		putDocumentIntoEnvironment(os.Otto, doc)
		expr_result := evaluateExpressionInEnvironment(os.Otto, os.expr)
		expr_exported, err := expr_result.Export()
		if err != nil {
			log.Printf("Error exporting evaluated expression %v", err)
		}
		if expr_exported != nil {
			os.OutputChannel <- expr_exported
		}
		cleanupDocumentFromEnvironment(os.Otto, doc)

	}

}

func (os *OttoSelecter) Explain() {

	defer close(os.OutputChannel)

	// tell our source to start explaining
	go os.Source.Explain()

	// get our sources channel
	sourceChannel := os.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {

		thisStep := map[string]interface{}{
			"_type":      "SELECT",
			"impl":       "Otto",
			"expression": fmt.Sprintf("%v", os.expr),
			"source":     doc}

		os.OutputChannel <- thisStep
	}

}

func (os *OttoSelecter) Cancel() {

}

func (os *OttoSelecter) SetSelect(e parser.Expression) {
	os.expr = e
}
