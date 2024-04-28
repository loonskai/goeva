package main

import (
	"errors"
	"reflect"
)

type Eva struct{}

func (eva *Eva) Eval(expression any) any {
	if isNumber(expression) {
		return expression
	}

	if isString(expression) {
		strExpression := expression.(string)
		return strExpression[1 : len(strExpression)-1]
	}

	if isValidAddition(expression.([]any)) {
		sum := 0
		addends := expression.([]any)[1:]
		for _, n := range addends {
			sum += n.(int)
		}
		return sum
	}

	panic(errors.New("invalid"))
}

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

func isValidAddition(expression []any) bool {
	if !isSlice(expression) || expression[0] != "+" || len(expression) < 3 {
		return false
	}
	addends := expression[1:]
	for _, n := range addends {
		if !isNumber(n) {
			return false
		}
	}
	return true
}
