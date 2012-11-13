package parser

import (
	//	"log"
	"fmt"
)

type AndExpression struct {
	Left  Expression
	Right Expression
}

func NewAndExpression(l, r Expression) *AndExpression {
	return &AndExpression{
		Left:  l,
		Right: r}
}

func (ae *AndExpression) String() string {
	return fmt.Sprintf("%v && %v", ae.Left, ae.Right)
}

func (ae *AndExpression) SybolsReferenced() []string {
	leftSymbols := ae.Left.SybolsReferenced()
	return concatStringSlices(leftSymbols, ae.Right.SybolsReferenced())
}

func (ae *AndExpression) PrefixSymbols(s string) {
    ae.Left.PrefixSymbols(s)
    ae.Right.PrefixSymbols(s)
}
