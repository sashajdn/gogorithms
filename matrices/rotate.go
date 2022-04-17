package matrices

// RotateMatrix ...
//
// T -> O(n ** 2) where `n` is the length of the matrix.
// S -> O(1) as we rotate inplace iteratively.
func RotateMatrix(matrix [][]int) {
	for j := 0; j < len(matrix); j++ {
		for i := j + 1; i < len(matrix[0]); i++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}

	for j := 0; j < len(matrix); j++ {
		var left, right = 0, len(matrix) - 1
		for left < right {
			matrix[j][left], matrix[j][right] = matrix[j][right], matrix[j][left]
		}
	}
}
