package solution

/*
Runtime: 4 ms, faster than 89.93% of Go online submissions for Roman to Integer.
Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Roman to Integer.
*/

// RomanToInt converts a roman numeral to an integer.
func RomanToInt(s string) int {
	// I [73]
	// V [86]
	// X [88]
	// L [76]
	// C [67]
	// D [68]
	// M [77]
	res := 0
	for b := range s {
		switch s[b] {
		case 73:
			res++
			break
		case 86:
			res += 5
			if b > 0 && s[b-1] == 73 {
				res -= 2
			}
			break
		case 88:
			res += 10
			if b > 0 && s[b-1] == 73 {
				res -= 2
			}
			break
		case 76:
			res += 50
			if b > 0 && s[b-1] == 88 {
				res -= 20
			}
			break
		case 67:
			res += 100
			if b > 0 && s[b-1] == 88 {
				res -= 20
			}
			break
		case 68:
			res += 500
			if b > 0 && s[b-1] == 67 {
				res -= 200
			}
			break
		case 77:
			res += 1000
			if b > 0 && s[b-1] == 67 {
				res -= 200
			}
			break
		}
	}
	return res
}
