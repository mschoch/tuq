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
		if r := recover(); r != nil {
			err = fmt.Errorf("Parse Error - %v", r)
		}
	}()

	yyParse(NewLexer(strings.NewReader(input)))
	returnQuery = parsingQuery
	return
}
