package parser

import (
	//"log"
	"fmt"
)

type EqualsExpression struct {
	Left  Expression
	Right Expression
}

func NewEqualsExpression(l, r Expression) *EqualsExpression {
	return &EqualsExpression{
		Left:  l,
		Right: r}
}

func (ee *EqualsExpression) String() string {
	return fmt.Sprintf("%v == %v", ee.Left, ee.Right)
}

func (ee *EqualsExpression) SymbolsReferenced() []string {
	leftSymbols := ee.Left.SymbolsReferenced()
	return concatStringSlices(leftSymbols, ee.Right.SymbolsReferenced())
}

func (ee *EqualsExpression) PrefixSymbols(s string) {
    ee.Left.PrefixSymbols(s)
    ee.Right.PrefixSymbols(s)
}
