package parser

import (
	"fmt"
	"strings"
	"sync"
)

var parsingStack *Stack
var parsingQuery *Select

var debugTokens = false
var debugGrammar = false
var crashHard = false

type UnqlParser struct {
	mutex sync.Mutex
}

func NewUnqlParser(dt, db, ch bool) *UnqlParser {
	debugTokens = dt
	debugGrammar = db
	crashHard = ch
	return &UnqlParser{}
}

func (u *UnqlParser) Parse(input string) (returnQuery *Select, err error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	parsingStack = new(Stack)
	parsingQuery = NewSelect()

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
