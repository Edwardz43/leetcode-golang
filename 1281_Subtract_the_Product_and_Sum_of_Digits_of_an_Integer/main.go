package main

import "log"

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Subtract the Product and Sum of Digits of an Integer.
Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Subtract the Product and Sum of Digits of an Integer.
*/

func main() {
	input := 4421
	result := subtractProductAndSum(input)
	log.Println(result)
}

func subtractProductAndSum(n int) int {
	digitSum := 0
	productDigits := 1
	for n != 0 {
		digitSum += n % 10
		productDigits *= n % 10
		n /= 10
	}
	return productDigits - digitSum
}
