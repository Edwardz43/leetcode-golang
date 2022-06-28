package solution

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
