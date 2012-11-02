package parser

import ()

type OrExpression struct {
	Left  Expression
	Right Expression
}

func NewOrExpression(l, r Expression) *OrExpression {
	return &OrExpression{
		Left:  l,
		Right: r}
}
