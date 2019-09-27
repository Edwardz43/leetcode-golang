package main

import "fmt"

func main() {
	input := []int{1, -4, 2, 3, -12, 0, 121, -8}
	output := sortedSquares(input)

	fmt.Println(output)
}

func sortedSquares(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] *= A[i]
	}

	return quickSort(A, 0, len(A)-1)
}

func quickSort(A []int, front int, end int) []int {
	if front < end {
		pivot := partition(A, front, end)
		quickSort(A, front, pivot-1)
		quickSort(A, pivot+1, end)
	}
	return A
}

func partition(A []int, front int, end int) int {
	pivot := A[end]
	i := front - 1
	for j := front; j < end; j++ {
		if A[j] < pivot {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	i++
	A[i], A[end] = A[end], A[i]
	return i
}
