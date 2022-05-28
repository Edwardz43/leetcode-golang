package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/268_Missing_Number/solution"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution01(t *testing.T) {
	input := []int{3, 0, 1}
	expected := 2
	output := solution.MissingNumber(input)
	assert.Equal(t, expected, output)
}
func TestSolution02(t *testing.T) {
	input := []int{0, 1}
	expected := 2
	output := solution.MissingNumber(input)
	assert.Equal(t, expected, output)
}
func TestSolution03(t *testing.T) {
	input := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	expected := 8
	output := solution.MissingNumber(input)
	assert.Equal(t, expected, output)
}

var in = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

func BenchmarkMissingNumber1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution.MissingNumber(in)
	}
}
func BenchmarkMissingNumber2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution.MissingNumber2(in)
	}
}
