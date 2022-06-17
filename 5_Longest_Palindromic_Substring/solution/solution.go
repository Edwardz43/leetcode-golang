package solution

/*
Runtime: 191 ms, faster than 24.76% of Go online submissions for Longest Palindromic Substring.
Memory Usage: 8.9 MB, less than 11.82% of Go online submissions for Longest Palindromic Substring.
*/
func longestPalindrome(s string) string {
	var (
		start, end int
	)
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for len := 0; len < n; len++ {
		for i := 0; i+len < n; i++ {
			dp[i][i+len] = s[i] == s[i+len] && (len < 2 || dp[i+1][i+len-1])
			if dp[i][i+len] && len > end-start {
				start = i
				end = i + len
			}
		}

	}
	return s[start : end+1]
}

/*
Runtime: 8 ms, faster than 78.57% of Go online submissions for Longest Palindromic Substring.
Memory Usage: 3.1 MB, less than 32.91% of Go online submissions for Longest Palindromic Substring.
*/
func longestPalindrome_v2(s string) string {
	if s == "" {
		return ""
	}

	runeStr := []rune(s)
	runeArrLength := len(runeStr)
	longest := []rune{rune(s[0])}

	for i := range runeStr {
		leftIdx, rightIdx := i, i
		for rightIdx+1 < runeArrLength && runeStr[leftIdx] == runeStr[rightIdx+1] {
			rightIdx++
		}

		for (leftIdx-1 >= 0) && (rightIdx+1 < runeArrLength) && runeStr[leftIdx-1] == runeStr[rightIdx+1] {
			leftIdx--
			rightIdx++
		}

		if rightIdx-leftIdx+1 > len(longest) {
			longest = runeStr[leftIdx : rightIdx+1]
		}
	}

	return string(longest)
}
