package main

import (
	"log"
)

/*
	Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it.
	That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].

	Return the answer in an array.

	Runtime: 4 ms, faster than 98.00% of Go online submissions for How Many Numbers Are Smaller Than the Current Number.
	Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for How Many Numbers Are Smaller Than the Current Number.
*/

func main() {

	// Input: nums = [8,1,2,2,3]
	// Output: [4,0,1,1,3]

	// Input: nums = [6,5,4,8]
	// Output: [2,1,0,3]

	// Input: nums = [7, 7, 7, 7]
	// Output: [0, 0, 0, 0]
	input := []int{7, 7, 7, 7}

	output := smallerNumbersThanCurrent(input)

	log.Println(output)
}

func smallerNumbersThanCurrent(nums []int) []int {

	result := make([]int, len(nums))
	for i, num := range nums {
		counter := 0
		for _, val := range nums {
			if val < num {
				counter = counter + 1
			}
		}
		result[i] = counter
	}
	return result
}
