package parser

import (
	//	"log"
	"fmt"
)

type MultiplyExpression struct {
	Left  Expression
	Right Expression
}

func NewMultiplyExpression(l, r Expression) *MultiplyExpression {
	return &MultiplyExpression{
		Left:  l,
		Right: r}
}

func (me *MultiplyExpression) String() string {
	return fmt.Sprintf("%v * %v", me.Left, me.Right)
}

func (me *MultiplyExpression) SymbolsReferenced() []string {
	leftSymbols := me.Left.SymbolsReferenced()
	return concatStringSlices(leftSymbols, me.Right.SymbolsReferenced())
}

func (me *MultiplyExpression) PrefixSymbols(s string) {
    me.Left.PrefixSymbols(s)
    me.Right.PrefixSymbols(s)
}
