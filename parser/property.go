package parser

import ()

type Property struct {
	Symbol string
}

func NewProperty(s string) *Property {
	return &Property{
		Symbol: s}
}
