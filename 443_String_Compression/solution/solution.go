package solution

import "strconv"

/*
	Given an array of characters, compress it in-place.
	The length after compression must always be smaller than or equal to the original array.
	Every element of the array should be a character (not int) of length 1.
	After you are done modifying the input array in-place, return the new length of the array.

	Example 1:
	Input: ["a","a","b","b","c","c","c"]
	Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
	Explanation: "aa" is replaced by "a2". "bb" is replaced by "b2". "ccc" is replaced by "c3".

	Example 2:
	Input: ["a"]
	Output: Return 1, and the first 1 characters of the input array should be: ["a"]
	Explanation: Nothing is replaced.

	Example 3:
	Input: ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
	Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
	Explanation:
	Since the character "a" does not repeat, it is not compressed. "bbbbbbbbbbbb" is replaced by "b12".
	Notice each digit has it's own entry in the array.

	Runtime: 4 ms, faster than 97.71% of Go online submissions for String Compression.
	Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for String Compression.
*/

// Compress return the new length of the array
func Compress(chars []byte) int {
	/*
		pseudo code:
		for c in chars
			initialize counter to 1
			assign c to current char and nextChar to the next char

			while c equal to nextChar
				counter adds 1
				if chars still has next one
					assign nextChar to the next one
					continue
				break

			if counter is greater than 1
				remove duplicates of c form chars
				convert counter to byte array bts
				for b in bts
					insert b to chars after c

			jump to next different c
	*/
	for i := 0; i < len(chars)-1; i++ {
		counter := 1
		c := chars[i]
		nextChar := chars[i+1]
		for c == nextChar {
			counter++
			if i+counter < len(chars) {
				nextChar = chars[i+counter]
				continue
			}
			break
		}
		if counter > 1 {
			chars = append(chars[:i+1], chars[i+counter:]...)
			bts := []byte(strconv.Itoa(counter))
			for _, b := range bts {
				chars = append(chars, 0)
				copy(chars[i+1:], chars[i:])
				chars[i+1] = b
				i++
			}
		}
	}
	return len(chars)
}
