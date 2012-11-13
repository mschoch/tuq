package parser

import (
    //  "log"
    "fmt"
)

type MinusExpression struct {
    Left  Expression
    Right Expression
}

func NewMinusExpression(l, r Expression) *MinusExpression {
    return &MinusExpression{
        Left:  l,
        Right: r}
}

func (me *MinusExpression) String() string {
    return fmt.Sprintf("%v - %v", me.Left, me.Right)
}

func (me *MinusExpression) SybolsReferenced() []string {
    leftSymbols := me.Left.SybolsReferenced()
    return concatStringSlices(leftSymbols, me.Right.SybolsReferenced())
}

func (me *MinusExpression) PrefixSymbols(s string) {
    me.Left.PrefixSymbols(s)
    me.Right.PrefixSymbols(s)
}
