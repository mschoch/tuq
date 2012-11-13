package main

import (
	"encoding/json"
	"fmt"
	"github.com/mschoch/go-unql-couchbase/planner"
	"io"
	"log"
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
		if !first {
			fmt.Fprint(w, ",\n")
		}
		body, err := json.MarshalIndent(row, "    ", "    ")
		if err != nil {
			log.Printf("Unable to format result to display %v", err)
		} else {
			fmt.Fprintf(w, "    %v", string(body))
		}
		first = false
	}
	fmt.Fprint(w, "\n]\n")
}
