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

func (me *MinusExpression) SymbolsReferenced() []string {
    leftSymbols := me.Left.SymbolsReferenced()
    return concatStringSlices(leftSymbols, me.Right.SymbolsReferenced())
}

func (me *MinusExpression) PrefixSymbols(s string) {
    me.Left.PrefixSymbols(s)
    me.Right.PrefixSymbols(s)
}
