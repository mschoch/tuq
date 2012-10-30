package main

import (
	"strings"
)

func EvaluateExpressionInContext(expression, context map[string]interface{}) (interface{}, error) {

	var result interface{} = nil
	var err error = nil

	_, isOperation := expression["op"]

	if isOperation {

	} else {

		val, isProperty := expression["property"]
		if isProperty {
			result = ValueOfPropertyInContext(val.(string), context)
		}

		val, isIdentifier := expression["identifier"]
		if isIdentifier {
			result = ValueOfPropertyInContext(val.(string), context)
		}

		// literals (with nested objects)

		val, isObject := expression["object"]
		if isObject {
			obj := make(map[string]interface{})
			val_map := val.(map[string]interface{})
			for k, v := range val_map {
				inner_expr, err := EvaluateExpressionInContext(v.(map[string]interface{}), context)
				if err != nil {
					return nil, err
				} else {
					obj[k] = inner_expr
					result = obj
				}
			}
		}

		val, isArray := expression["array"]
		if isArray {
			obj := make([]interface{}, 0)
			val_map := val.([]interface{})
			for _, v := range val_map {
				inner_expr, err := EvaluateExpressionInContext(v.(map[string]interface{}), context)
				if err != nil {
					return nil, err
				} else {
					obj = append(obj, inner_expr)
					result = obj
				}
			}
		}

		// functions
		val, isFunction := expression["function"]
		if isFunction {
			expr_list := expression["expression_list"].([]interface{})
			first_arg := expr_list[0].(map[string]interface{})

			prop, isFirstArgProperty := first_arg["property"]
			if isFirstArgProperty {
				result = ValueOfPropertyInContext(FunctionPrefix+val.(string)+"."+prop.(string), context)
			}
		}

		// literals that are terminal

		val, isString := expression["string"]
		if isString {
			result = val
		}

		val, isInt := expression["int"]
		if isInt {
			result = val
		}
		val, isReal := expression["real"]
		if isReal {
			result = val
		}
		val, isBool := expression["bool"]
		if isBool {
			result = val
		}

	}

	return result, err
}

// walk through the map (recursively)
// finding any properties or identifiers
// this is used to peek at the from clause
// and find any fields we need stats for (other than the aggregate column)
func FindPropertiesAndIdentifiers(expression map[string]interface{}) []string {

	var result []string = make([]string, 0)

	_, isOperation := expression["op"]

	if isOperation {

	} else {

		val, isProperty := expression["property"]
		if isProperty {
			result = []string{val.(string)}
		}

		val, isIdentifier := expression["identifier"]
		if isIdentifier {
			result = []string{val.(string)}
		}

		// literals (with nested objects)

		val, isObject := expression["object"]
		if isObject {
			obj := make([]string, 0)
			val_map := val.(map[string]interface{})
			for _, v := range val_map {
				inner_expr := FindPropertiesAndIdentifiers(v.(map[string]interface{}))
				if inner_expr != nil {
					obj = AppendStringSliceIfMissing(obj, inner_expr)
					result = obj
				}
			}
		}

		val, isArray := expression["array"]
		if isArray {
			obj := make([]string, 0)
			val_map := val.([]interface{})
			for _, v := range val_map {
				inner_expr := FindPropertiesAndIdentifiers(v.(map[string]interface{}))
				if inner_expr != nil {
					obj = AppendStringSliceIfMissing(obj, inner_expr)
					result = obj
				}
			}
		}

		// functions
		val, isFunction := expression["function"]
		if isFunction {
			expr_list := expression["expression_list"].([]interface{})
			first_arg := expr_list[0].(map[string]interface{})

			prop, isFirstArgProperty := first_arg["property"]
			if isFirstArgProperty {
				result = []string{prop.(string)}
			}
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
