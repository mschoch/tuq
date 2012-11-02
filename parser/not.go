package parser

import (
//	"log"
)

type NotExpression struct {
	Oper Expression
}

func NewNotExpression(o Expression) *NotExpression {
	return &NotExpression{
		Oper: o}
}
