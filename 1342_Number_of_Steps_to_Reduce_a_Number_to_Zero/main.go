package main

import "log"

/*
	Given a non-negative integer num, return the number of steps to reduce it to zero.
	If the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.

	Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Steps to Reduce a Number to Zero.
	Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Number of Steps to Reduce a Number to Zero.
*/
func main() {
	input := 123
	res := numberOfSteps(input)
	log.Println(res)
}

func numberOfSteps(num int) int {
	res := 0
	reduce(num, &res)
	return res
}

func reduce(num int, count *int) {
	if num == 0 {
		return
	}
	if num&1 == 1 {
		num--
	} else {
		num = num >> 1
	}
	*count++
	reduce(num, count)

}
