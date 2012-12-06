package parser

import (
	//"log"
	"fmt"
)

type LessThanOrEqualExpression struct {
	Left  Expression
	Right Expression
}

func NewLessThanOrEqualExpression(l, r Expression) *LessThanOrEqualExpression {
	return &LessThanOrEqualExpression{
		Left:  l,
		Right: r}
}

func (lte *LessThanOrEqualExpression) String() string {
	return fmt.Sprintf("%v <= %v", lte.Left, lte.Right)
}

func (lte *LessThanOrEqualExpression) SymbolsReferenced() []string {
	leftSymbols := lte.Left.SymbolsReferenced()
	return concatStringSlices(leftSymbols, lte.Right.SymbolsReferenced())
}

func (lte *LessThanOrEqualExpression) PrefixSymbols(s string) {
	lte.Left.PrefixSymbols(s)
	lte.Right.PrefixSymbols(s)
}
