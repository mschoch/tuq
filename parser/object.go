package parser

import ()

type Object map[string]interface{}

type ObjectLiteral struct {
	Val Object
}

func NewObjectLiteral(v Object) *ObjectLiteral {
	return &ObjectLiteral{
		Val: v}
}

func (o *ObjectLiteral) AddAll(other ObjectLiteral) {
	for k, v := range other.Val {
		o.Val[k] = v
	}
}
