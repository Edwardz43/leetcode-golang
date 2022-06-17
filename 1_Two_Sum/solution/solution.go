package solution

/*
Runtime: 85 ms, faster than 5.23% of Go online submissions for Two Sum.
Memory Usage: 3.7 MB, less than 71.67% of Go online submissions for Two Sum.
time complexity: O(N^2)
*/
func twoSum(nums []int, target int) []int {
	var (
		l = 0
		r = 1
	)
	for l < r && r < len(nums) {
		if nums[l]+nums[r] == target {
			break
		}
		r++
		if r >= len(nums) {
			l++
			r = l + 1
		}
	}
	return []int{l, r}
}

/*
Runtime: 7 ms, faster than 80.21% of Go online submissions for Two Sum.
Memory Usage: 4.1 MB, less than 66.44% of Go online submissions for Two Sum.
time complexity: O(N)
*/
func twoSum_v2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		index, ok := m[target-num]
		if ok && index != i {
			return []int{index, i}
		}
		m[num] = i
	}
	return []int{}
}
