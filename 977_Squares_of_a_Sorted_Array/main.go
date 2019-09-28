package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{-12, -8, -4, 0, 1, 8, 12, 121}
	output := sortedSquares2(input)
	fmt.Println(output)
}

/*
Runtime: 500 ms, faster than 52.57% of Go online submissions for Squares of a Sorted Array.
Memory Usage: 8.4 MB, less than 100.00% of Go online submissions for Squares of a Sorted Array.
*/
func sortedSquares(A []int) []int {
	B := []int{}
	for _, v := range A {
		B = append(B, v*v)
	}
	sort.Ints(B)
	return B
}

/*
Runtime: 476 ms, faster than 96.42% of Go online submissions for Squares of a Sorted Array.
Memory Usage: 8.3 MB, less than 100.00% of Go online submissions for Squares of a Sorted Array.
*/
func sortedSquares2(A []int) []int {
	size := len(A)
	res := make([]int, size)
	for l, r, k := 0, size-1, size-1; l <= r; k-- {
		if A[l]+A[r] < 0 {
			res[k] = A[l] * A[l]
			l++
		} else {
			res[k] = A[r] * A[r]
			r--
		}
	}
	return res
}
