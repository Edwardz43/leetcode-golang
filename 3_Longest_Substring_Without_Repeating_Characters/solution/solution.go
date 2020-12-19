package solution

// LengthOfLongestSubstring find the length of the longest substring without repeating characters.
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	set := make(map[byte]bool, len(s))
	max := 0
	for i, j := 0, 1; i <= j && j < len(s); {
		// fmt.Printf("i=%d, j=%d, i: %s, j: %s\n", i, j, string(s[i]), string(s[j]))
		set[s[i]] = true
		cursor := s[j]
		if set[cursor] {
			if j-i > max {
				max = j - i
				// fmt.Printf("max=%d\n", max)
			}
			set = make(map[byte]bool, 0)
			i++
			j = i + 1
			continue
		}

		if j-i+1 > max {
			max = j - i + 1
			// fmt.Printf("max=%d\n", max)
		}
		// fmt.Printf("max=%d\n", max)
		set[cursor] = true
		j++
	}
	return max
}
