package main

import "log"

/*
Runtime: 28 ms, faster than 90.00% of Go online submissions for N-Repeated Element in Size 2N Array.
Memory Usage: 6.2 MB, less than 100.00% of Go online submissions for N-Repeated Element in Size 2N Array.
*/

func main() {
	input := []int{5, 1, 5, 2, 5, 3, 5, 4}
	output := repeatedNTimes(input)
	log.Println(output)
}

func repeatedNTimes(A []int) int {
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] == A[j] {
				return A[i]
			}
		}
	}

	return 0
}
