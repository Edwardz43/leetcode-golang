package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/1486_XOR_Operation_in_an_Array/solution"
	"github.com/bmizerany/assert"
	"testing"
)

func TestSolution01(t *testing.T) {
	n := 5
	start := 0
	expected := 8
	output := solution.XorOperation(n, start)
	assert.Equal(t, expected, output)
}

func TestSolution02(t *testing.T) {
	n := 4
	start := 3
	expected := 8
	output := solution.XorOperation(n, start)
	assert.Equal(t, expected, output)
}

func TestSolution03(t *testing.T) {
	n := 1
	start := 7
	expected := 7
	output := solution.XorOperation(n, start)
	assert.Equal(t, expected, output)
}

func TestSolution04(t *testing.T) {
	n := 10
	start := 5
	expected := 2
	output := solution.XorOperation(n, start)
	assert.Equal(t, expected, output)
}
