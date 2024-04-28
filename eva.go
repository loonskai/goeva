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

	if isValidMathExpression(expression.([]any), "+") {
		return eva.add(expression.([]any))
	}

	if isValidMathExpression(expression.([]any), "-") {
		return eva.subtract(expression.([]any))
	}

	if isValidMathExpression(expression.([]any), "*") {
		return eva.multiply(expression.([]any))
	}

	if isValidMathExpression(expression.([]any), "/") {
		return eva.divide(expression.([]any))
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

func isValidMathExpression(expression []any, operationSymbol string) bool {
	if !isSlice(expression) || expression[0] != operationSymbol {
		return false
	}
	nums := expression[1:]
	for _, n := range nums {
		if !isNumber(n) && !isValidMathExpression(n.([]any), operationSymbol) {
			return false
		}
	}
	return true
}

// ------ operations ----
// TODO: Replace recursion with iteration
func (eva *Eva) add(expression []any) int {
	sum := 0
	nums := expression[1:]
	for _, n := range nums {
		sum += eva.Eval(n).(int)
	}
	return sum
}

func (eva *Eva) subtract(expression []any) int {
	sum := 0
	nums := expression[1:]
	for i, n := range nums {
		if i == 0 {
			sum += eva.Eval(n).(int)
		} else {
			sum -= eva.Eval(n).(int)
		}
	}
	return sum
}
func (eva *Eva) multiply(expression []any) int {
	product := 1
	nums := expression[1:]
	for _, n := range nums {
		product *= eva.Eval(n).(int)
	}
	return product
}

func (eva *Eva) divide(expression []any) int {
	quotinent := 1
	nums := expression[1:]
	for i, n := range nums {
		if i == 0 {
			quotinent = eva.Eval(n).(int)
		} else {
			if n == 0 {
				panic(errors.New("can't divide by zero"))
			}
			quotinent /= eva.Eval(n).(int)
		}
	}
	return quotinent
}
