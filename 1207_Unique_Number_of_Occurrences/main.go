package main

import "fmt"

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Number of Occurrences.
Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Unique Number of Occurrences.
*/

func main() {
	// input := []int{1, 2, 2, 1, 1, 3}
	// input := []int{1, 2}
	input := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	output := uniqueOccurrences(input)
	fmt.Println(output)
}

func uniqueOccurrences(arr []int) bool {
	check := make(map[int]int)
	for _, v := range arr {
		check[v]++
	}

	check2 := make(map[int]bool)
	for _, count := range check {
		if check2[count] == true {
			return false
		}
		check2[count] = true
	}
	return true
}
