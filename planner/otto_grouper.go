package planner

import (
	"encoding/json"
	"fmt"
	"github.com/mschoch/tuq/parser"
	"github.com/robertkrimen/otto"
	"log"
	"math"
	"reflect"
)

type OttoGrouper struct {
	Source        PlanPipelineComponent
	OutputChannel DocumentChannel
	Otto          *otto.Otto
	GroupBy       parser.ExpressionList
	Stats         []string
	groupValues   map[string]ExpressionStatsMap
	groupDocs     map[string]Document
}

func NewOttoGrouper() *OttoGrouper {
	return &OttoGrouper{
		OutputChannel: make(DocumentChannel),
		Otto:          otto.New(),
		groupValues:   make(map[string]ExpressionStatsMap),
		groupDocs:     make(map[string]Document)}
}

// PlanPipelineComponent interface

func (og *OttoGrouper) SetSource(s PlanPipelineComponent) {
	og.Source = s
}

func (og *OttoGrouper) GetSource() PlanPipelineComponent {
    return og.Source
}

func (og *OttoGrouper) GetDocumentChannel() DocumentChannel {
	return og.OutputChannel
}

func (og *OttoGrouper) Run() {

	// tell our source to start
	go og.Source.Run()

	// get our sources channel
	sourceChannel := og.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {
		putDocumentIntoEnvironment(og.Otto, doc)

		groupDocument := Document{}
		groupKey := ""
		for _, expr := range og.GroupBy {
			value := evaluateExpressionInEnvironment(og.Otto, expr)

			val := convertToPrimitive(value)
			//groupDocument[expr.String()] = val
			SetDocumentProperty(groupDocument, expr, val)

			key, err := value.ToString() // this probably doesnt do what we want for Objects or Arrays
			if err != nil {
				log.Printf("Error converting to string %v", err)
			} // FIXME handle error

			groupKey += key
		}

		// now walk through the group by expressions again
		// this time to compute any necessary stats
		// we need a complete list of fields to compute stats on
		// this should be og.Stats and og.Group, preferrably without duplicates
		// we dont expect huge volume here, lets just build a map
		// and then walk the keys
		statsFieldMap := make(map[string]interface{})
		for _, expr := range og.GroupBy {
		  statsFieldMap[expr.String()] = nil
		}
		for _, field := range og.Stats {
		  statsFieldMap[field] = nil
		}

        for expr, _ := range statsFieldMap {
			// find the entry in the group by map (if it exists
			statsMap, ok := og.groupValues[groupKey]

			value := evaluateExpressionInEnvironment(og.Otto, parser.NewProperty(expr)) 
			if ok {
				expr_stats, ok := statsMap[expr]
				if ok {
					expr_stats.ConsiderValue(value)
					statsMap[expr] = expr_stats
				} else {
					stats := NewExpressionStats()
					stats.ConsiderValue(value)
					statsMap[expr] = *stats
				}
			} else {
				// first time we've seen this value
				statsMap = make(ExpressionStatsMap)
				stats := NewExpressionStats()
				stats.ConsiderValue(value)
				statsMap[expr] = *stats
				og.groupValues[groupKey] = statsMap
			}

			for stat_field, stats := range statsMap {
				SetDocumentProperty(groupDocument, parser.NewProperty(fmt.Sprintf("__func__.sum.%v", stat_field)), stats.Sum)
				SetDocumentProperty(groupDocument, parser.NewProperty(fmt.Sprintf("__func__.avg.%v", stat_field)), stats.Avg)
				SetDocumentProperty(groupDocument, parser.NewProperty(fmt.Sprintf("__func__.count.%v", stat_field)), stats.Count)

				if math.IsInf(stats.Min, 1) {
					SetDocumentProperty(groupDocument, parser.NewProperty(fmt.Sprintf("__func__.min.%v", stat_field)), "Infinity")
				} else if math.IsInf(stats.Min, -1) {
					SetDocumentProperty(groupDocument, parser.NewProperty(fmt.Sprintf("__func__.min.%v", stat_field)), "-Infinity")
				} else {
					SetDocumentProperty(groupDocument, parser.NewProperty(fmt.Sprintf("__func__.min.%v", stat_field)), stats.Min)
				}
				if math.IsInf(stats.Max, 1) {
					SetDocumentProperty(groupDocument, parser.NewProperty(fmt.Sprintf("__func__.max.%v", stat_field)), "Infinity")
				} else if math.IsInf(stats.Max, -1) {
					SetDocumentProperty(groupDocument, parser.NewProperty(fmt.Sprintf("__func__.max.%v", stat_field)), "-Infinity")
				} else {
					SetDocumentProperty(groupDocument, parser.NewProperty(fmt.Sprintf("__func__.max.%v", stat_field)), stats.Max)
				}
			}
		}

		og.groupDocs[groupKey] = groupDocument
		cleanupDocumentFromEnvironment(og.Otto, doc)
	}

	for _, v := range og.groupDocs {
		og.OutputChannel <- v
	}

	close(og.OutputChannel)

}

func (og *OttoGrouper) Explain() {

	defer close(og.OutputChannel)

	// tell our source to explain
	go og.Source.Explain()

	// get our sources channel
	sourceChannel := og.Source.GetDocumentChannel()

	// read from the channel until its closed
	for doc := range sourceChannel {

		thisStep := map[string]interface{}{
			"_type":      "GROUP BY",
			"impl":       "Otto",
			"expression": fmt.Sprintf("%v", og.GroupBy),
			"source":     doc}

		og.OutputChannel <- thisStep
	}

}

func (og *OttoGrouper) Cancel() {

}

// Grouper interface
func (og *OttoGrouper) SetGroupByWithStatsFields(g parser.ExpressionList, s []string) {
	og.GroupBy = g
	og.Stats = s
}

func (og *OttoGrouper) GetGroupByWithStatsFields() (parser.ExpressionList, []string) {
	return og.GroupBy, og.Stats
}

type ExpressionStatsMap map[string]ExpressionStats

type ExpressionStats struct {
	Count int               `json:"count"`
	Sum   float64           `json:"sum"`
	Min   float64           `json:"min"`
	Max   float64           `json:"max"`
	Avg   float64           `json:"avg"`
}

func (es ExpressionStats) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{}

	val := reflect.ValueOf(es)
	for i := 0; i < val.NumField(); i++ {
		v := (interface{})(val.Field(i).Interface())
		x, ok := v.(float64)
		if ok && math.IsInf(x, 1) {
			v = "Infinity"
		} else if ok && math.IsInf(x, -1) {
			v = "-Infinity"
		}

		m[jsonFieldName(val.Type().Field(i))] = v
	}

	return json.Marshal(m)
}

func jsonFieldName(sf reflect.StructField) string {
	fieldName := sf.Tag.Get("json")
	if fieldName == "" {
		fieldName = sf.Name
	}
	return fieldName
}

func NewExpressionStats() *ExpressionStats {
	return &ExpressionStats{
		Count: 0,
		Sum:   0,
		Min:   math.Inf(+1),
		Max:   math.Inf(-1),
		Avg:   0}
}

func (es *ExpressionStats) ConsiderValue(val otto.Value) {
	// increment the count
	es.Count += 1

	if val.IsNumber() {
		f, err := val.ToFloat()
		if err != nil {
			log.Printf("Error converting number to float %v", err)
		} else {
			// update the sum
			es.Sum += f

			// if this is smaller than anything we've seen so far update the min
			if f < es.Min {
				es.Min = f
			}

			// if this is larger than anything we've seen so far update the max
			if f > es.Max {
				es.Max = f
			}

			// update the average (perhaps wasteful, could be done once at the end
			// but i'd have to walk the whole tree again, for now will do update it each time
			es.Avg = es.Sum / float64(es.Count) // we incremented es.count in this function, can not divide by 0
		}
	}
}

func SetDocumentProperty(doc Document, property parser.Expression, value interface{}) {
	switch prop := property.(type) {
	case *parser.Property:
		if prop.HasSubProperty() {
			next, exists := doc[prop.Head()]
			if !exists {
				next = make(Document)
				doc[prop.Head()] = next
			}
			SetDocumentProperty(next.(Document), prop.Tail(), value)
		} else {
			doc[prop.Head()] = value
		}
	}
}
