package main

import (
	"encoding/json"
	"fmt"
	naiveoptimizer "github.com/mschoch/tuq/optimizer/naive"
	nulloptimizer "github.com/mschoch/tuq/optimizer/null"
	"github.com/mschoch/tuq/parser/tuql"
	"github.com/mschoch/tuq/planner"
	naiveplanner "github.com/mschoch/tuq/planner/naive"
	"log"
	"net/http"
)

var fileServer http.Handler

func handleHttpMode() {

	fileServer = http.FileServer(http.Dir("html"))

	s := &http.Server{
		Addr:        *bindAddr,
		Handler:     http.HandlerFunc(httpHandler),
		ReadTimeout: *readTimeout,
	}
	log.Printf("Listening to web requests on %s",
		*bindAddr)
	log.Fatal(s.ListenAndServe())

}

func httpHandler(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	switch req.Method {
	case "POST":
		doPost(w, req)
	case "GET":
		doGet(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func doGet(w http.ResponseWriter, req *http.Request) {
	fileServer.ServeHTTP(w, req)
}

func doPost(w http.ResponseWriter, req *http.Request) {

	d := json.NewDecoder(req.Body)
	request := map[string]interface{}{}

	err := d.Decode(&request)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error parsing request JSON: %v", err)
		return
	}

	line, ok := request["query"].(string)
	if !ok {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error reading query string: %v", err)
		return
	}

	unqlParser := tuql.NewTuqlParser(false, false, *crashHard)
	naiveOptimizer := naiveoptimizer.NewNaiveOptimizer()
	nullOptimizer := nulloptimizer.NewNullOptimizer()
	query, err := unqlParser.Parse(line)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error parsing query string: %v", err)
		return
	} else {

		if query.WasParsedSuccessfully() {
			// check to make sure the query is semantically valid
			err := query.Validate()

			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Error validating query string: %v", err)
				return
			} else {
				naivePlanner := naiveplanner.NewNaivePlanner()
				plans := naivePlanner.Plan(*query)

				if plans != nil {
					var plan planner.Plan
					if *disableOptimizer {
						plan = nullOptimizer.Optimize(plans)
					} else {
						plan = naiveOptimizer.Optimize(plans)
					}

					if query.IsExplainOnly {
						result := plan.Explain()
						if err != nil {
							w.WriteHeader(500)
							fmt.Fprintf(w, "Error explaining query string: %v", err)
							return
						} else {
							FormatChannelOutput(result, w)
						}
					} else {
						result := plan.Run()
						if err != nil {
							w.WriteHeader(500)
							fmt.Fprintf(w, "Error running query string: %v", err)
							return
						} else {
							FormatChannelOutput(result, w)
						}
					}

				}
			}
		} else {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Error parsing query string: %v", err)
			return
		}
	}
}
