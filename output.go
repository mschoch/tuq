package main

import (
	"encoding/json"
	"fmt"
	"github.com/mschoch/tuq/planner"
	"io"
	"log"
	"math"
)

func FormatOutput(result interface{}) {
	body, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Printf("Unable to format result to display %v", err)
	} else {
		fmt.Printf(string(body) + "\n")
	}
}

func FormatChannelOutput(result planner.RowChannel, w io.Writer) {
	fmt.Fprint(w, "[\n")
	first := true
	for row := range result {
		row = ReplaceNaNAndInfRecursive(row)
		if !first {
			fmt.Fprint(w, ",\n")
		}
		body, err := json.MarshalIndent(row, "    ", "    ")
		if err != nil {
			log.Printf("Unable to format result to display %#v, %v", row, err)
		} else {
			fmt.Fprintf(w, "    %v", string(body))
		}
		first = false
	}
	fmt.Fprint(w, "\n]\n")
}

// this should go away once this bug is resolved:
// http://code.google.com/p/go/issues/detail?id=3480
func ReplaceNaNAndInfRecursive(row planner.Row) planner.Row {
	switch row := row.(type) {
	case map[string]interface{}:
		for k, v := range row {
			switch v := v.(type) {
			case planner.Row:
				row[k] = ReplaceNaNAndInfRecursive(v)
			case float64:
				if math.IsNaN(v) {
					row[k] = nil
				} else if math.IsInf(v, 1) {
					row[k] = "Infinity"
				} else if math.IsInf(v, -1) {
					row[k] = "-Infinity"
				}
			}
		}
	case []interface{}:
		for i, v := range row {
			switch v := v.(type) {
			case planner.Row:
				row[i] = ReplaceNaNAndInfRecursive(v)
			case float64:
				if math.IsNaN(v) {
					row[i] = nil
				} else if math.IsInf(v, 1) {
					row[i] = "Infinity"
				} else if math.IsInf(v, -1) {
					row[i] = "-Infinity"
				}
			}
		}
	case planner.Document:
		for k, v := range row {
			switch v := v.(type) {
			case planner.Row:
				row[k] = ReplaceNaNAndInfRecursive(v)
			case float64:
				if math.IsNaN(v) {
					row[k] = nil
				} else if math.IsInf(v, 1) {
					row[k] = "Infinity"
				} else if math.IsInf(v, -1) {
					row[k] = "-Infinity"
				}
			}
		}

	case float64:
		if math.IsNaN(row) {
			return nil
		} else if math.IsInf(row, 1) {
			return "Infinity"
		} else if math.IsInf(row, -1) {
			return "-Infinity"
		}

	case string:
		return row

	default:
		log.Printf("unexpected type %T", row)
	}
	return row
}
