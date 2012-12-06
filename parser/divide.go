package parser

import (
	//	"log"
	"fmt"
)

type DivideExpression struct {
	Left  Expression
	Right Expression
}

func NewDivideExpression(l, r Expression) *DivideExpression {
	return &DivideExpression{
		Left:  l,
		Right: r}
}

func (de *DivideExpression) String() string {
	return fmt.Sprintf("%v / %v", de.Left, de.Right)
}

func (de *DivideExpression) SymbolsReferenced() []string {
	leftSymbols := de.Left.SymbolsReferenced()
	return concatStringSlices(leftSymbols, de.Right.SymbolsReferenced())
}

func (de *DivideExpression) PrefixSymbols(s string) {
	de.Left.PrefixSymbols(s)
	de.Right.PrefixSymbols(s)
}
