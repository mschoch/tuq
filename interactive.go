package main

import (
	naiveoptimizer "github.com/mschoch/tuq/optimizer/naive"
	nulloptimizer "github.com/mschoch/tuq/optimizer/null"
	"github.com/mschoch/tuq/parser/tuql"
	"github.com/mschoch/tuq/planner"
	naiveplanner "github.com/mschoch/tuq/planner/naive"
	"github.com/sbinet/liner"
	"log"
	"os"
	"os/signal"
	"os/user"
	"syscall"
)

func handleInteractiveMode() {

	unqlParser := tuql.NewTuqlParser(*debugTokens, *debugGrammar, *crashHard)
	naivePlanner := naiveplanner.NewNaivePlanner()
	naiveOptimizer := naiveoptimizer.NewNaiveOptimizer()
	nullOptimizer := nulloptimizer.NewNullOptimizer()

	currentUser, err := user.Current()
	if err != nil {
		log.Printf("Unable to determine home directory, history file disabled")
	}

	var liner = liner.NewLiner()
	defer liner.Close()

	LoadHistory(liner, currentUser)

	go signalCatcher(liner)

	for {
		line, err := liner.Prompt("tuq> ")
		if err != nil {
			break
		}

		if line == "" {
			continue
		}

		UpdateHistory(liner, currentUser, line)

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

/**
 *  Attempt to clean up after ctrl-C otherwise
 *  terminal is left in bad shape
 */
func signalCatcher(liner *liner.State) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
	liner.Close()
	os.Exit(0)
}
