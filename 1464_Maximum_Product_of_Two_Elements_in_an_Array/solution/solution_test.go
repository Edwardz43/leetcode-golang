package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/1464_Maximum_Product_of_Two_Elements_in_an_Array/solution"
	"github.com/bmizerany/assert"
)

/*
	Example 1:
	Input: nums = [3,4,5,2]
	Output: 12

	Example 2:
	Input: nums = [1,5,4,5]
	Output: 16

	Example 3:
	Input: nums = [3,7]
	Output: 12
*/
func TestMaxProduct01(t *testing.T) {
	input := []int{3, 4, 5, 2}
	expexted := 12
	output := solution.MaxProduct(input)
	assert.Equal(t, expexted, output)
}

func TestMaxProduct02(t *testing.T) {
	input := []int{1, 5, 4, 5}
	expexted := 16
	output := solution.MaxProduct(input)
	assert.Equal(t, expexted, output)
}

func TestMaxProduct03(t *testing.T) {
	input := []int{3, 7}
	expexted := 12
	output := solution.MaxProduct(input)
	assert.Equal(t, expexted, output)
}
