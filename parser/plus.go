package parser

import (
    //  "log"
    "fmt"
)

type PlusExpression struct {
    Left  Expression
    Right Expression
}

func NewPlusExpression(l, r Expression) *PlusExpression {
    return &PlusExpression{
        Left:  l,
        Right: r}
}

func (pe *PlusExpression) String() string {
    return fmt.Sprintf("%v + %v", pe.Left, pe.Right)
}

func (pe *PlusExpression) SybolsReferenced() []string {
    leftSymbols := pe.Left.SybolsReferenced()
    return concatStringSlices(leftSymbols, pe.Right.SybolsReferenced())
}

func (pe *PlusExpression) PrefixSymbols(s string) {
    pe.Left.PrefixSymbols(s)
    pe.Right.PrefixSymbols(s)
}