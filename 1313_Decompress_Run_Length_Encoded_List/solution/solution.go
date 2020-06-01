package solution

/*
	We are given a list nums of integers representing a list compressed with run-length encoding.
	Consider each adjacent pair of elements [freq, val] = [nums[2*i], nums[2*i+1]] (with i >= 0).
	For each such pair, there are freq elements with value val concatenated in a sublist.
	Concatenate all the sublists from left to right to generate the decompressed list.

	Example :
	Input: nums = [1,2,3,4]
	Output: [2,4,4,4]

	Input: nums = [1,1,2,3]
	Output: [1,3,3]

	Runtime: 4 ms, faster than 92.78% of Go online submissions for Decompress Run-Length Encoded List.
	Memory Usage: 5.9 MB, less than 100.00% of Go online submissions for Decompress Run-Length Encoded List.
*/

//DecompressRLElist return the decompressed list.
func DecompressRLElist(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}
	return res
}
