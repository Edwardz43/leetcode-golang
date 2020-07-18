package solution

/**
Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Good Pairs.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Number of Good Pairs.
*/

// NumIdenticalPairs returns the number of good pairs
func NumIdenticalPairs(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				res += 1
			}
		}
	}
	return res
}
