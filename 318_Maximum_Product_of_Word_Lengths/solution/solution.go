package solution

// MaxProduct is the original version
func MaxProduct(words []string) int {
	max := 0
	for i := 0; i < len(words)-1; i++ {
		a := words[i]
		for j := i + 1; j < len(words); j++ {
			b := words[j]
			sum := uint8(1)
		loop:
			for i, _ := range a {
				for k, _ := range b {
					sum = a[i] ^ b[k]
					if a[i]^b[k] == 0 {
						break loop
					}
				}
			}
			if sum != 0 {
				if len(a)*len(b) > max {
					max = len(a) * len(b)
				}
			}
		}
	}
	return max
}

// MaxProduct2 is the optimal version
func MaxProduct2(words []string) int {
	best, buf := 0, make([]int, len(words))
	for i, x := range words {
		mask, n := 0, len(x)
		for _, y := range []byte(x) {
			mask |= 1 << (y - 'a')
		}
		buf[i] = mask
		for j, m := range buf[:i] {
			if m&mask != 0 {
				continue
			}
			if now := len(words[j]) * n; best < now {
				best = now
			}
		}
	}
	return best
}
