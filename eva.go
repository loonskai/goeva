package main

import (
	"errors"
	"reflect"
	"regexp"
)

type Eva struct {
	global Environment
}

func (eva *Eva) Eval(expression any, env *Environment) any {
	if env == nil {
		env = &eva.global
	}
	if isNumber(expression) {
		return expression
	}

	if isString(expression) {
		strExpression := expression.(string)
		return strExpression[1 : len(strExpression)-1]
	}

	if isVariableName(expression) {
		return env.Lookup(expression.(string))
	}

	if isValidMathExpression(expression.([]any), "+") {
		return eva.add(expression.([]any), env)
	}

	if isValidMathExpression(expression.([]any), "-") {
		return eva.subtract(expression.([]any), env)
	}

	if isValidMathExpression(expression.([]any), "*") {
		return eva.multiply(expression.([]any), env)
	}

	if isValidMathExpression(expression.([]any), "/") {
		return eva.divide(expression.([]any), env)
	}

	if isValidVariableDeclaration(expression) {
		return env.Define(expression.([]any)[1].(string), eva.Eval(expression.([]any)[2], nil))
	}

	if isValidVariableAssignement(expression.([]any)) {
		return env.Set(expression.([]any)[1].(string), eva.Eval(expression.([]any)[2], nil))
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

func isValidVariableDeclaration(expression any) bool {
	return isSlice(expression) && len(expression.([]any)) > 2 && expression.([]any)[0] == "var"
}

func isValidVariableAssignement(expression any) bool {
	return isSlice(expression) && len(expression.([]any)) > 2 && expression.([]any)[0] == "set"
}

func isVariableName(expression any) bool {
	if str, ok := expression.(string); ok {
		r, err := regexp.Compile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
		if err != nil {
			panic("cannot construct variables format")
		}
		return r.MatchString(str)
	}
	return false
}

// ------ operations ----
// TODO: Replace recursion with iteration
func (eva *Eva) add(expression []any, env *Environment) int {
	sum := 0
	nums := expression[1:]
	for _, n := range nums {
		sum += eva.Eval(n, env).(int)
	}
	return sum
}

func (eva *Eva) subtract(expression []any, env *Environment) int {
	sum := 0
	nums := expression[1:]
	for i, n := range nums {
		if i == 0 {
			sum += eva.Eval(n, env).(int)
		} else {
			sum -= eva.Eval(n, env).(int)
		}
	}
	return sum
}
func (eva *Eva) multiply(expression []any, env *Environment) int {
	product := 1
	nums := expression[1:]
	for _, n := range nums {
		product *= eva.Eval(n, env).(int)
	}
	return product
}

func (eva *Eva) divide(expression []any, env *Environment) int {
	quotinent := 1
	nums := expression[1:]
	for i, n := range nums {
		if i == 0 {
			quotinent = eva.Eval(n, env).(int)
		} else {
			if n == 0 {
				panic(errors.New("can't divide by zero"))
			}
			quotinent /= eva.Eval(n, env).(int)
		}
	}
	return quotinent
}
