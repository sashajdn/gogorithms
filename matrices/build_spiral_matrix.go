package matrices

// BuildSpiralMatrix ...
//
// T -> O(n ** 2)
// S -> O(1)
func BuildSpiralMatrix(n int) [][]int {
	var matrix = make([][]int, 0, n)
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}

	var (
		directions = [][]int{
			{0, 1},  // Right.
			{1, 0},  // Down.
			{0, -1}, // Left.
			{-1, 0}, // Up.
		}
		currentDirection int
		row, column      int
		count            = 1
	)

	for count < (n*n)+1 {
		matrix[row][column] = count
		count++

		dj, di := directions[currentDirection][0], directions[currentDirection][1]

		// Change direction if next step is out of bounds or we have previously visited.
		switch {
		case row+dj < 0 || row+dj >= len(matrix) || (dj != 0 && matrix[row+dj][column] != 0):
			currentDirection = (currentDirection + 1) % 4
		case column+di < 0 || column+di >= len(matrix[0]) || (di != 0 && matrix[row][column+di] != 0):
			currentDirection = (currentDirection + 1) % 4
		}

		dj, di = directions[currentDirection][0], directions[currentDirection][1]

		row += dj
		column += di
	}

	return matrix
}
