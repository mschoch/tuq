package planner

import (
	"fmt"
	"github.com/mschoch/tuq/parser"
	"github.com/robertkrimen/otto"
	"log"
)

type OttoSortMergeJoiner struct {
	LeftSource    PlanPipelineComponent
	RightSource   PlanPipelineComponent
	OutputChannel DocumentChannel
	peekLeft      Document
	peekRight     Document
	LeftExpr      parser.Expression
	RightExpr     parser.Expression
	Otto          *otto.Otto
	joinExpr      parser.Expression
}

func NewOttoSortMergeJoiner() *OttoSortMergeJoiner {
	return &OttoSortMergeJoiner{
		OutputChannel: make(DocumentChannel),
		Otto:          otto.New()}
}

// PlanPipelineComponent interface

func (cpj *OttoSortMergeJoiner) SetSource(s PlanPipelineComponent) {
	log.Fatalf("SetSource called on DataSource, use ")
}

func (cpj *OttoSortMergeJoiner) GetSource() PlanPipelineComponent {
	return nil
}

func (cpj *OttoSortMergeJoiner) GetDocumentChannel() DocumentChannel {
	return cpj.OutputChannel
}

func (cpj *OttoSortMergeJoiner) Run() {

	defer close(cpj.OutputChannel)

	// tell our sources to start
	go cpj.LeftSource.Run()
	go cpj.RightSource.Run()

	// get our sources channel
	//	leftSourceChannel := cpj.LeftSource.GetDocumentChannel()
	//	rightSourceChannel := cpj.RightSource.GetDocumentChannel()

	ldocs, lkey := cpj.advanceLeft()
	rdocs, rkey := cpj.advanceRight()

	for len(ldocs) != 0 && len(rdocs) != 0 {
		cpj.Otto.Set("left", lkey)
		cpj.Otto.Set("right", rkey)
		keysEqual := evaluateExpressionStringInEnvironmentAsBoolean(cpj.Otto, "left == right")

		if keysEqual {
			for _, l := range ldocs {
				for _, r := range rdocs {
					combined := combineDocs(l, r)
					cpj.OutputChannel <- combined
				}
			}
			ldocs, lkey = cpj.advanceLeft()
			rdocs, rkey = cpj.advanceRight()
		} else {
			leftKeyLess := evaluateExpressionStringInEnvironmentAsBoolean(cpj.Otto, "left < right")

			if leftKeyLess {
				ldocs, lkey = cpj.advanceLeft()
			} else {
				rdocs, rkey = cpj.advanceRight()
			}

		}
		cpj.Otto.Set("left", otto.UndefinedValue())
		cpj.Otto.Set("right", otto.UndefinedValue())
	}
}

func (cpj *OttoSortMergeJoiner) peekFromLeft() Document {
	if cpj.peekLeft == nil {
		cpj.peekLeft = <-cpj.LeftSource.GetDocumentChannel()
	}
	return cpj.peekLeft
}

func (cpj *OttoSortMergeJoiner) consumeFromLeft() Document {
	result := cpj.peekLeft
	if result != nil {
		cpj.peekLeft = nil
	} else {
		result = <-cpj.LeftSource.GetDocumentChannel()
	}
	return result
}

func (cpj *OttoSortMergeJoiner) peekFromRight() Document {
	if cpj.peekRight == nil {
		cpj.peekRight = <-cpj.RightSource.GetDocumentChannel()
	}
	return cpj.peekRight
}

func (cpj *OttoSortMergeJoiner) consumeFromRight() Document {
	result := cpj.peekRight
	if result != nil {
		cpj.peekRight = nil
	} else {
		result = <-cpj.RightSource.GetDocumentChannel()
	}
	return result
}

func (cpj *OttoSortMergeJoiner) advanceLeft() ([]Document, otto.Value) {
	result := make([]Document, 0)

	var key otto.Value
	peekDoc := cpj.peekFromLeft()
	if peekDoc != nil {
		putDocumentIntoEnvironment(cpj.Otto, peekDoc)
		key = evaluateExpressionInEnvironment(cpj.Otto, cpj.LeftExpr)
		thisDocVal := key
		cleanupDocumentFromEnvironment(cpj.Otto, peekDoc)

		for peekDoc != nil && key == thisDocVal {
			result = append(result, cpj.consumeFromLeft())

			peekDoc = cpj.peekFromLeft()
			if peekDoc != nil {
				putDocumentIntoEnvironment(cpj.Otto, peekDoc)
				thisDocVal = evaluateExpressionInEnvironment(cpj.Otto, cpj.LeftExpr)
				cleanupDocumentFromEnvironment(cpj.Otto, peekDoc)
			}
		}
	}

	return result, key
}

func (cpj *OttoSortMergeJoiner) advanceRight() ([]Document, otto.Value) {
	result := make([]Document, 0)

	var key otto.Value
	peekDoc := cpj.peekFromRight()
	if peekDoc != nil {
		putDocumentIntoEnvironment(cpj.Otto, peekDoc)
		key = evaluateExpressionInEnvironment(cpj.Otto, cpj.RightExpr)
		thisDocVal := key
		cleanupDocumentFromEnvironment(cpj.Otto, peekDoc)

		for peekDoc != nil && key == thisDocVal {
			result = append(result, cpj.consumeFromRight())

			peekDoc = cpj.peekFromRight()
			if peekDoc != nil {
				putDocumentIntoEnvironment(cpj.Otto, peekDoc)
				thisDocVal = evaluateExpressionInEnvironment(cpj.Otto, cpj.RightExpr)
				cleanupDocumentFromEnvironment(cpj.Otto, peekDoc)
			}
		}
	}

	return result, key
}

func (cpj *OttoSortMergeJoiner) Explain() {

	defer close(cpj.OutputChannel)

	// tell our sources to explain
	go cpj.LeftSource.Explain()
	go cpj.RightSource.Explain()

	// get our sources channel
	leftSourceChannel := cpj.LeftSource.GetDocumentChannel()
	rightSourceChannel := cpj.RightSource.GetDocumentChannel()

	// read from the channel until its closed
	var l Document
	for doc := range leftSourceChannel {
		l = doc
	}

	// read from the channel until its closed
	var r Document
	for doc := range rightSourceChannel {
		r = doc
	}

	thisStep := map[string]interface{}{
		"_type":     "JOIN",
		"impl":      "Otto Sort Merge",
		"condition": fmt.Sprintf("%v", cpj.joinExpr),
		"left":      l,
		"right":     r}

	cpj.OutputChannel <- thisStep

}

func (cpj *OttoSortMergeJoiner) Cancel() {

}

// Joiner interface

func (cpj *OttoSortMergeJoiner) SetLeftSource(l PlanPipelineComponent) {
	cpj.LeftSource = l
}

func (cpj *OttoSortMergeJoiner) GetLeftSource() PlanPipelineComponent {
	return cpj.LeftSource
}

func (cpj *OttoSortMergeJoiner) SetRightSource(r PlanPipelineComponent) {
	cpj.RightSource = r
}

func (cpj *OttoSortMergeJoiner) GetRightSource() PlanPipelineComponent {
	return cpj.RightSource
}

func (cpj *OttoSortMergeJoiner) SetCondition(e parser.Expression) error {
	// make sure this is a condition we can support
	equalsExpr, isEqualsExpression := e.(*parser.EqualsExpression)
	if isEqualsExpression {
		// now check that left and right are simple properties
		leftExpr, isLeftProperty := equalsExpr.Left.(*parser.Property)
		rightExpr, isRightProperty := equalsExpr.Right.(*parser.Property)
		if isLeftProperty && isRightProperty {
			//ok we are good to go
			cpj.LeftExpr = leftExpr
			cpj.RightExpr = rightExpr
			cpj.joinExpr = e
			return nil
		}
	}

	return fmt.Errorf("Sort merge can only be performed on equals expressions of simple properties")
}

func (cpj *OttoSortMergeJoiner) GetCondition() parser.Expression {
	return cpj.joinExpr
}
