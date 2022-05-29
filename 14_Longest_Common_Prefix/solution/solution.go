package solution

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
Memory Usage: 2.3 MB, less than 50.53% of Go online submissions for Longest Common Prefix.
*/

// LongestCommonPrefix returns the longest common prefix string amongst an array of strings.
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var pre []byte
	for cursor := 0; ; cursor++ {
		for _, s := range strs[1:] {
			if cursor >= len(strs[0]) || cursor >= len(s) {
				return string(pre)
			}
			if s[cursor] != strs[0][cursor] {
				return string(pre)
			}
		}
		pre = append(pre, strs[0][cursor])
	}
	return string(pre)
}
