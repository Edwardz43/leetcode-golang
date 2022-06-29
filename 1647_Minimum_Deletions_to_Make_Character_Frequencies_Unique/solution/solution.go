package solution

import (
	"sort"
)

/*
Runtime: 37 ms, faster than 45.90% of Go online submissions for Minimum Deletions to Make Character Frequencies Unique.
Memory Usage: 6.9 MB, less than 68.85% of Go online submissions for Minimum Deletions to Make Character Frequencies Unique.
*/
func minDeletions(s string) int {
	f := make([]int, 26)
	for _, v := range s {
		f[v-'a']++
	}
	d := 0
	set := make(map[int]bool, 26)
	for i := 0; i < 26; i++ {
		for f[i] > 0 && set[f[i]] {
			f[i]--
			d++
		}
		set[f[i]] = true
	}
	return d
}

/*
Runtime: 12 ms, faster than 98.36% of Go online submissions for Minimum Deletions to Make Character Frequencies Unique.
Memory Usage: 6.4 MB, less than 98.36% of Go online submissions for Minimum Deletions to Make Character Frequencies Unique.
*/
func minDeletionsV2(s string) int {
	f := make([]int, 26)
	// var f [26]int
	for _, v := range s {
		f[v-'a']++
	}
	d := 0
	sort.Slice(f[:], func(i, j int) bool {
		return f[i] > f[j]
	})
	for i := 1; i < len(f); i++ {
		if f[i] == 0 {
			break
		}
		if f[i] < f[i-1] {
			continue
		}
		for f[i] > 0 && f[i] >= f[i-1] {
			f[i]--
			d++
		}
	}
	return d
}
