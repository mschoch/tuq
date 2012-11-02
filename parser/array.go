package parser

import ()

type ArrayLiteral struct {
	Val ExpressionList
}

func NewArrayLiteral(v ExpressionList) *ArrayLiteral {
	return &ArrayLiteral{
		Val: v}
}