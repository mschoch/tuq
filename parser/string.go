package parser

import ()

type StringLiteral struct {
	Val string
}

func NewStringLiteral(v string) *StringLiteral {
	return &StringLiteral{
		Val: v}
}
