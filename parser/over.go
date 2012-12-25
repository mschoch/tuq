package parser

import ()

type Over struct {
	Path *Property
	As   string
}

func NewOver(path *Property, as string) *Over {
	return &Over{
		Path: path,
		As:   as}
}
