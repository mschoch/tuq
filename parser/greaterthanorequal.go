package parser

import (
//"log"
)

type GreaterThanOrEqualExpression struct {
	Left  Expression
	Right Expression
}

func NewGreaterThanOrEqualExpression(l, r Expression) *GreaterThanOrEqualExpression {
	return &GreaterThanOrEqualExpression{
		Left:  l,
		Right: r}
}
