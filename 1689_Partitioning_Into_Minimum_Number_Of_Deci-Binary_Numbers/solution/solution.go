package solution

/*
Runtime: 14 ms, faster than 67.69% of Go online submissions for Partitioning Into Minimum Number Of Deci-Binary Numbers.
Memory Usage: 6.5 MB, less than 60.00% of Go online submissions for Partitioning Into Minimum Number Of Deci-Binary Numbers.
*/

func minPartitions(n string) int {
	var max rune
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	return int(max) - 48
}

func minPartitions_v2(n string) int {
	r := 0
	for i := range n {
		t := int(n[i]) - int('0')
		if t > r {
			r = t
		}
	}
	return r
}
