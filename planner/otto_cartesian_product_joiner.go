package planner

import (
	"fmt"
	"github.com/mschoch/tuq/parser"
	"github.com/robertkrimen/otto"
	"log"
)

type OttoCartesianProductJoiner struct {
	LeftSource    PlanPipelineComponent
	RightSource   PlanPipelineComponent
	OutputChannel DocumentChannel
	leftDocs      []Document
	rightDocs     []Document
	Otto          *otto.Otto
	joinExpr      parser.Expression
}

func NewOttoCartesianProductJoiner() *OttoCartesianProductJoiner {
	return &OttoCartesianProductJoiner{
		OutputChannel: make(DocumentChannel),
		leftDocs:      make([]Document, 0),
		rightDocs:     make([]Document, 0),
		Otto:          otto.New()}
}

// PlanPipelineComponent interface

func (cpj *OttoCartesianProductJoiner) SetSource(s PlanPipelineComponent) {
	log.Fatalf("SetSource called on DataSource, use ")
}

func (cpj *OttoCartesianProductJoiner) GetSource() PlanPipelineComponent {
	return nil
}

func (cpj *OttoCartesianProductJoiner) GetDocumentChannel() DocumentChannel {
	return cpj.OutputChannel
}

func (cpj *OttoCartesianProductJoiner) Run() {

	defer close(cpj.OutputChannel)

	// tell our sources to start
	go cpj.LeftSource.Run()
	go cpj.RightSource.Run()

	// get our sources channel
	leftSourceChannel := cpj.LeftSource.GetDocumentChannel()
	rightSourceChannel := cpj.RightSource.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range leftSourceChannel {
		cpj.leftDocs = append(cpj.leftDocs, doc)
	}

	// read from the channel until its closed
	for doc := range rightSourceChannel {
		cpj.rightDocs = append(cpj.rightDocs, doc)
	}

	for _, l := range cpj.leftDocs {
		for _, r := range cpj.rightDocs {
			combined := combineDocs(l, r)

			putDocumentIntoEnvironment(cpj.Otto, combined)
			expr_result := evaluateExpressionInEnvironmentAsBoolean(cpj.Otto, cpj.joinExpr)
			if expr_result {
				cpj.OutputChannel <- combined
			}
			cleanupDocumentFromEnvironment(cpj.Otto, combined)
		}
	}

}

func (cpj *OttoCartesianProductJoiner) Explain() {

	defer close(cpj.OutputChannel)

	// tell our sources to explain
	go cpj.LeftSource.Explain()
	go cpj.RightSource.Explain()

	// get our sources channel
	leftSourceChannel := cpj.LeftSource.GetDocumentChannel()
	rightSourceChannel := cpj.RightSource.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range leftSourceChannel {
		cpj.leftDocs = append(cpj.leftDocs, doc)
	}

	// read from the channel until its closed
	for doc := range rightSourceChannel {
		cpj.rightDocs = append(cpj.rightDocs, doc)
	}

	for _, l := range cpj.leftDocs {
		for _, r := range cpj.rightDocs {

			thisStep := map[string]interface{}{
				"_type":     "JOIN",
				"impl":      "Otto Full Cartesian Product",
				"condition": fmt.Sprintf("%v", cpj.joinExpr),
				"left":      l,
				"right":     r}

			cpj.OutputChannel <- thisStep

		}
	}

}

func (cpj *OttoCartesianProductJoiner) Cancel() {

}

// Joiner interface

func (cpj *OttoCartesianProductJoiner) SetLeftSource(l PlanPipelineComponent) {
	cpj.LeftSource = l
}

func (cpj *OttoCartesianProductJoiner) GetLeftSource() PlanPipelineComponent {
	return cpj.LeftSource
}

func (cpj *OttoCartesianProductJoiner) SetRightSource(r PlanPipelineComponent) {
	cpj.RightSource = r
}

func (cpj *OttoCartesianProductJoiner) GetRightSource() PlanPipelineComponent {
	return cpj.RightSource
}

func (cpj *OttoCartesianProductJoiner) SetCondition(e parser.Expression) error {
	cpj.joinExpr = e
	return nil
}

func (cpj *OttoCartesianProductJoiner) GetCondition() parser.Expression {
	return cpj.joinExpr
}
