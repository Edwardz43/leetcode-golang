package solution

/*
Runtime: 20 ms, faster than 61.38% of Go online submissions for Missing Number.
Memory Usage: 6.2 MB, less than 97.21% of Go online submissions for Missing Number.
*/

// MissingNumber returns the missing number in the given array.
func MissingNumber(nums []int) int {
	sum := (1 + len(nums)) * len(nums) >> 1
	for _, v := range nums {
		sum -= v
	}
	return sum
}
func MissingNumber2(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return (1+len(nums))*len(nums)>>1 - s
}
