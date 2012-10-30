package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func FormatOutput(result interface{}) {
	body, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Printf("Unable to format result to display")
	} else {
		fmt.Printf(string(body) + "\n")
	}
}
