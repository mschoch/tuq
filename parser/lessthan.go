package parser

import (
	//"log"
	"fmt"
)

type LessThanExpression struct {
	Left  Expression
	Right Expression
}

func NewLessThanExpression(l, r Expression) *LessThanExpression {
	return &LessThanExpression{
		Left:  l,
		Right: r}
}

func (lt *LessThanExpression) String() string {
	return fmt.Sprintf("%v < %v", lt.Left, lt.Right)
}

func (lt *LessThanExpression) SybolsReferenced() []string {
	leftSymbols := lt.Left.SybolsReferenced()
	return concatStringSlices(leftSymbols, lt.Right.SybolsReferenced())
}

func (lt *LessThanExpression) PrefixSymbols(s string) {
	lt.Left.PrefixSymbols(s)
	lt.Right.PrefixSymbols(s)
}
