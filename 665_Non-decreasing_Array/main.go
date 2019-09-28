package main

import (
	"fmt"
)

/*
Runtime: 24 ms, faster than 85.83% of Go online submissions for Non-decreasing Array.
Memory Usage: 6.2 MB, less than 100.00% of Go online submissions for Non-decreasing Array.
*/

func main() {
	input := []int{-1, 4, 2, 3}
	output := checkPossibility(input)

	fmt.Println(output)
}

func checkPossibility(nums []int) bool {
	breakpoint := false // mark has break or not
	for i, v := range nums {
		if i == 0 || nums[i-1] <= v {
			continue
		}

		if breakpoint { // found second break, return false
			return false
		}

		breakpoint = true               // found break, mark
		if i == 1 || i == len(nums)-1 { // 1 and len(nums)-1 is OK
			continue
		} else if nums[i-2] <= v || nums[i-1] <= nums[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}
