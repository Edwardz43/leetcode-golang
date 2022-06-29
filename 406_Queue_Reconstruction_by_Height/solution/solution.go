package solution

import "sort"

/*
Runtime: 17 ms, faster than 80.70% of Go online submissions for Queue Reconstruction by Height.
Memory Usage: 6.6 MB, less than 63.16% of Go online submissions for Queue Reconstruction by Height.
*/
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})
	indexes := make([]int, len(people))
	for i := range indexes {
		indexes[i] = i
	}
	queued := make([][]int, len(people))
	for _, v := range people {
		queued[indexes[v[1]]] = v
		indexes = append(indexes[:v[1]], indexes[v[1]+1:]...)
	}
	return queued
}

/*
Runtime: 10 ms, faster than 100% of Go online submissions for Queue Reconstruction by Height.
Memory Usage: 6.6 MB, less than 63.16% of Go online submissions for Queue Reconstruction by Height.
*/
func reconstructQueueV2(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	q := make([][]int, len(people))
	for _, p := range people {
		k := p[1]
		copy(q[k+1:], q[k:])
		q[k] = p
	}
	return q
}

/*
Runtime: 27 ms, faster than 49.12% of Go online submissions for Queue Reconstruction by Height.
Memory Usage: 6.2 MB, less than 98.25% of Go online submissions for Queue Reconstruction by Height.
*/
func reconstructQueueV3(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	n := len(people)
	for i := 0; i < n; i++ {
		hi := people[i][1]

		for j := i; j > hi; j-- {
			people[j], people[j-1] = people[j-1], people[j]
		}
	}
	return people
}
