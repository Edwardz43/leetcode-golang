package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/867_Transpose_Matrix/solution"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution01(t *testing.T) {
	input := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	expected := [][]int{
		[]int{1, 4, 7},
		[]int{2, 5, 8},
		[]int{3, 6, 9},
	}
	output := solution.Transpose(input)
	assert.Equal(t, expected, output)
}
func TestSolution02(t *testing.T) {
	input := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}
	expected := [][]int{
		[]int{1, 4},
		[]int{2, 5},
		[]int{3, 6},
	}
	output := solution.Transpose(input)
	assert.Equal(t, expected, output)
}

func TestSolution03(t *testing.T) {
	input := [][]int{
		[]int{1, 2},
		[]int{3, 4},
		[]int{5, 6},
	}
	expected := [][]int{
		[]int{1, 3, 5},
		[]int{2, 4, 6},
	}
	output := solution.Transpose(input)
	assert.Equal(t, expected, output)
}
