package parser

import (
	"fmt"
)

type OrExpression struct {
	Left  Expression
	Right Expression
}

func NewOrExpression(l, r Expression) *OrExpression {
	return &OrExpression{
		Left:  l,
		Right: r}
}

func (oe *OrExpression) String() string {
	return fmt.Sprintf("%v || %v", oe.Left, oe.Right)
}

func (oe *OrExpression) SymbolsReferenced() []string {
	leftSymbols := oe.Left.SymbolsReferenced()
	return concatStringSlices(leftSymbols, oe.Right.SymbolsReferenced())
}

func (oe *OrExpression) PrefixSymbols(s string) {
    oe.Left.PrefixSymbols(s)
    oe.Right.PrefixSymbols(s)
}
