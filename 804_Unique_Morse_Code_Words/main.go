package main

import "log"

/*
Runtime: 0 ms
Memory Usage: 2.4 MB
*/
func main() {
	input := []string{"gin", "zen", "gig", "msg"}
	result := uniqueMorseRepresentations(input)
	log.Println(result)
}

var morsecode = map[rune]string{
	'a': ".-", 'b': "-...", 'c': "-.-.", 'd': "-..", 'e': ".", 'f': "..-.", 'g': "--.", 'h': "....", 'i': "..", 'j': ".---",
	'k': "-.-", 'l': ".-..", 'm': "--", 'n': "-.", 'o': "---", 'p': ".--.", 'q': "--.-", 'r': ".-.", 's': "...", 't': "-",
	'u': "..-", 'v': "...-", 'w': ".--", 'x': "-..-", 'y': "-.--", 'z': "--..",
}

func uniqueMorseRepresentations(words []string) int {

	check := make(map[string]byte, 0)

	for _, word := range words {
		s := []byte{}
		for _, b := range word {
			v, _ := morsecode[b]
			s = append(s, v...)
		}
		check[string(s)] = 0
	}

	return len(check)
}
