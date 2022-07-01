package solution

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	sum := 0
	for i := 0; i < len(boxTypes); i++ {

		if truckSize >= boxTypes[i][0] {
			sum += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		} else {
			sum += truckSize * boxTypes[i][1]
			break
		}
	}
	return sum
}
