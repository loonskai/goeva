package main

import (
	"fmt"
)

type Environment struct {
	Record map[string]any
	Parent *Environment
}

func (env *Environment) Define(name string, value any) any {
	if env.Record == nil {
		env.Record = map[string]any{}
	}
	env.Record[name] = value
	return value
}

func (env *Environment) Lookup(name string) any {
	if _, ok := env.Record[name]; !ok {
		panic(fmt.Errorf("cannot access undeclared variable: %v", name))
	}
	return env.Record[name]
}

func (env *Environment) Set(name string, value any) any {
	if _, ok := env.Record[name]; !ok {
		panic(fmt.Errorf("cannot set undeclared variable: %v", name))
	}
	env.Record[name] = value
	return value
}
