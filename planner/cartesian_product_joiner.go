package planner

import (
	"fmt"
	"github.com/mschoch/tuq/parser"
	"log"
)

type CartesianProductJoiner struct {
	LeftSource    PlanPipelineComponent
	RightSource   PlanPipelineComponent
	OutputChannel DocumentChannel
	leftDocs      []Document
	rightDocs     []Document
}

func NewCartesianProductJoiner() *CartesianProductJoiner {
	return &CartesianProductJoiner{
		OutputChannel: make(DocumentChannel),
		leftDocs:      make([]Document, 0),
		rightDocs:     make([]Document, 0)}
}

// PlanPipelineComponent interface

func (cpj *CartesianProductJoiner) SetSource(s PlanPipelineComponent) {
	log.Fatalf("SetSource called on DataSource, use ")
}

func (cpj *CartesianProductJoiner) GetSource() PlanPipelineComponent {
	return nil
}

func (cpj *CartesianProductJoiner) GetDocumentChannel() DocumentChannel {
	return cpj.OutputChannel
}

func (cpj *CartesianProductJoiner) Run() {

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
			cpj.OutputChannel <- combined
		}
	}

}

func (cpj *CartesianProductJoiner) Explain() {

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
				"impl":      "Full Cartesian Product",
				"left":      l,
				"right":     r,
				"cost":      cpj.Cost(),
				"totalCost": cpj.TotalCost()}

			cpj.OutputChannel <- thisStep

		}
	}

}

func (cpj *CartesianProductJoiner) Cancel() {

}

// Joiner interface

func (cpj *CartesianProductJoiner) SetLeftSource(l PlanPipelineComponent) {
	cpj.LeftSource = l
}

func (cpj *CartesianProductJoiner) SetRightSource(r PlanPipelineComponent) {
	cpj.RightSource = r
}

func (cpj *CartesianProductJoiner) SetCondition(parser.Expression) error {
	return fmt.Errorf("Cartesan Product Jointer does not support conditions")
}

// internal

// in theory, when joining, all datasources have already "named" their documents "as"
// something unique.  so we'd expect to only have 1 top-level element in each
// and that would be the name
func combineDocs(l, r Document) Document {
	combined := make(Document)
	for k, v := range l {
		combined[k] = v
	}
	for k, v := range r {
		combined[k] = v
	}
	return combined
}

func (cpj *CartesianProductJoiner) Cost() float64 {
	return cpj.LeftSource.TotalCost() * cpj.RightSource.TotalCost()
}

func (cpj *CartesianProductJoiner) TotalCost() float64 {
	return cpj.Cost() + cpj.LeftSource.TotalCost() + cpj.RightSource.TotalCost()
}
