package main

import (
	"flag"
	"fmt"
	"github.com/mschoch/elastigo/api"
	"github.com/sbinet/liner"
	"log"
	"net/url"
	"os"
	"os/signal"
	"os/user"
	"strings"
	"sync"
	"syscall"
	"time"
)

var debugTokens = flag.Bool("debugTokens", false, "Enable debug of all tokens seen by the lexer")
var debugGrammar = flag.Bool("debugGrammar", false, "Enable debug of all grammar rules processed by the parser")
var debugElasticSearch = flag.Bool("debugElasticSearch", false, "Enable debug of all ElasticSearch requests and responses")
var debugCouchbase = flag.Bool("debugCouchbase", false, "Enable debug of Couchbase operations")
var esBatchSize = flag.Int("esBatchSize", 10000, "Number of documents to retrieve from ES in a single query")
var esMaxAggregate = flag.Int("esMaxAggregate", 1000000, "Maximum number of aggregate results to return from ES in a single query")
var esURLString = flag.String("elasticsearch", "http://localhost:9200/", "The URL of your ElasticSearch server")
var esDefaultExcludeType = flag.String("esDefaultExcludeType", "couchbaseCheckpoint", "A document type to exclude by default typically to exclude Couchbase replication checkpoint documents")
var couchbaseServer = flag.String("couchbase", "http://localhost:8091/", "Couchbase URL")
var viewTimeout = flag.Duration("viewTimeout", 5*time.Second, "Couchbase view client read timeout")
var couchbaseBatchSize = flag.Int("couchbaseBatchSize", 10000, "Number of documents to retrieve from Couchbase in a single operation")

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

/**
 *  Wrap processing so that we can recover from panic
 */

// these variables represent global state
// they should only be modified by the owner of the lock (this function) 
var parsingMutex sync.Mutex
var parsingStack *Stack
var parsingQuery *Select

// the idea here is:
// 1 acquire the lock
// 2 allocate new stack and query objects
// 3 parse - mutating items from #2
// 4 return pointer to the query object
// 5 relase the lock (deferred)

func processNextLine(line string) (returnQuery *Select, err error) {

	parsingMutex.Lock()
	defer parsingMutex.Unlock()

	parsingStack = new(Stack)
	parsingQuery = NewSelect()

	defer func() {
		if r := recover(); r != nil {
            err = fmt.Errorf("Parse Error - %v", r)
		}
	}()

	yyParse(NewLexer(strings.NewReader(line)))
	returnQuery = parsingQuery
	return
}

func main() {

	flag.Parse()

	esURL, err := url.Parse(*esURLString)
	if err != nil {
		log.Fatalf("Unable to parse esURL: %s error: %v", esURL, err)
	} else {
		api.Protocol = esURL.Scheme
		colonIndex := strings.Index(esURL.Host, ":")
		if colonIndex < 0 {
			api.Domain = esURL.Host
			api.Port = "9200"
		} else {
			api.Domain = esURL.Host[0:colonIndex]
			api.Port = esURL.Host[colonIndex+1:]
		}

	}

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

		query, err := processNextLine(line)
		if err != nil {
			log.Printf("%v", err)
		} else {
			if *debugGrammar {
				log.Printf("Query is: %#v", query)
			}
			if query.parsedSuccessfully {
				result, err := query.Execute()
				if err != nil {
					log.Printf("Error: %v", err)
				} else {
				    FormatOutput(result)
				}
			}
		}
	}

}
