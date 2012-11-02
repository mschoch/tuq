package main

import (
	"github.com/mschoch/go-unql-couchbasev2/parser"
	"strings"
)

func EvaluateExpressionInContext(expression parser.Expression, context map[string]interface{}) (interface{}, error) {

	var result interface{} = nil
	var err error = nil
	switch exp := expression.(type) {
	case parser.Property:
		result = ValueOfPropertyInContext(exp.Symbol, context)
	case parser.ObjectLiteral:
		obj := make(map[string]interface{})
		for k, v := range exp.Val {
			inner_expr, err := EvaluateExpressionInContext(v, context)
			if err != nil {
				return nil, err
			} else {
				obj[k] = inner_expr
				result = obj
			}
		}
	case parser.ArrayLiteral:
		obj := make([]interface{}, 0)
		for _, v := range exp.Val {
			inner_expr, err := EvaluateExpressionInContext(v, context)
			if err != nil {
				return nil, err
			} else {
				obj = append(obj, inner_expr)
				result = obj
			}
		}
	case parser.Function:
		expr_list := exp.Args
		first_arg := expr_list[0].(parser.Expression)

		switch fa := first_arg.(type) {
		case parser.Property:
			result = ValueOfPropertyInContext(FunctionPrefix+exp.Name+"."+fa.Symbol, context)
		}

	case parser.StringLiteral:
		result = exp.Val
	case parser.IntegerLiteral:
		result = exp.Val
	case parser.FloatLiteral:
		result = exp.Val
	case parser.BoolLiteral:
		result = exp.Val
	}

	return result, err
}

// walk through the map (recursively)
// finding any properties or identifiers
// this is used to peek at the from clause
// and find any fields we need stats for (other than the aggregate column)
func FindPropertiesAndIdentifiers(expression parser.Expression) []string {

	var result []string = make([]string, 0)

	switch exp := expression.(type) {
	case parser.Property:
		result = []string{exp.Symbol}
	case parser.ObjectLiteral:
		obj := make([]string, 0)
		for _, v := range exp.Val {
			inner_expr := FindPropertiesAndIdentifiers(v)
			if inner_expr != nil {
				obj = AppendStringSliceIfMissing(obj, inner_expr)
				result = obj
			}
		}
	case parser.ArrayLiteral:
		obj := make([]string, 0)
		for _, v := range exp.Val {
			inner_expr := FindPropertiesAndIdentifiers(v)
			if inner_expr != nil {
				obj = AppendStringSliceIfMissing(obj, inner_expr)
				result = obj
			}
		}
	case parser.Function:
		expr_list := exp.Args
		first_arg := expr_list[0].(parser.Expression)

		switch fa := first_arg.(type) {
		case parser.Property:
			result = []string{fa.Symbol}
		}
	}

	return result
}

func ValueOfPropertyInContext(property string, context map[string]interface{}) interface{} {

	dotIndex := strings.Index(property, ".")
	if dotIndex > 0 {
		curr := property[0:dotIndex]
		remaining := property[dotIndex+1:]
		innerContext, ok := context[curr]
		if ok {
			return ValueOfPropertyInContext(remaining, innerContext.(map[string]interface{}))
		} else {
			return nil
		}
	} else {
		val, ok := context[property]
		if ok {
			return val
		} else {
			return nil
		}
	}

	return nil
}
