package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/1389_Create_Target_Array_in_the_Given_Order/solution"
	"github.com/bmizerany/assert"
)

/*
	Example 1:
	Input: nums = [0,1,2,3,4], index = [0,1,2,2,1]
	Output: [0,4,1,3,2]

	Example 2:
	Input: nums = [1,2,3,4,0], index = [0,1,2,3,0]
	Output: [0,1,2,3,4]

	Example 3:
	Input: nums = [1], index = [0]
	Output: [1]
*/

func TestCreateTargetArray01(t *testing.T) {
	input := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	expexted := []int{0, 4, 1, 3, 2}
	output := solution.CreateTargetArray(input, index)

	assert.Equal(t, expexted, output)
}

func TestCreateTargetArray02(t *testing.T) {
	input := []int{1, 2, 3, 4, 0}
	index := []int{0, 1, 2, 3, 0}
	expexted := []int{0, 1, 2, 3, 4}
	output := solution.CreateTargetArray(input, index)

	assert.Equal(t, expexted, output)
}

func TestCreateTargetArray03(t *testing.T) {
	input := []int{1}
	index := []int{0}
	expexted := []int{1}
	output := solution.CreateTargetArray(input, index)

	assert.Equal(t, expexted, output)
}
