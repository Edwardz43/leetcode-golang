package solution

/*
	Given the array of integers nums,
	you will choose two different indices i and j of that array.
	Return the maximum value of (nums[i]-1)*(nums[j]-1).

	Runtime: 4 ms, faster than 92.96% of Go online submissions for Maximum Product of Two Elements in an Array.
	Memory Usage: 3 MB, less than 87.60% of Go online submissions for Maximum Product of Two Elements in an Array.
*/

// MaxProduct return the maximum value
func MaxProduct(nums []int) int {
	max1 := nums[0]
	max2 := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > max1 {
			max2 = max1
			max1 = nums[i]
		} else if nums[i] > max2 {
			max2 = nums[i]
		}
	}

	return (max1 - 1) * (max2 - 1)
}

type ss struct {
	a string
	b int
	c bool
}

//
func newSs(a string, b int, c bool) *ss {
	return &ss{a: a, b: b, c: c}
}
