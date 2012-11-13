package parser

import (
	"fmt"
)

type Function struct {
	Name string
	Args ExpressionList
}

func NewFunction(n string) *Function {
	return &Function{
		Name: n}
}

func (f *Function) AddArguments(v ExpressionList) {
	f.Args = v
}

func (f *Function) String() string {
	items := ""
	for i, v := range f.Args {
		if i == 0 {
			items += fmt.Sprintf("%v", v)
		} else {
			items += fmt.Sprintf(",%v", v)
		}
	}
	return fmt.Sprintf("__func__.%s.%s", f.Name, items)
}

func (f *Function) SybolsReferenced() []string {
	result := make([]string, 0)
	for _, expr := range f.Args {
		result = concatStringSlices(result, expr.SybolsReferenced())
	}
	return result
}

func (f *Function) PrefixSymbols(s string) {
	for _, expr := range f.Args {
		expr.PrefixSymbols(s)
	}
}
