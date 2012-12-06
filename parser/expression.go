package parser

import (
	"fmt"
)

type Expression interface {
	String() string
	SymbolsReferenced() []string
	PrefixSymbols(string)
}

// NOTE: this should be OK even if
// Expression eventually gets methods
// as an Expression List is not an Expression
// by itself (though it can become one inside
// a literal array)
type ExpressionList []Expression

func (el ExpressionList) String() string {
	result := ""
	for _, v := range el {
		result += fmt.Sprintf("%v", v)
	}
	return result
}
