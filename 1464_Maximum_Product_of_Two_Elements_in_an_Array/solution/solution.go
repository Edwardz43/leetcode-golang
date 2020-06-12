package solution

import "sort"

/*
	Given the array of integers nums,
	you will choose two different indices i and j of that array.
	Return the maximum value of (nums[i]-1)*(nums[j]-1).

	Runtime: 4 ms, faster than 92.96% of Go online submissions for Maximum Product of Two Elements in an Array.
	Memory Usage: 3 MB, less than 60.47% of Go online submissions for Maximum Product of Two Elements in an Array.
*/

// MaxProduct return the maximum value
func MaxProduct(nums []int) int {
	var d data = nums
	sort.Sort(d)

	return (d[0] - 1) * (d[1] - 1)
}

type data []int

func (d data) Len() int { return len(d) }

func (d data) Less(i, j int) bool { return d[i] > d[j] }

func (d data) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
