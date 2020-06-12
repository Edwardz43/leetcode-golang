package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/1470_Shuffle_the_Array/solution"
	"github.com/bmizerany/assert"
)

func TestShuffle01(t *testing.T) {
	input := []int{2, 5, 1, 3, 4, 7}
	n := 3
	output := solution.Shuffle(input, n)

	expexted := []int{2, 3, 5, 4, 1, 7}
	assert.Equal(t, expexted, output)
}

func TestShuffle02(t *testing.T) {
	input := []int{1, 2, 3, 4, 4, 3, 2, 1}
	n := 4
	output := solution.Shuffle(input, n)

	expexted := []int{1, 4, 2, 3, 3, 2, 4, 1}
	assert.Equal(t, expexted, output)
}

func TestShuffle03(t *testing.T) {
	input := []int{1, 1, 2, 2}
	n := 2
	output := solution.Shuffle(input, n)

	expexted := []int{1, 2, 1, 2}
	assert.Equal(t, expexted, output)
}
