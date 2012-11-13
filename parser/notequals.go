package parser

import (
	//"log"
	"fmt"
)

type NotEqualsExpression struct {
	Left  Expression
	Right Expression
}

func NewNotEqualsExpression(l, r Expression) *NotEqualsExpression {
	return &NotEqualsExpression{
		Left:  l,
		Right: r}
}

func (ne *NotEqualsExpression) String() string {
	return fmt.Sprintf("%v != %v", ne.Left, ne.Right)
}

func (ne *NotEqualsExpression) SybolsReferenced() []string {
	leftSymbols := ne.Left.SybolsReferenced()
	return concatStringSlices(leftSymbols, ne.Right.SybolsReferenced())
}

func (ne *NotEqualsExpression) PrefixSymbols(s string) {
    ne.Left.PrefixSymbols(s)
    ne.Right.PrefixSymbols(s)
}
