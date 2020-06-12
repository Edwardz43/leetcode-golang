package solution

/*
	Given two arrays of integers nums and index. Your task is to create target array under the following rules:
	Initially target array is empty.
	From left to right read nums[i] and index[i], insert at index index[i] the value nums[i] in target array.
	Repeat the previous step until there are no elements to read in nums and index.
	Return the target array.
	It is guaranteed that the insertion operations will be valid.

	Runtime: 0 ms, faster than 100.00% of Go online submissions for Create Target Array in the Given Order.
	Memory Usage: 2.3 MB, less than 30.66% of Go online submissions for Create Target Array in the Given Order.
*/

//CreateTargetArray returns the target array
func CreateTargetArray(nums []int, index []int) []int {
	res := []int{nums[0]}
	for i := 1; i < len(index); i++ {
		if index[i] > len(res)-1 {
			res = append(res, nums[i])
		} else {
			tmp := res[index[i]:]
			res = append(res[:index[i]], append([]int{nums[i]}, tmp...)...)
		}
	}
	return res
}
