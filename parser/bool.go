package parser

import ()

type BoolLiteral struct {
	Val bool
}

func NewBoolLiteral(v bool) *BoolLiteral {
	return &BoolLiteral{
		Val: v}
}
