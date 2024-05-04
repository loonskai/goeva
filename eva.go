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

	if isValidVariableDeclaration(expression) {
		return env.Define(expression.([]any)[1].(string), eva.Eval(expression.([]any)[2], env))
	}

	if isValidStatement(expression) {
		result := eva.Eval(expression.([]any)[1], env)
		if result == true {
			return eva.Eval(expression.([]any)[2], env)
		}
		if expression.([]any)[3] != nil {
			return eva.Eval(expression.([]any)[3], env)
		}
	}

	if isValidWhileStatement(expression) {
		var result any
		for eva.Eval(expression.([]any)[1], env).(bool) {
			result = eva.Eval(expression.([]any)[2], env)
		}
		return result
	}

	if isValidVariableName(expression) {
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

	if isValidConditionalExpression(expression.([]any)) {
		return eva.compare(expression.([]any)[0].(string), eva.Eval(expression.([]any)[1], env), eva.Eval(expression.([]any)[2], env))
	}

	if isValidVariableAssignement(expression.([]any)) {
		return env.Set(expression.([]any)[1].(string), eva.Eval(expression.([]any)[2], env))
	}

	panic(errors.New("invalid"))
}

func (eva *Eva) evalBlock(expression []any, env *Environment) any {
	blockEnv := Environment{
		Parent: env,
	}
	var result any
	expressions := expression[1:]
	for _, e := range expressions {
		result = eva.Eval(e, &blockEnv)
	}
	return result
}

func (eva *Eva) compare(operator string, x any, y any) bool {
	if !isNumber(x) || !isNumber(y) {
		return false
	}
	switch operator {
	case ">":
		return x.(int) > y.(int)
	case "<":
		return x.(int) < y.(int)
	case "==":
		return x.(int) == y.(int)
	}
	return false
}
