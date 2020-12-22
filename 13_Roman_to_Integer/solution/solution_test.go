package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/13_Roman_to_Integer/solution"
	"github.com/bmizerany/assert"
)

func TestSolution01(t *testing.T) {
	input := "III"
	expected := 3
	output := solution.RomanToInt(input)
	assert.Equal(t, expected, output)
}

func TestSolution02(t *testing.T) {
	input := "IV"
	expected := 4
	output := solution.RomanToInt(input)
	assert.Equal(t, expected, output)
}

func TestSolution03(t *testing.T) {
	input := "IX"
	expected := 9
	output := solution.RomanToInt(input)
	assert.Equal(t, expected, output)
}

func TestSolution04(t *testing.T) {
	input := "LVIII"
	expected := 58
	output := solution.RomanToInt(input)
	assert.Equal(t, expected, output)
}

func TestSolution05(t *testing.T) {
	input := "MCMXCIV"
	expected := 1994
	output := solution.RomanToInt(input)
	assert.Equal(t, expected, output)
}

func TestSolution07(t *testing.T) {
	input := "DXL"
	expected := 540
	output := solution.RomanToInt(input)
	assert.Equal(t, expected, output)
}

func TestSolution06(t *testing.T) {
	input := "MCDXCIV"
	expected := 1494
	output := solution.RomanToInt(input)
	assert.Equal(t, expected, output)
}
