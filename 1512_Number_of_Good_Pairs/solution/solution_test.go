package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/1512_Number_of_Good_Pairs/solution"
	"github.com/bmizerany/assert"
	"testing"
)

func TestSolution01(t *testing.T) {
	input := []int{1, 2, 3, 1, 1, 3}
	output := solution.NumIdenticalPairs(input)
	expect := 4

	assert.Equal(t, expect, output)
}

func TestSolution02(t *testing.T) {
	input := []int{1, 1, 1, 1}
	output := solution.NumIdenticalPairs(input)
	expect := 6

	assert.Equal(t, expect, output)
}

func TestSolution03(t *testing.T) {
	input := []int{1, 2, 3}
	output := solution.NumIdenticalPairs(input)
	expect := 0

	assert.Equal(t, expect, output)
}
