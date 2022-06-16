package solution_test

import (
	"github.com/Edwardz43/leetcode-golang/2_Add_Two_Numbers/solution"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution01(t *testing.T) {
	l1 := generateNode([]int{2, 4, 3})
	l2 := generateNode([]int{5, 6, 4})
	expected := generateNode([]int{7, 0, 8})
	actual := solution.AddTwoNumbers(l1, l2)
	assert.Equal(t, expected, actual)
}

func generateNode(arr []int) *solution.ListNode {
	l := &solution.ListNode{}
	cursor := l
	for count := 0; count < len(arr); count++ {
		cursor.Val = arr[count]
		if count != len(arr)-1 {
			cursor.Next = &solution.ListNode{}
			cursor = cursor.Next
		}
	}
	return l
}
