package parser

import (
	//"log"
	"fmt"
)

type GreaterThanExpression struct {
	Left  Expression
	Right Expression
}

func NewGreaterThanExpression(l, r Expression) *GreaterThanExpression {
	return &GreaterThanExpression{
		Left:  l,
		Right: r}
}

func (gt *GreaterThanExpression) String() string {
	return fmt.Sprintf("%v > %v", gt.Left, gt.Right)
}

func (gt *GreaterThanExpression) SybolsReferenced() []string {
	leftSymbols := gt.Left.SybolsReferenced()
	return concatStringSlices(leftSymbols, gt.Right.SybolsReferenced())
}

func (gt *GreaterThanExpression) PrefixSymbols(s string) {
	gt.Left.PrefixSymbols(s)
	gt.Right.PrefixSymbols(s)
}
