package main

import "log"

/*
	Given the array candies and the integer extraCandies, where candies[i] represents the number of candies that the ith kid has.
	For each kid check if there is a way to distribute extraCandies among the kids such that he or she can have the greatest number of candies among them. Notice that multiple kids can have the greatest number of candies.

	Runtime: 0 ms, faster than 100.00% of Go online submissions for Kids With the Greatest Number of Candies.
	Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Kids With the Greatest Number of Candies.
*/
func main() {
	// candies := []int{2, 3, 5, 1, 3}
	// candies := []int{4, 2, 1, 1, 2}
	candies := []int{12, 1, 12}

	// extraCandies := 3
	// extraCandies := 1
	extraCandies := 10

	res := kidsWithCandies(candies, extraCandies)

	log.Println(res)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))

	max := 0
	for _, c := range candies {
		if c > max {
			max = c
		}
	}
	for i, c := range candies {
		if c+extraCandies >= max {
			res[i] = true
		} else {
			res[i] = false
		}
	}

	return res
}
