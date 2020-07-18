package solution

/**
Runtime: 0 ms, faster than 100.00% of Go online submissions for XOR Operation in an Array.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for XOR Operation in an Array.
*/

// XorOperation returns the bitwise XOR of all elements of nums
func XorOperation(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res ^= start + i*2
	}
	return res
}
