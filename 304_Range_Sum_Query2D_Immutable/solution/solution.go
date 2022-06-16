package solution

/*
Runtime: 616 ms, faster than 100.00% of Go online submissions for Range Sum Query 2D - Immutable.
Memory Usage: 15.3 MB, less than 84.92% of Go online submissions for Range Sum Query 2D - Immutable.
*/

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := len(matrix) + 1
	m := len(matrix[0]) + 1
	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]int, m)
	}
	for i, mv := range matrix {
		for j, v := range mv {
			sum[i+1][j+1] =
				v + sum[i+1][j] + sum[i][j+1] - sum[i][j]
		}
	}
	return NumMatrix{
		matrix: sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 < 0 || col1 < 0 {
		return 0
	}
	if row2 >= len(this.matrix) || col2 >= len(this.matrix[0]) {
		return 0
	}
	return this.matrix[row2+1][col2+1] -
		this.matrix[row1][col2+1] - this.matrix[row2+1][col1] + this.matrix[row1][col1]
}
