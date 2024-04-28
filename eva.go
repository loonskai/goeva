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
		return add(expression.([]any))
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
		if !isNumber(n) && !isValidAddition(n.([]any)) {
			return false
		}
	}
	return true
}

// ------ operations ----
// TODO: Replace recursion with iteration
func add(expression []any) int {
	sum := 0
	addends := expression[1:]
	for _, n := range addends {
		if isSlice(n) {
			sum += add(n.([]any))
		} else {
			sum += n.(int)
		}
	}
	return sum
}
