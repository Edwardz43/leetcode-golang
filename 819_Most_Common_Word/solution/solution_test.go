package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/819_Most_Common_Word/solution"
	"github.com/bmizerany/assert"
)

func TestSolution01(t *testing.T) {
	input := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	expected := "ball"
	output := solution.MostCommonWord(input, banned)
	assert.Equal(t, expected, output)
}

func TestSolution02(t *testing.T) {
	input := "a."
	banned := []string{}
	expected := "a"
	output := solution.MostCommonWord(input, banned)
	assert.Equal(t, expected, output)
}
