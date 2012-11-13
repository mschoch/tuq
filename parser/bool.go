package parser

import (
	"fmt"
)

type BoolLiteral struct {
	Val bool
}

func NewBoolLiteral(v bool) *BoolLiteral {
	return &BoolLiteral{
		Val: v}
}

func (bl *BoolLiteral) String() string {
	return fmt.Sprintf("%t", bl.Val)
}

func (bl *BoolLiteral) SybolsReferenced() []string {
	return []string{}
}

func (bl *BoolLiteral) PrefixSymbols(string) {
    
}
