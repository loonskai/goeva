package main

import (
	"reflect"
	"regexp"
)

func isNumber(expression any) bool {
	return reflect.TypeOf(expression).Kind() == reflect.Int
}

func isString(expression any) bool {
	if str, ok := expression.(string); ok {
		return str[0] == '"' && str[len(str)-1] == '"'
	}
	return false
}

func isSlice(expression any) bool {
	return reflect.TypeOf(expression).Kind() == reflect.Slice
}

func isValidMathExpression(expression []any) bool {
	if isSlice(expression) && !isValidMathOperator(expression[0].(string)) {
		return false
	}
	terms := expression[1:]
	for _, term := range terms {
		if isNumber(term) || isValidVariableName(term) {
			continue
		}
		if !isValidMathExpression(term.([]any)) {
			return false
		}
	}
	return true
}

func isValidVariableDeclaration(expression any) bool {
	return isSlice(expression) && len(expression.([]any)) > 2 && expression.([]any)[0] == "var"
}

func isValidVariableAssignement(expression any) bool {
	return isSlice(expression) && len(expression.([]any)) > 2 && expression.([]any)[0] == "set"
}

func isValidVariableName(expression any) bool {
	if str, ok := expression.(string); ok {
		r, err := regexp.Compile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
		if err != nil {
			panic("cannot construct variables format")
		}
		return r.MatchString(str)
	}
	return false
}

func isValidMathOperator(operator string) bool {
	switch operator {
	case "+":
		return true
	case "-":
		return true
	case "/":
		return true
	case "*":
		return true
	}
	return false
}

func isBlock(expression any) bool {
	if !isSlice(expression) {
		return false
	}
	if str, ok := expression.([]any)[0].(string); ok {
		return str == "begin"
	}
	return false
}

func isValidStatement(expression any) bool {
	if !isSlice(expression) {
		return false
	}
	if str, ok := expression.([]any)[0].(string); ok {
		return str == "if"
	}
	return false
}

func isValidConditionalExpression(expression any) bool {
	if !isSlice(expression) {
		return false
	}
	if len(expression.([]any)) < 3 {
		return false
	}
	if str, ok := expression.([]any)[0].(string); ok {
		return str == ">" || str == "<" || str == "=="
	}
	return false
}

func isValidWhileStatement(expression any) bool {
	if !isSlice(expression) {
		return false
	}
	if len(expression.([]any)) < 3 {
		return false
	}
	if str, ok := expression.([]any)[0].(string); ok {
		return str == "while"
	}
	return false
}
