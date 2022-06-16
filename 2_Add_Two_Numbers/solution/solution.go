package solution

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers returns the sum as a linked list.
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	carry := false
	cursor := res
	for l1.Next != nil || l2.Next != nil {
		cursor.Val = l1.Val + l2.Val
		if carry {
			cursor.Val += 1
		}
		if cursor.Val > 9 {
			cursor.Val = 0
			carry = true
		}
		if l1.Next != nil || l2.Next != nil {
			cursor.Next = &ListNode{}
			cursor = cursor.Next
		}
		if l1.Next != nil {
			l1 = l1.Next
		}
		if l2.Next != nil {
			l2 = l2.Next
		}
	}
	return res
}
