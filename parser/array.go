package parser

import (
	"fmt"
)

type ArrayLiteral struct {
	Val ExpressionList
}

func NewArrayLiteral(v ExpressionList) *ArrayLiteral {
	return &ArrayLiteral{
		Val: v}
}

func (al *ArrayLiteral) String() string {
	items := ""
	for i, v := range al.Val {
		if i == 0 {
			items += fmt.Sprintf("%v", v)
		} else {
			items += fmt.Sprintf(",%v", v)
		}
	}
	return fmt.Sprintf("[%s]", items)
}

func (al *ArrayLiteral) SybolsReferenced() []string {
	result := make([]string, 0)
	for _, expr := range al.Val {
		result = concatStringSlices(result, expr.SybolsReferenced())
	}
	return result
}

func (al *ArrayLiteral) PrefixSymbols(s string) {
    for _, expr := range al.Val {
        expr.PrefixSymbols(s)
    }
}
