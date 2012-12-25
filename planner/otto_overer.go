package planner

import (
	"fmt"
	"github.com/mschoch/tuq/parser"
	"github.com/robertkrimen/otto"
	"log"
)

type OttoOver struct {
	Source        PlanPipelineComponent
	OutputChannel DocumentChannel
	Path          *parser.Property
	As            string
	Otto          *otto.Otto
}

func NewOttoOver() *OttoOver {
	return &OttoOver{
		OutputChannel: make(DocumentChannel),
		Otto:          otto.New()}
}

// PlanPipelineComponent interface

func (of *OttoOver) SetSource(s PlanPipelineComponent) {
	of.Source = s
}

func (of *OttoOver) GetSource() PlanPipelineComponent {
	return of.Source
}

func (of *OttoOver) GetDocumentChannel() DocumentChannel {
	return of.OutputChannel
}

func (of *OttoOver) Run() {

	defer close(of.OutputChannel)

	// tell our source to start
	go of.Source.Run()

	// get our sources channel
	sourceChannel := of.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {
		putDocumentIntoEnvironment(of.Otto, doc)

		expr_result := evaluateExpressionInEnvironment(of.Otto, of.Path)
		expr_exported, err := expr_result.Export()
		if err != nil {
			if err.Error() != "undefined" {
				log.Printf("Error exporting evaluated expression %v", err)
				continue
			}
		}
		if expr_exported != nil {
			switch val := expr_exported.(type) {
			case []interface{}:
				for _, v := range val {
					newdoc := make(Document)
					// start with copy of existing doc
					for dk, dv := range doc {
						newdoc[dk] = dv
					}
					newdoc[of.As] = v
					of.OutputChannel <- newdoc
				}
			}

		}

		cleanupDocumentFromEnvironment(of.Otto, doc)

	}

}

func (of *OttoOver) Explain() {

	defer close(of.OutputChannel)

	// tell our source to start
	go of.Source.Explain()

	// get our sources channel
	sourceChannel := of.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {

		thisStep := map[string]interface{}{
			"_type":     "OVER",
			"impl":      "Otto",
			"path":      fmt.Sprintf("%v", of.Path),
			"as":        fmt.Sprintf("%v", of.As),
			"source":    doc,
			"cost":      of.Cost(),
			"totalCost": of.TotalCost()}

		of.OutputChannel <- thisStep
	}

}

func (of *OttoOver) Cancel() {

}

func (of *OttoOver) Cost() float64 {
	return 1000
}

func (of *OttoOver) TotalCost() float64 {
	return of.Cost() + of.Source.TotalCost()
}

func (of *OttoOver) SetPath(path *parser.Property) {
	of.Path = path
}

func (of *OttoOver) SetAs(as string) {
	of.As = as
}
