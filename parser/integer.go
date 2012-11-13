package parser

import (
	"fmt"
)

type IntegerLiteral struct {
	Val int
}

func NewIntegerLiteral(v int) *IntegerLiteral {
	return &IntegerLiteral{
		Val: v}
}

func (il *IntegerLiteral) String() string {
	return fmt.Sprintf("%d", il.Val)
}

func (il *IntegerLiteral) SybolsReferenced() []string {
	return []string{}
}

func (il *IntegerLiteral) PrefixSymbols(string) {

}
