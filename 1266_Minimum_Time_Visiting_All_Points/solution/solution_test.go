package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/1266_Minimum_Time_Visiting_All_Points/solution"
	"github.com/bmizerany/assert"
	"testing"
)

func TestSolution01(t *testing.T) {
	input := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	expected := 7
	output := solution.MinTimeToVisitAllPoints(input)
	assert.Equal(t, expected, output)
}

func TestSolution02(t *testing.T) {
	input := [][]int{{3, 2}, {-2, 2}}
	expected := 5
	output := solution.MinTimeToVisitAllPoints(input)
	assert.Equal(t, expected, output)
}

func TestSolution03(t *testing.T) {
	input := [][]int{{0, 0}, {-1, 3}, {2, 2}}
	expected := 6
	output := solution.MinTimeToVisitAllPoints(input)
	assert.Equal(t, expected, output)
}
