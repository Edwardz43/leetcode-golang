package solution

/*
Runtime: 4 ms, faster than 88.05% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 2.7 MB, less than 76.99% of Go online submissions for Longest Substring Without Repeating Characters.
*/

// LengthOfLongestSubstring find the length of the longest substring without repeating characters.
func LengthOfLongestSubstring(s string) int {
	cache := [128]bool{}
	max, lengh, cursor := 0, 0, 0
	for i := range s {
		for cache[s[i]] {
			cache[s[cursor]] = false
			lengh--
			cursor++
		}
		lengh++
		cache[s[i]] = true
		if lengh > max {
			max = lengh
		}
	}
	return max
}
