package parser

import ()

type Null struct {
}

func NewNull() *Null {
	return &Null{}
}

func (n *Null) String() string {
	return "null"
}

func (n *Null) SymbolsReferenced() []string {
	return []string{}
}

func (n *Null) PrefixSymbols(string) {
}
