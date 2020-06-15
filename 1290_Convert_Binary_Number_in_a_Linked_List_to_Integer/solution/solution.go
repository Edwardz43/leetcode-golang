package solution

/*
	Given head which is a reference node to a singly-linked list.
	The value of each node in the linked list is either 0 or 1.
	The linked list holds the binary representation of a number.
	Return the decimal value of the number in the linked list.

	Runtime: 0 ms, faster than 100.00% of Go online submissions for Convert Binary Number in a Linked List to Integer.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Convert Binary Number in a Linked List to Integer.
*/

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Newnode create a new ListNode with specific int array
func Newnode(nums []int) *ListNode {
	res := &ListNode{Val: nums[0]}

	tmp := res
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{Val: nums[i]}
		tmp = tmp.Next
	}
	return res
}

// Len returns the length of ListNode
func (n *ListNode) Len() int {
	res := 0
	for ; n != nil; n = n.Next {
		res++
	}
	return res
}

// GetDecimalValue returns the decimal value of the number in the linked list
func GetDecimalValue(head *ListNode) int {
	res := 0
	for ; head != nil; head = head.Next {
		res = res<<1 | head.Val
	}
	return res
}
