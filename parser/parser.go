package parser

var DebugTokens = false
var DebugGrammar = false

type Parser interface {
	Parse(input string) (returnQuery *Select, err error)
}
