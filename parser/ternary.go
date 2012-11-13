package parser

import (
	//	"log"
	"fmt"
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

func (te *TernaryExpression) String() string {
	return fmt.Sprintf("%v ? %v : %v", te.Iff, te.Thenn, te.Elsee)
}

func (te *TernaryExpression) SybolsReferenced() []string {
	iffSymbols := te.Iff.SybolsReferenced()
	thennSymbols := concatStringSlices(iffSymbols, te.Thenn.SybolsReferenced())
	return concatStringSlices(thennSymbols, te.Elsee.SybolsReferenced())
}

func (te *TernaryExpression) PrefixSymbols(s string) {
    te.Iff.PrefixSymbols(s)
    te.Thenn.PrefixSymbols(s)
    te.Elsee.PrefixSymbols(s)
}
