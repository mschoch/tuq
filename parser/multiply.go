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

func (me *MultiplyExpression) SybolsReferenced() []string {
	leftSymbols := me.Left.SybolsReferenced()
	return concatStringSlices(leftSymbols, me.Right.SybolsReferenced())
}

func (me *MultiplyExpression) PrefixSymbols(s string) {
    me.Left.PrefixSymbols(s)
    me.Right.PrefixSymbols(s)
}
