package main

import "fmt"

/*
Runtime: 48 ms, faster than 99.15% of Go online submissions for Sort Array By Parity.
Memory Usage: 8.1 MB, less than 100.00% of Go online submissions for Sort Array By Parity.
*/
func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := sortArrayByParity(input)
	fmt.Println(output)
}

func sortArrayByParity(A []int) []int {
	e := make([]int, 0)
	o := make([]int, 0)
	for _, v := range A {
		if v%2 == 0 {
			o = append(o, v)
		} else {
			e = append(e, v)
		}
	}
	return append(o, e...)
}
