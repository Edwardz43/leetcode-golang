package solution

/*
Runtime: 7 ms, faster than 78.43% of Go online submissions for Transpose Matrix.
Memory Usage: 6.3 MB, less than 84.31% of Go online submissions for Transpose Matrix.
*/

// Transpose return the transpose of matrix.
func Transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		res[i] = make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			res[i][j] = matrix[j][i]
		}
	}
	return res
}
