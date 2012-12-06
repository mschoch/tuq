package parser

import (
	"fmt"
)

type BracketMemberExpression struct {
	Left  *Property
	Right Expression
}

func NewBracketMemberExpression(l *Property, r Expression) *BracketMemberExpression {
	return &BracketMemberExpression{
		Left:  l,
		Right: r}
}

func (ae *BracketMemberExpression) String() string {
	return fmt.Sprintf("%v[%v]", ae.Left, ae.Right)
}

func (ae *BracketMemberExpression) SymbolsReferenced() []string {
	leftSymbols := ae.Left.SymbolsReferenced()
	return concatStringSlices(leftSymbols, ae.Right.SymbolsReferenced())
}

func (ae *BracketMemberExpression) PrefixSymbols(s string) {
	ae.Left.PrefixSymbols(s)
	ae.Right.PrefixSymbols(s)
}
