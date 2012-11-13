package planner

import (
	"encoding/json"
	"fmt"
	"github.com/mschoch/go-unql-couchbase/parser"
	"github.com/robertkrimen/otto"
	"log"
	"sort"
)

type OttoOrderer struct {
	Source        PlanPipelineComponent
	OutputChannel DocumentChannel
	SortList      parser.SortList
	Otto          *otto.Otto
	docs          []Document
}

func NewOttoOrderer() *OttoOrderer {
	return &OttoOrderer{
		OutputChannel: make(DocumentChannel),
		Otto:          otto.New(),
		docs:          make([]Document, 0)}
}

// PlanPipelineComponent interface

func (oo *OttoOrderer) SetSource(s PlanPipelineComponent) {
	oo.Source = s
}

func (oo *OttoOrderer) GetSource() PlanPipelineComponent {
	return oo.Source
}

func (oo *OttoOrderer) GetDocumentChannel() DocumentChannel {
	return oo.OutputChannel
}

func (oo *OttoOrderer) Run() {
	defer close(oo.OutputChannel)

	// tell our source to start
	go oo.Source.Run()

	// get our sources channel
	sourceChannel := oo.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {
		oo.docs = append(oo.docs, doc)
	}

	// sort them
	sort.Sort(oo)

	// write the sorted docs out
	for _, v := range oo.docs {
		oo.OutputChannel <- v
	}

}

func (oo *OttoOrderer) Explain() {
	defer close(oo.OutputChannel)

	// tell our source to explain
	go oo.Source.Explain()

	// get our sources channel
	sourceChannel := oo.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {
		thisStep := map[string]interface{}{
			"_type":      "ORDER BY",
			"impl":       "Otto",
			"expression": fmt.Sprintf("%v", oo.SortList),
			"source":     doc}

		oo.OutputChannel <- thisStep
	}
}

func (oo *OttoOrderer) Cancel() {

}

// Orderer interface
func (oo *OttoOrderer) SetOrderBy(sl parser.SortList) {
	oo.SortList = sl
}

func (oo *OttoOrderer) GetOrderBy() parser.SortList {
	return oo.SortList
}

// sort.Interface interface

func (oo *OttoOrderer) Len() int      { return len(oo.docs) }
func (oo *OttoOrderer) Swap(i, j int) { oo.docs[i], oo.docs[j] = oo.docs[j], oo.docs[i] }
func (oo *OttoOrderer) Less(i, j int) bool {
	elementA := oo.docs[i]
	elementB := oo.docs[j]

	// serialize a
	a_json, err := json.Marshal(elementA)
	if err != nil {
		log.Printf("JSON serialization failed: %v", err)
	}

	// serialize b
	b_json, err := json.Marshal(elementB)
	if err != nil {
		log.Printf("JSON serialization failed: %v", err)
	}

	// put a in the js environment
	_, err = oo.Otto.Run("a=" + string(a_json))
	if err != nil {
		log.Printf("Error running otto: %v", err)
	}

	// put b in the js environment
	_, err = oo.Otto.Run("b=" + string(b_json))
	if err != nil {
		log.Printf("Error running otto: %v", err)
	}

	for _, order := range oo.SortList {
		// FIXME, really need to add "a." and "b." to any symbols in the
		// expression, this wont work if expression is somehting like
		// len(doc.name) -> a.len(doc.name), should be len(a.doc.name)
		compareExpression := fmt.Sprintf("a.%v < b.%v", order.Sort, order.Sort)
		if order.Ascending == false {
			compareExpression = fmt.Sprintf("a.%v > b.%v", order.Sort, order.Sort)
		}
		result, err := oo.Otto.Run(compareExpression)
		if err != nil {
			log.Printf("Error running otto %v", err)
		} else {
			//log.Printf("result was %v", result)
			result, err := result.ToBoolean()
			if err != nil {
				log.Printf("Error converting otto result to boolean %v", err)
			} else if result {
				return true
			} else {
				// all we know now is that a does not sort before b
				// however, if they are the same, and there are more sort items
				// we may need to proceed deeper, check now to see if tehy were equal
				equalExpression := fmt.Sprintf("a.%v == b.%v", order.Sort, order.Sort)
				result, err := oo.Otto.Run(equalExpression)
				if err != nil {
					log.Printf("Error running otto %v", err)
				} else {
					result, err := result.ToBoolean()
					if err != nil {
						log.Printf("Error converting otto result to boolean %v", err)
					} else if result == false {
						// not equal either, we should break out of the loop and return false
						break
					}
				}
				//otherwise they are equal at this level, and we should allow the loop to continue
				//checking the next level of order by
			}
		}
	}

	return false
}
