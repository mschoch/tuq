package main

import (
	"flag"
	"github.com/mschoch/elastigo/api"
	"log"
	"net/url"
	"strings"
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
var stdinMode = flag.Bool("stdin", false, "Read statements from STDIN, execute them, print results to STDOUT, then exit")
var httpMode = flag.Bool("http", false, "Answer queries received over HTTP")

func main() {

	flag.Parse()

	setupElasticSearch()

	if *stdinMode {
		handleStdinMode()
	} else if *httpMode {
		handleHttpMode()
	} else {
		handleInteractiveMode()
	}

}

func setupElasticSearch() {

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
}
