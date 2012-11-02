package parser

import ()

type FloatLiteral struct {
	Val float64
}

func NewFloatLiteral(v float64) *FloatLiteral {
	return &FloatLiteral{
		Val: v}
}