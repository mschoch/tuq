package parser

import (
	"fmt"
	"strings"
)

type Property struct {
	Symbol string
}

func NewProperty(s string) *Property {
	return &Property{
		Symbol: s}
}

func (p *Property) String() string {
	return fmt.Sprintf("%s", p.Symbol)
}

func (p *Property) SymbolsReferenced() []string {
	return []string{p.Symbol}
}

func (p *Property) PrefixSymbols(s string) {
	if !strings.HasPrefix(p.Symbol, s) {
		p.Symbol = s + p.Symbol
	}
}

func (p *Property) HasSubProperty() bool {
	dotIndex := strings.Index(p.Symbol, ".")
	return dotIndex >= 0
}

func (p *Property) Head() string {
	dotIndex := strings.Index(p.Symbol, ".")
	if dotIndex >= 0 {
		return p.Symbol[0:dotIndex]
	}
	return p.Symbol
}

func (p *Property) Tail() *Property {
	dotIndex := strings.Index(p.Symbol, ".")
	if dotIndex >= 0 {
		return NewProperty(p.Symbol[dotIndex+1:])
	}
	return nil
}

func (p *Property) IsReferencingDataSource(ds string) bool {
	if strings.HasPrefix(p.Symbol, ds+".") {
		return true
	}
	return false
}
