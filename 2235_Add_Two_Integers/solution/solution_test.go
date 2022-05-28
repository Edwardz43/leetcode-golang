package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/2235_Add_Two_Integers/solution"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution01(t *testing.T) {
	input := []int{12, 5}
	expected := 17
	output := solution.Sum(input[0], input[1])
	assert.Equal(t, expected, output)
}
func TestSolution02(t *testing.T) {
	input := []int{-10, 4}
	expected := -6
	output := solution.Sum(input[0], input[1])
	assert.Equal(t, expected, output)
}
