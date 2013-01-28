package tuql

import (
	"fmt"
	"strings"
	"sync"
	"github.com/mschoch/tuq/parser"
)

var parsingStack *parser.Stack
var parsingQuery *parser.Select
var crashHard = false

type TuqlParser struct {
	mutex sync.Mutex
}

func NewTuqlParser(dt, db, ch bool) *TuqlParser {
	parser.DebugTokens = dt
	parser.DebugGrammar = db
	crashHard = ch
	return &TuqlParser{}
}

func (u *TuqlParser) Parse(input string) (returnQuery *parser.Select, err error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	parsingStack = new(parser.Stack)
	parsingQuery = parser.NewSelect()

	defer func() {
		r := recover()
		if r != nil && r == "syntax error" {
			// if we're panicing over a syntax error, chill
			err = fmt.Errorf("Parse Error - %v", r)
		} else if r != nil {
			// otherise continue to panic
			if crashHard {
				panic(r)
			} else {
				err = fmt.Errorf("Other Error - %v", r)
			}
		}
	}()

	yyParse(NewLexer(strings.NewReader(input)))
	returnQuery = parsingQuery
	return
}
