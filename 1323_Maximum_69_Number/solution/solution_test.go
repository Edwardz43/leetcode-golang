package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/1323_Maximum_69_Number/solution"

	"github.com/bmizerany/assert"
)

func TestSolution01(t *testing.T) {
	input := 9669
	expected := 9969
	output := solution.Maximum69Number(input)
	assert.Equal(t, expected, output)
}

func TestSolution02(t *testing.T) {
	input := 9996
	expected := 9999
	output := solution.Maximum69Number(input)
	assert.Equal(t, expected, output)
}

func TestSolution03(t *testing.T) {
	input := 9999
	expected := 9999
	output := solution.Maximum69Number(input)
	assert.Equal(t, expected, output)
}

func TestSolution04(t *testing.T) {
	input := 6996
	expected := 9996
	output := solution.Maximum69Number(input)
	assert.Equal(t, expected, output)
}

func TestSolution05(t *testing.T) {
	input := 66
	expected := 96
	output := solution.Maximum69Number(input)
	assert.Equal(t, expected, output)
}
