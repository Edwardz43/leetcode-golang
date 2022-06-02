package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/12_Integer_to_Roman/solution"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution01(t *testing.T) {
	input := 3
	expected := "III"
	output := solution.IntToRoman(input)
	assert.Equal(t, expected, output)
}
func TestSolution02(t *testing.T) {
	input := 58
	expected := "LVIII"
	output := solution.IntToRoman(input)
	assert.Equal(t, expected, output)
}
func TestSolution03(t *testing.T) {
	input := 1994
	expected := "MCMXCIV"
	output := solution.IntToRoman(input)
	assert.Equal(t, expected, output)
}

var input = 1994

func BenchmarkIntToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution.IntToRoman(input)
	}
}
func BenchmarkIntToRoman2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution.IntToRoman2(input)
	}
}
