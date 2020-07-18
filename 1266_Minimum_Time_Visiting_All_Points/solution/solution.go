package solution

/**
Runtime: 0 ms, faster than 100.00% of Go online submissions for Minimum Time Visiting All Points.
Memory Usage: 3.5 MB, less than 10.00% of Go online submissions for Minimum Time Visiting All Points.
*/

// MinTimeToVisitAllPoints returns the minimum time in seconds to visit all points
func MinTimeToVisitAllPoints(points [][]int) int {
	res := 0
	start := points[0]
	points = append(points[1:])
	for _, tuple := range points {
		x := max(start[0], tuple[0]) - min(start[0], tuple[0])
		y := max(start[1], tuple[1]) - min(start[1], tuple[1])
		res += max(x, y)
		start = tuple
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
