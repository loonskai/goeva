package main

import (
	"errors"
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

	if isBlock(expression) {
		return eva.evalBlock(expression.([]any), env)
	}

	if isVariableName(expression) {
		return env.Lookup(expression.(string))
	}

	if isValidMathExpression(expression.([]any)) {
		if expression.([]any)[0].(string) == "+" {
			return eva.add(expression.([]any), env)
		}
		if expression.([]any)[0].(string) == "-" {
			return eva.subtract(expression.([]any), env)
		}
		if expression.([]any)[0].(string) == "/" {
			return eva.divide(expression.([]any), env)
		}
		if expression.([]any)[0].(string) == "*" {
			return eva.multiply(expression.([]any), env)
		}
	}

	if isValidVariableDeclaration(expression) {
		return env.Define(expression.([]any)[1].(string), eva.Eval(expression.([]any)[2], nil))
	}

	if isValidVariableAssignement(expression.([]any)) {
		return env.Set(expression.([]any)[1].(string), eva.Eval(expression.([]any)[2], nil))
	}

	panic(errors.New("invalid"))
}

func (eva *Eva) evalBlock(expression []any, env *Environment) any {
	var result any
	expressions := expression[1:]
	for _, e := range expressions {
		result = eva.Eval(e, env)
	}
	return result
}
