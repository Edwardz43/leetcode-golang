package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/1480_Running_Sum_of1d_Array/solution"
	"github.com/bmizerany/assert"
)

/*
	Example 1:
	Input: nums = [1,2,3,4]
	Output: [1,3,6,10]


	Example 2:
	Input: nums = [1,1,1,1,1]
	Output: [1,2,3,4,5]

	Example 3:
	Input: nums = [3,1,2,10,1]
	Output: [3,4,6,16,17]
*/
func TestRunningSum01(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expexted := []int{1, 3, 6, 10}
	output := solution.RunningSum(input)

	assert.Equal(t, expexted, output)
}

func TestRunningSum02(t *testing.T) {
	input := []int{1, 1, 1, 1, 1}
	expected := []int{1, 2, 3, 4, 5}
	output := solution.RunningSum(input)

	assert.Equal(t, expected, output)
}

func TestRunningSum03(t *testing.T) {
	input := []int{3, 1, 2, 10, 1}
	expexted := []int{3, 4, 6, 16, 17}
	output := solution.RunningSum(input)

	assert.Equal(t, expexted, output)
}
