package parser

import (
//"log"
)

type EqualsExpression struct {
	Left  Expression
	Right Expression
}

func NewEqualsExpression(l, r Expression) *EqualsExpression {
	return &EqualsExpression{
		Left:  l,
		Right: r}
}