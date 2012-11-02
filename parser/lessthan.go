package parser

import (
//"log"
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
