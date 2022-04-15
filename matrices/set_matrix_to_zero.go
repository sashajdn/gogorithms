package matrices

// SetMatrixToZero ...
//
// T -> O(w * h)
// S -> O(1)
func SetMatrixToZero(matrix [][]int) [][]int {
	var (
		setRow, setCol bool
	)

	for j := 0; j < len(matrix); j++ {
		if matrix[j][0] == 0 {
			setCol = true
			break
		}
	}

	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			setRow = true
			break
		}
	}

	for j := 1; j < len(matrix); j++ {
		for i := 1; i < len(matrix); i++ {
			if matrix[j][i] == 0 {
				matrix[j][0] = 0
				matrix[0][i] = 0
			}
		}
	}

	for j := 1; j < len(matrix); j++ {
		for i := 1; i < len(matrix); i++ {
			if matrix[j][0] == 0 || matrix[0][i] == 0 {
				matrix[j][i] = 0
			}
		}
	}

	if setCol {
		for j := 0; j < len(matrix); j++ {
			matrix[j][0] = 0
		}
	}
	if setRow {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	return matrix
}
