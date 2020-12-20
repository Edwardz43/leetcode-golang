package solution

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
Memory Usage: 2.3 MB, less than 12.92% of Go online submissions for Reverse Integer.
*/

// Reverse returns a reverse digits of an integer.
func Reverse(x int) int {

	res := make([]int, 0)
	tmp := x
	if x < 0 {
		tmp *= -1
	}
	for tmp > 0 {
		t := tmp % 10
		res = append(res, t)
		tmp = tmp / 10
	}
	r := 0
	for i := range res {
		r = r*10 + res[i]
	}
	if x < 0 {
		r *= -1
		if r < 1<<31*-1 {
			return 0
		}
		return r
	}
	if r > 1<<31-1 {
		return 0
	}
	return r
}
