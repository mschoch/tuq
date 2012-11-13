package parser

import (
	"fmt"
)

type StringLiteral struct {
	Val string
}

func NewStringLiteral(v string) *StringLiteral {
	return &StringLiteral{
		Val: v}
}

func (sl *StringLiteral) String() string {
	return fmt.Sprintf("\"%s\"", sl.Val)
}

func (sl *StringLiteral) SybolsReferenced() []string {
	return []string{}
}

func (sl *StringLiteral) PrefixSymbols(string) {

}
