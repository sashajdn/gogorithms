package arrays

// RotateMatrix ...
//
// T -> O(n ** 2)
// S -> O(1)
func RotateMatrix(matrix [][]int) {
	l := len(matrix) - 1

	for j := 0; j <= l/2; j++ {
		for i := j; i < l; i++ {
			swapMatrix(matrix, i, j, l, i)
			swapMatrix(matrix, i, j, l-i+j, l)
			swapMatrix(matrix, i, j, j, l-i+j)
		}

		l = l - 1
	}
}

func swapMatrix(matrix [][]int, i, j, m, n int) {
	matrix[j][i], matrix[n][m] = matrix[n][m], matrix[j][i]
}
