package main

import (
	"fmt"
)

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Self Dividing Numbers.
Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Self Dividing Numbers.
*/

func main() {
	input := []int{1, 22}
	output := selfDividingNumbers(input[0], input[1])
	fmt.Println(output)
}

func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)
	for i := left; i <= right; i++ {
		c := i
		check := true
		for c != 0 {
			d := c % 10
			if d == 0 || i%d != 0 {
				check = false
				break
			}
			c /= 10
		}
		if check {
			result = append(result, i)
		}

	}

	return result
}
