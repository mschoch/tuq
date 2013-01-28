package main

import (
	"bufio"
	naiveoptimizer "github.com/mschoch/tuq/optimizer/naive"
	nulloptimizer "github.com/mschoch/tuq/optimizer/null"
	"github.com/mschoch/tuq/parser/tuql"
	"github.com/mschoch/tuq/planner"
	naiveplanner "github.com/mschoch/tuq/planner/naive"
	"log"
	"os"
)

func handleStdinMode() {

	unqlParser := tuql.NewTuqlParser(*debugTokens, *debugGrammar, *crashHard)
	naivePlanner := naiveplanner.NewNaivePlanner()
	naiveOptimizer := naiveoptimizer.NewNaiveOptimizer()
	nullOptimizer := nulloptimizer.NewNullOptimizer()

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		if line == "" {
			continue
		}

		query, err := unqlParser.Parse(line)
		if err != nil {
			log.Printf("%v", err)
		} else {
			if *debugGrammar {
				log.Printf("Query is: %#v", query)
			}
			if query.WasParsedSuccessfully() {
				// check to make sure the query is semantically valid
				err := query.Validate()
				if err != nil {
					log.Printf("%v", err)
				} else {
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
								log.Printf("Error: %v", err)
							} else {
								FormatChannelOutput(result, os.Stdout)
							}
						} else {
							result := plan.Run()
							if err != nil {
								log.Printf("Error: %v", err)
							} else {
								FormatChannelOutput(result, os.Stdout)
							}
						}

					}
				}

			}
		}
	}

}
