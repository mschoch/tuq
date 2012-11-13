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

type UnqlParser struct {
	mutex sync.Mutex
}

func NewUnqlParser(dt, db bool) *UnqlParser {
	debugTokens = dt
	debugGrammar = db
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
		} else if r != nil{
			// otherise continue to panic
			panic(r)
		}
	}()

	yyParse(NewLexer(strings.NewReader(input)))
	returnQuery = parsingQuery
	return
}
