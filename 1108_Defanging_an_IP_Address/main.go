package main

import (
	"log"
)

/*
Runtime: 0 ms
Memory Usage: 2 MB
*/
func main() {
	s := defangIPaddr("1.2.3.4")
	log.Println(s)

}

func defangIPaddr(address string) string {
	result := []byte{}
	for _, v := range address {
		if v == 46 {
			result = append(result, []byte("[.]")...)
		} else {
			result = append(result, byte(v))
		}
	}
	return string(result)
}
