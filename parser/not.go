package parser

import (
	//	"log"
	"fmt"
)

type NotExpression struct {
	Oper Expression
}

func NewNotExpression(o Expression) *NotExpression {
	return &NotExpression{
		Oper: o}
}

func (n *NotExpression) String() string {
	return fmt.Sprintf("!(%v)", n.Oper)
}

func (n *NotExpression) SymbolsReferenced() []string {
	return n.Oper.SymbolsReferenced()
}

func (n *NotExpression) PrefixSymbols(s string) {
    n.Oper.PrefixSymbols(s)
}
