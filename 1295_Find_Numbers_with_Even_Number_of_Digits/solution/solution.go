package solution

import "strconv"

/*
	Given an array nums of integers, return how many of them contain an even number of digits.

	Example 1:
	Input: nums = [12,345,2,6,7896]
	Output: 2
	Explanation:
	12 contains 2 digits (even number of digits).
	7896 contains 4 digits (even number of digits).
	Therefore only 12 and 7896 contain an even number of digits.

	Example 2:
	Input: nums = [555,901,482,1771]
	Output: 1

	Runtime: 4 ms, faster than 90.22% of Go online submissions for Find Numbers with Even Number of Digits.
	Memory Usage: 3.2 MB, less than 100.00% of Go online submissions for Find Numbers with Even Number of Digits.
*/

// FindNumbers return how many of them contain an even number of digits
func FindNumbers(nums []int) int {
	counter := 0
	for _, num := range nums {
		s := strconv.Itoa(num)
		if len(s)%2 == 0 {
			counter++
		}
	}

	return counter
}
