package main

import "errors"

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
