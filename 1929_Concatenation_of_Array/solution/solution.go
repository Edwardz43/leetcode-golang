package solution

/*
Runtime: 15 ms, faster than 31.24% of Go online submissions for Concatenation of Array.
Memory Usage: 6.5 MB, less than 10.06% of Go online submissions for Concatenation of Array.
*/
func getConcatenation(nums []int) []int {
	res := make([]int, len(nums)*2)
	for i, num := range nums {
		res[i] = num
		res[i+len(nums)] = num
	}
	return res
}

/*
Runtime: 8 ms, faster than 86.79% of Go online submissions for Concatenation of Array.
Memory Usage: 6.2 MB, less than 73.17% of Go online submissions for Concatenation of Array.
*/
func getConcatenation_v2(nums []int) []int {
	res := make([]int, len(nums)*2, 2*cap(nums))
	for i, num := range nums {
		res[i] = num
		res[i+len(nums)] = num
	}
	return res
}
