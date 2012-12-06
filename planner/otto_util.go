package planner

import (
	"encoding/json"
	"fmt"
	"github.com/mschoch/tuq/parser"
	"github.com/robertkrimen/otto"
	"log"
)

// serialize all the top level elements of doc to JSON and
// put them in the environment
// FIXME there has to be a better way than this
func putDocumentIntoEnvironment(o *otto.Otto, doc Document) {
	for k, v := range doc {
		v_json, err := json.Marshal(v)
		if err != nil {
			log.Printf("JSON serialization failed: %v", err)
		} else {
			_, err := o.Run(k + "=" + string(v_json))
			if err != nil {
				log.Printf("Error running otto put: %v", err)
			}
		}
	}

}

// now we be sure to cleanup the environment as best we can
func cleanupDocumentFromEnvironment(o *otto.Otto, doc Document) {
	for k, _ := range doc {
		_, err := o.Run(k + "=undefined")
		if err != nil {
			log.Printf("Error running otto cleanup: %v", err)
		}
	}
}

func evaluateExpressionInEnvironment(o *otto.Otto, expr parser.Expression) otto.Value {
	result, err := o.Run(fmt.Sprintf("ignore = %v", expr))
	if err != nil {
		log.Printf("Error running otto eval %v, %v", expr, err)
	} else {
		return result
	}
	return otto.UndefinedValue()
}

func evaluateExpressionInEnvironmentAsBoolean(o *otto.Otto, expr parser.Expression) bool {
	expression := fmt.Sprintf("%v", expr)
	return evaluateExpressionStringInEnvironmentAsBoolean(o, expression)
}

func evaluateExpressionStringInEnvironmentAsBoolean(o *otto.Otto, expression string) bool {
	result, err := o.Run(expression)
	if err != nil {
		log.Printf("Error running otto %v, %v", expression, err)
	} else {
		//log.Printf("result was %v", result)
		result, err := result.ToBoolean()
		if err != nil {
			log.Printf("Error converting otto result to boolean %v", err)
		} else if result {
			return true
		}
	}
	return false
}

func evaluateExpressionInEnvironmentAsInteger(o *otto.Otto, expr parser.Expression) int64 {
	result, err := o.Run(fmt.Sprintf("%v", expr))
	if err != nil {
		log.Printf("Error running otto %v", err)
	} else {
		result, err := result.ToInteger()
		if err != nil {
			log.Printf("Error converting otto result to integer %v", err)
		} else {
			return result
		}
	}
	return -1
}

func convertToPrimitive(v otto.Value) interface{} {
	if v.IsBoolean() {
		v_b, err := v.ToBoolean()
		if err != nil {
			log.Printf("Error converting to boolean")
		}
		return v_b
	} else if v.IsNumber() {
		v_f, err := v.ToFloat()
		if err != nil {
			log.Printf("Error converting to float")
		}
		return v_f
	} else {
		v_s, err := v.ToString()
		if err != nil {
			log.Printf("Error converting to boolean")
		}
		return v_s
	}
	return nil
}
