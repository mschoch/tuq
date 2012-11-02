package main

import (
	"github.com/mschoch/go-unql-couchbasev2/parser"
	"github.com/sbinet/liner"
	"log"
	"os"
	"os/signal"
	"os/user"
	"syscall"
)

func handleInteractiveMode() {

	unqlParser := parser.NewUnqlParser(*debugTokens, *debugGrammar)
	queryExecutor := NewQueryExecutor()

	currentUser, err := user.Current()
	if err != nil {
		log.Printf("Unable to determine home directory, history file disabled")
	}

	var liner = liner.NewLiner()
	defer liner.Close()

	LoadHistory(liner, currentUser)

	go signalCatcher(liner)

	for {
		line, err := liner.Prompt("unql-couchbase> ")
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
				result, err := queryExecutor.Execute(*query)
				if err != nil {
					log.Printf("Error: %v", err)
				} else {
					FormatOutput(result)
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
