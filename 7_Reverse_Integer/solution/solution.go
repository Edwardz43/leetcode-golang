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

/*
Runtime: 4 ms, faster than 37.46% of Go online submissions for Reverse Integer.
Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Reverse Integer.
*/

// ReverseV2 uses less memory but slower than original one.
func ReverseV2(x int) int {
	res := [10]int{}
	negative := x < 0
	if negative {
		x *= -1
	}
	counter := 0
	for x > 0 {
		t := x % 10
		res[counter] = t
		x = x / 10
		counter++
	}
	r := 0
	for i := 0; i < counter; i++ {
		r = r*10 + res[i]
	}
	if negative {
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
