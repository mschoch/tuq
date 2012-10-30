package main

func ProcessPragma(left, right map[string]interface{}) {
	name, isPragma := left["pragma"]
	val, isBoolean := right["bool"]
	if isPragma && name == "debug" && isBoolean {
		*debugTokens = val.(bool)
		*debugGrammar = val.(bool)
	}

}
