package parser

import ()

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