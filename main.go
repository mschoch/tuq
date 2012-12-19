package main

import (
	"flag"
	"github.com/mschoch/tuq/datasources"
	"time"

	// register the datasources you want to support
	_ "github.com/mschoch/tuq/datasources/couchbase"
	_ "github.com/mschoch/tuq/datasources/csv"
	_ "github.com/mschoch/tuq/datasources/datasources"
	_ "github.com/mschoch/tuq/datasources/elasticsearch"
	_ "github.com/mschoch/tuq/datasources/mongodb"
	_ "github.com/mschoch/tuq/datasources/jsondir"
)

var debugTokens = flag.Bool("debugTokens", false, "Enable debug of all tokens seen by the lexer")
var debugGrammar = flag.Bool("debugGrammar", false, "Enable debug of all grammar rules processed by the parser")
var crashHard = flag.Bool("crashHard", false, "Enable hard crashing during parsing (useful when developing)")
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
