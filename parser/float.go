package parser

import (
	"fmt"
)

type FloatLiteral struct {
	Val float64
}

func NewFloatLiteral(v float64) *FloatLiteral {
	return &FloatLiteral{
		Val: v}
}

func (fl *FloatLiteral) String() string {
	return fmt.Sprintf("%f", fl.Val)
}

func (fl *FloatLiteral) SymbolsReferenced() []string {
	return []string{}
}

func (fl *FloatLiteral) PrefixSymbols(string) {

}
