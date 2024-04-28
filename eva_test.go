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
	eva := Eva{}

	assert.Equal(t, eva.Eval(1), 1)
	assert.Equal(t, eva.Eval(`"hello"`), `hello`)
	assert.Equal(t, eva.Eval([]any{"+", 1, 5}), 6)
	assert.Equal(t, eva.Eval([]any{"+", 1, 5, 5}), 11)
	assert.Equal(t, eva.Eval([]any{"+", []any{"+", 1, 5}, 5}), 11)
	assert.Equal(t, eva.Eval([]any{"-", 10, 2}), 8)
	assert.Equal(t, eva.Eval([]any{"-", []any{"-", 10, 5}, 5}), 0)
	// assert.Equal(t, eva.Eval([]any{"-", []any{"+", 15, 15}, 5}), 25) // FIXME
	assert.Equal(t, eva.Eval([]any{"*", []any{"*", 10, 5}, 5}), 250)
	assert.Equal(t, eva.Eval([]any{"/", []any{"/", 100, 4}, 5}), 5)

}

func TestEvalInvalidSingleQuoteString(t *testing.T) {
	eva := Eva{}

	defer recoverable(t)

	eva.Eval("'hello'")
}

func TestEvalInvalidSumOfNumAndString(t *testing.T) {
	eva := Eva{}

	defer recoverable(t)

	eva.Eval([]any{"+", 1, "hello"})
}

func TestEvalInvalidNestedSlice(t *testing.T) {
	eva := Eva{}

	defer recoverable(t)

	eva.Eval([]any{"+", 2, 1, []any{"+", 1, "hello"}})
}
