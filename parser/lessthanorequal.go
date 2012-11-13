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

func (lte *LessThanOrEqualExpression) SybolsReferenced() []string {
	leftSymbols := lte.Left.SybolsReferenced()
	return concatStringSlices(leftSymbols, lte.Right.SybolsReferenced())
}

func (lte *LessThanOrEqualExpression) PrefixSymbols(s string) {
	lte.Left.PrefixSymbols(s)
	lte.Right.PrefixSymbols(s)
}
