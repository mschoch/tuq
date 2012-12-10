package planner

import (
	"fmt"
	"github.com/mschoch/tuq/parser"
	"github.com/robertkrimen/otto"
)

type OttoOffsetter struct {
	Source        PlanPipelineComponent
	OutputChannel DocumentChannel
	Offset        parser.Expression
	Otto          *otto.Otto
}

func NewOttoOffsetter() *OttoOffsetter {
	return &OttoOffsetter{
		OutputChannel: make(DocumentChannel),
		Otto:          otto.New()}
}

// PlanPipelineComponent interface

func (oo *OttoOffsetter) SetSource(s PlanPipelineComponent) {
	oo.Source = s
}

func (oo *OttoOffsetter) GetSource() PlanPipelineComponent {
	return oo.Source
}

func (oo *OttoOffsetter) GetDocumentChannel() DocumentChannel {
	return oo.OutputChannel
}

func (oo *OttoOffsetter) Run() {
	defer close(oo.OutputChannel)

	offset := evaluateExpressionInEnvironmentAsInteger(oo.Otto, oo.Offset)

	// tell our source to start
	go oo.Source.Run()

	// get our sources channel
	sourceChannel := oo.Source.GetDocumentChannel()

	rowsSkipped := int64(1)
	// read from the channel until its closed
	for doc := range sourceChannel {
		// skip over the first Offset, then pass through
		if rowsSkipped <= offset {
			rowsSkipped += 1
		} else {
			oo.OutputChannel <- doc
		}
	}

}

func (oo *OttoOffsetter) Explain() {
	defer close(oo.OutputChannel)

	// tell our source to explain
	go oo.Source.Explain()

	// get our sources channel
	sourceChannel := oo.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {

		thisStep := map[string]interface{}{
			"_type":      "OFFSET",
			"impl":       "Otto",
			"expression": fmt.Sprintf("%v", oo.Offset),
			"source":     doc,
			"cost":       oo.Cost(),
			"totalCost":  oo.TotalCost()}

		oo.OutputChannel <- thisStep
	}

}

func (oo *OttoOffsetter) Cancel() {

}

// Offsetter Interface

func (oo *OttoOffsetter) SetOffset(l parser.Expression) {
	oo.Offset = l
}

func (oo *OttoOffsetter) GetOffset() parser.Expression {
	return oo.Offset
}

func (oo *OttoOffsetter) Cost() float64 {
	return 1000
}

func (oo *OttoOffsetter) TotalCost() float64 {
	return oo.Cost() + oo.Source.TotalCost()
}
