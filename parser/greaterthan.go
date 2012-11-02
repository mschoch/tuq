package parser

import (
//"log"
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
