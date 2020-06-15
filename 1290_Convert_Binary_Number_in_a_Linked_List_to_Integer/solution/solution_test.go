package solution_test

import (
	"testing"

	"github.com/Edwardz43/leetcode-golang/1290_Convert_Binary_Number_in_a_Linked_List_to_Integer/solution"
	"github.com/bmizerany/assert"
)

/*
	Example 1:
	Input: head = [1,0,1]
	Output: 5

	Example 2:
	Input: head = [0]
	Output: 0

	Example 3:
	Input: head = [1]
	Output: 1

	Example 4:
	Input: head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
	Output: 18880

	Example 5:
	Input: head = [0,0]
	Output: 0
*/
func TestNewNode01(t *testing.T) {
	input := solution.Newnode([]int{1, 0, 1})
	assert.Equal(t, 1, input.Next.Next.Val)
	assert.Equal(t, 3, input.Len())
}

func TestNewNode02(t *testing.T) {
	input := solution.Newnode([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0})
	assert.Equal(t, 15, input.Len())
}

func TestGetDecimalValue01(t *testing.T) {
	input := solution.Newnode([]int{1, 0, 1})
	expexted := 5
	output := solution.GetDecimalValue(input)

	assert.Equal(t, expexted, output)
}

func TestGetDecimalValue02(t *testing.T) {
	input := solution.Newnode([]int{0})
	expexted := 0
	output := solution.GetDecimalValue(input)

	assert.Equal(t, expexted, output)
}

func TestGetDecimalValue03(t *testing.T) {
	input := solution.Newnode([]int{1})
	expexted := 1
	output := solution.GetDecimalValue(input)

	assert.Equal(t, expexted, output)
}

func TestGetDecimalValue04(t *testing.T) {
	input := solution.Newnode([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0})
	expexted := 18880
	output := solution.GetDecimalValue(input)

	assert.Equal(t, expexted, output)
}

func TestGetDecimalValue05(t *testing.T) {
	input := solution.Newnode([]int{0, 0})
	expexted := 0
	output := solution.GetDecimalValue(input)

	assert.Equal(t, expexted, output)
}
