package main

import "log"

/*
Runtime: 4 ms, faster than 94.08% of Go online submissions for Flipping an Image.
Memory Usage: 3.7 MB, less than 100.00% of Go online submissions for Flipping an Image.
*/
func main() {
	input := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	output := flipAndInvertImage(input)
	log.Println(output)
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, p := range A {
		for i := 0; i < len(p)/2; i++ {
			p[i], p[len(p)-i-1] = p[len(p)-i-1]^1, p[i]^1
		}
		if len(p)%2 != 0 {
			p[len(p)/2] ^= 1
		}
	}

	return A
}
