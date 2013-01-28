package parser

import (
	"fmt"
)

func ProcessPragma(left, right Expression) {

	switch l := left.(type) {
	case *Property:
		if l.Symbol == "debugTokens" {
			switch r := right.(type) {
			case *BoolLiteral:
				DebugTokens = r.Val
			default:
				fmt.Printf("Pragma debugTokens only supports boolean value")
			}
		} else if l.Symbol == "debugGrammar" {
			switch r := right.(type) {
			case *BoolLiteral:
				DebugGrammar = r.Val
			default:
				fmt.Printf("Pragma debugGrammar only supports boolean value")
			}
		}

	default:
		fmt.Printf("Unsupported expression in pragma: %v", left)
	}

}
