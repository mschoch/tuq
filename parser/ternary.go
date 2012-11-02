package parser

import (
//	"log"
)

type TernaryExpression struct {
	Iff   Expression
	Thenn Expression
	Elsee Expression
}

func NewTernaryExpression(i, t, e Expression) *TernaryExpression {
	return &TernaryExpression{
		Iff:   i,
		Thenn: t,
		Elsee: e}
}
