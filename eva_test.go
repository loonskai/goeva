package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func recoverable(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("Should not work")
	} else {
		t.Logf("Recovered in Eval: %v", r)
	}
}

func TestEvalValid(t *testing.T) {
	eva := Eva{
		global: Environment{},
	}

	// Self-assigned expressions
	assert.Equal(t, eva.Eval(1, &eva.global), 1)
	assert.Equal(t, eva.Eval(`"hello"`, &eva.global), `hello`)
	// Math operations
	assert.Equal(t, eva.Eval([]any{"+", 1, 5}, &eva.global), 6)
	assert.Equal(t, eva.Eval([]any{"+", 1, 5, 5}, &eva.global), 11)
	assert.Equal(t, eva.Eval([]any{"+", []any{"+", 1, 5}, 5}, &eva.global), 11)
	assert.Equal(t, eva.Eval([]any{"-", 10, 2}, &eva.global), 8)
	assert.Equal(t, eva.Eval([]any{"-", []any{"-", 10, 5}, 5}, &eva.global), 0)
	assert.Equal(t, eva.Eval([]any{"-", []any{"+", 15, 15}, 5}, &eva.global), 25)
	assert.Equal(t, eva.Eval([]any{"*", []any{"*", 10, 5}, 5}, &eva.global), 250)
	assert.Equal(t, eva.Eval([]any{"/", []any{"/", 100, 4}, 5}, &eva.global), 5)
	// Variables
	assert.Equal(t, eva.Eval([]any{"var", "foo", `"hello"`}, &eva.global), `hello`)
	assert.Equal(t, eva.Eval("foo", &eva.global), `hello`)
	assert.Equal(t, eva.Eval([]any{"set", "foo", []any{"+", 2, 3}}, &eva.global), 5)
	assert.Equal(t, eva.Eval("foo", &eva.global), 5)
}

func TestEvalInvalidSingleQuoteString(t *testing.T) {
	eva := Eva{
		global: Environment{},
	}

	defer recoverable(t)

	eva.Eval("'hello'", &eva.global)
}

func TestEvalInvalidSumOfNumAndString(t *testing.T) {
	eva := Eva{
		global: Environment{},
	}

	defer recoverable(t)

	eva.Eval([]any{"+", 1, "hello"}, &eva.global)
}

func TestEvalInvalidNestedSlice(t *testing.T) {
	eva := Eva{
		global: Environment{},
	}

	defer recoverable(t)

	eva.Eval([]any{"+", 2, 1, []any{"+", 1, "hello"}}, &eva.global)
}

func TestEvalInvalidDivisionByZero(t *testing.T) {
	eva := Eva{
		global: Environment{},
	}

	defer recoverable(t)

	eva.Eval([]any{"/", 25, 5, 0}, &eva.global)
}

func TestAccessUndeclaredVariable(t *testing.T) {
	eva := Eva{
		global: Environment{},
	}

	defer recoverable(t)

	eva.Eval([]any{"var", "foo", 5}, &eva.global)
	eva.Eval("bar", &eva.global)
}

func TestAssignUndeclaredVariable(t *testing.T) {
	eva := Eva{
		global: Environment{},
	}

	defer recoverable(t)

	eva.Eval([]any{"set", "foo", 5}, &eva.global)
}

func TestAccessGlobalVariable(t *testing.T) {
	globalEnv := Environment{}
	globalEnv.Define("VERSION", "0.1")
	eva := Eva{
		global: globalEnv,
	}

	assert.Equal(t, eva.Eval("VERSION", nil), "0.1")
}

// ----- BLOCKS -----
func TestEvalBlocks(t *testing.T) {
	eva := Eva{
		global: Environment{},
	}

	assert.Equal(t, eva.Eval([]any{"begin", []any{"var", "x", 10}, []any{"var", "y", 20}, []any{"+", []any{"*", "x", "y"}, 30}}, &eva.global), 230)
}

func TestEvalBlockShadowing(t *testing.T) {
	eva := Eva{
		global: Environment{},
	}

	assert.Equal(t, eva.Eval([]any{"begin",
		[]any{"var", "x", 10},
		[]any{
			"begin",
			[]any{"var", "x", 20},
			[]any{"+", "x", 5},
		},
		"x",
	}, &eva.global), 10)
}

func TestEvalBlockParentLookup(t *testing.T) {
	eva := Eva{
		global: Environment{},
	}

	assert.Equal(t, eva.Eval(
		[]any{
			"begin",
			[]any{"var", "x", 10},
			[]any{
				"var", "result",
				[]any{
					"begin",
					[]any{"var", "y", 20},
					[]any{"var", "sum", []any{"+", "x", "y"}},
					"sum",
				},
			},
			"result",
		}, nil),
		30)
}

func TestEvalParentVariableAssignment(t *testing.T) {
	eva := Eva{
		global: Environment{},
	}

	assert.Equal(t, eva.Eval(
		[]any{
			"begin",
			[]any{"var", "x", 10},
			[]any{
				"begin",
				[]any{"set", "x", 20},
			},
			"x",
		}, nil),
		20)
}

func TestEvalIfExpression(t *testing.T) {
	eva := Eva{
		global: Environment{},
	}

	assert.Equal(t, eva.Eval(
		[]any{
			"begin",
			[]any{"var", "x", 10},
			[]any{"var", "y", 0},
			[]any{
				"if",
				[]any{">", "x", 10},
				[]any{"set", "y", 20},
				[]any{"set", "y", 30},
			},
			"y",
		}, nil), 30)
	assert.Equal(t, eva.Eval(
		[]any{
			"begin",
			[]any{"var", "x", 10},
			[]any{"var", "y", 0},
			[]any{
				"if",
				[]any{">", "x", 5},
				[]any{"set", "y", 20},
				[]any{"set", "y", 30},
			},
			"y",
		}, nil), 20)
	assert.Equal(t, eva.Eval(
		[]any{
			"begin",
			[]any{"var", "x", 10},
			[]any{"var", "y", 0},
			[]any{
				"if",
				[]any{"==", "x", 10},
				[]any{"set", "y", 20},
				[]any{"set", "y", 30},
			},
			"y",
		}, nil), 20)
}
