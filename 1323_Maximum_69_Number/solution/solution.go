package solution

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Maximum 69 Number.
Memory Usage: 1.9 MB, less than 16.67% of Go online submissions for Maximum 69 Number.
*/

// Maximum69Number returns the maximum number by changing at most one digit (6 becomes 9, and 9 becomes 6).
func Maximum69Number(num int) int {
	if num/1000 == 6 {
		return num + 3000
	}
	if num%1000/100 == 6 {
		return num + 300
	}
	if num%100/10 == 6 {
		return num + 30
	}
	if num%10 == 6 {
		return num + 3
	}
	return num
}
