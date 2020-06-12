package solution

/*
	Balanced strings are those who have equal quantity of 'L' and 'R' characters.
	Given a balanced string s split it in the maximum amount of balanced strings.
	Return the maximum amount of splitted balanced strings.

	Runtime: 0 ms, faster than 100.00% of Go online submissions for Split a String in Balanced Strings.
	Memory Usage: 2.1 MB, less than 11.45% of Go online submissions for Split a String in Balanced Strings.
*/

// BalancedStringSplit returns the maximum amount of splitted balanced strings
func BalancedStringSplit(s string) int {
	count := 0
	tmp := 0
	for _, v := range s {
		if v == 82 {
			tmp--
		}
		if v == 76 {
			tmp++
		}
		if tmp == 0 {
			count++
		}
	}
	return count
}
