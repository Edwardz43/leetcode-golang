package solution

import (
	"sort"
	"strings"
)

/*
Runtime: 51 ms, faster than 30.23% of Go online submissions for Longest String Chain.
Memory Usage: 7.5 MB, less than 8.72% of Go online submissions for Longest String Chain.
*/

func longestStrChain(words []string) int {
	if len(words) == 1 {
		return 1
	}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	m := make(map[string]int)
	l := 1
	for i := 0; i < len(words); i++ {
		b := words[i]
		if m[b] == 0 {
			m[b] = 1
		}
		for j := 0; j < len(b); j++ {
			builder := strings.Builder{}
			builder.WriteString(b[0:j])
			builder.WriteString(b[j+1:])
			if m[builder.String()] != 0 {
				m[b] = m[builder.String()] + 1
			}
			if m[b] > l {
				l = m[b]
			}
		}
	}
	return l
}

func checkIfPredecessor(a, b string) bool {
	if len(a)+1 != len(b) {
		return false
	}

	var i, j int

	var inserted bool
	for i < len(a) && j < len(b) {
		if a[i] != b[j] && inserted {
			return false
		}
		if a[i] != b[j] {
			j++
			inserted = true
			continue
		}
		i++
		j++
	}
	return true
}

/*
Runtime: 36 ms, faster than 57.06% of Go online submissions for Longest String Chain.
Memory Usage: 4.9 MB, less than 97.65% of Go online submissions for Longest String Chain.
*/
func longestStrChain_v2(words []string) int {

	dp := make([]int, len(words))

	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })

	var max int
	for i := range words {
		for j := 0; j < i; j++ {

			if checkIfPredecessor(words[j], words[i]) {
				cnt := 1 + dp[j]

				if cnt > dp[i] {
					dp[i] = cnt
				}
				if cnt > max {
					max = cnt
				}
			}
			// fmt.Printf("words[%d]: %s, words[%d]: %s, dp: %d, max: %d\n", i, words[i], j, words[j], dp[i], max)
		}
	}

	return max + 1
}

func isPred(p, s string) (result bool) {

	if len(s)-len(p) != 1 {
		return false
	}

	i := 0
	j := 0
	inserted := false

	for i < len(p) && j < len(s) {
		if p[i] != s[j] {
			if !inserted {
				inserted = true
				j++
				continue
			} else {
				return false
			}
		}

		i++
		j++
	}

	if i == len(p) && j == len(s) && inserted {
		return true
	} else if i == len(p) && j+1 == len(s) && !inserted {
		return true
	}

	return false
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func longestStrChain_v3(words []string) int {
	maxLen := 0
	wordsByLen := make(map[int][]string, 0)

	for _, w := range words {
		maxLen = max(maxLen, len(w))
		wordsByLen[len(w)] = append(wordsByLen[len(w)], w)
	}

	result := 0
	cache := make(map[string]int, len(words))

	for i := maxLen; i > 0; i-- {
		predWords := wordsByLen[i]
		succWords := wordsByLen[i+1]

		for _, p := range predWords {
			cache[p] = 1
			for _, s := range succWords {
				if isPred(p, s) {
					cache[p] = max(cache[p], cache[s]+1)
				}
			}

			result = max(result, cache[p])
		}
	}

	return result
}
