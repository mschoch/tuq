package parser

import (
//"log"
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
