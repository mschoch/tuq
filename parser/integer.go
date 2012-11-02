package parser

import ()

type IntegerLiteral struct {
	Val int
}

func NewIntegerLiteral(v int) *IntegerLiteral {
	return &IntegerLiteral{
		Val: v}
}

