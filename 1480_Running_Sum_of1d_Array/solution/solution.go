package solution

/*
	Given an array nums.
	We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
	Return the running sum of nums.

	Runtime: 0 ms, faster than 100.00% of Go online submissions for Running Sum of 1d Array.
	Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Running Sum of 1d Array.
*/

// RunningSum returns the running sum of nums
func RunningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
