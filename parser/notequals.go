package parser

import (
//"log"
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