package main

import (
	"flag"
	"github.com/mschoch/go-unql-couchbase/datasources"
	_ "github.com/mschoch/go-unql-couchbase/datasources/csv"
	_ "github.com/mschoch/go-unql-couchbase/datasources/elasticsearch"
	_ "github.com/mschoch/go-unql-couchbase/datasources/couchbase"
	"time"
)

var debugTokens = flag.Bool("debugTokens", false, "Enable debug of all tokens seen by the lexer")
var debugGrammar = flag.Bool("debugGrammar", false, "Enable debug of all grammar rules processed by the parser")
var debugElasticSearch = flag.Bool("debugElasticSearch", false, "Enable debug of all ElasticSearch requests and responses")
var debugCouchbase = flag.Bool("debugCouchbase", false, "Enable debug of Couchbase operations")
var esMaxAggregate = flag.Int("esMaxAggregate", 1000000, "Maximum number of aggregate results to return from ES in a single query")
var couchbaseServer = flag.String("couchbase", "http://localhost:8091/", "Couchbase URL")
var viewTimeout = flag.Duration("viewTimeout", 5*time.Second, "Couchbase view client read timeout")
var couchbaseBatchSize = flag.Int("couchbaseBatchSize", 10000, "Number of documents to retrieve from Couchbase in a single operation")
var stdinMode = flag.Bool("stdin", false, "Read statements from STDIN, execute them, print results to STDOUT, then exit")
var httpMode = flag.Bool("http", false, "Answer queries received over HTTP")
var disableOptimizer = flag.Bool("disableOptimizer", false, "Disable query optimization")
var bindAddr = flag.String("bind", ":6009", "Address to bind web thing to")
var readTimeout = flag.Duration("serverTimeout", 5*time.Minute,
	"Web server read timeout")

func main() {

	flag.Parse()

	datasources.LoadDataSources()

	if *stdinMode {
		handleStdinMode()
	} else if *httpMode {
		handleHttpMode()
	} else {
		handleInteractiveMode()
	}

}
