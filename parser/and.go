package parser

import (
//	"log"
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
