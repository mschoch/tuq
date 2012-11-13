package parser

import (
	"fmt"
)

type Object map[string]Expression

type ObjectLiteral struct {
	Val Object
}

func NewObjectLiteral(v Object) *ObjectLiteral {
	return &ObjectLiteral{
		Val: v}
}

func (o *ObjectLiteral) AddAll(other *ObjectLiteral) {
	for k, v := range other.Val {
		o.Val[k] = v
	}
}

func (ol *ObjectLiteral) String() string {
	items := ""
	i := 0
	for k, v := range ol.Val {
		if i == 0 {
			items += fmt.Sprintf("\"%s\": %v", k, v)
		} else {
			items += fmt.Sprintf(", \"%s\": %v", k, v)
		}
		i += 1
	}
	return fmt.Sprintf("{%s}", items)
}

func (ol *ObjectLiteral) SybolsReferenced() []string {
	result := make([]string, 0)
	for _, expr := range ol.Val {
		result = concatStringSlices(result, expr.SybolsReferenced())
	}
	return result
}

func (ol *ObjectLiteral) PrefixSymbols(s string) {
    for _, expr := range ol.Val {
        expr.PrefixSymbols(s)
    }
}
