package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/7_Reverse_Integer/solution"

	"github.com/bmizerany/assert"
)

func TestSolution01(t *testing.T) {
	input := 123
	expected := 321
	output := solution.Reverse(input)
	assert.Equal(t, expected, output)
}

func TestSolution02(t *testing.T) {
	input := -123
	expected := -321
	output := solution.Reverse(input)
	assert.Equal(t, expected, output)
}

func TestSolution03(t *testing.T) {
	input := 120
	expected := 21
	output := solution.Reverse(input)
	assert.Equal(t, expected, output)
}

func TestSolution04(t *testing.T) {
	input := 0
	expected := 0
	output := solution.Reverse(input)
	assert.Equal(t, expected, output)
}

func TestSolution05(t *testing.T) {
	input := 4294967295
	expected := 0
	output := solution.Reverse(input)
	assert.Equal(t, expected, output)
}

func TestSolution06(t *testing.T) {
	input := -1563847412
	expected := 0
	output := solution.Reverse(input)
	assert.Equal(t, expected, output)
}

func TestSolution07(t *testing.T) {
	input := 1563847412
	expected := 0
	output := solution.Reverse(input)
	assert.Equal(t, expected, output)
}
