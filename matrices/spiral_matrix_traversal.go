package matrices

// SpiralTraversal ...
//
// T -> O(n * m)
// S -> O(1)
func SpiralTraversal(matrix [][]int) []int {
	var (
		output                               = make([]int, 0, len(matrix)*len(matrix[0]))
		count, column, row, currentDirection int
		directions                           = [][]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		}
	)

	for count < len(matrix)*len(matrix[0]) {
		output = append(output, matrix[row][column])
		count++

		direction := directions[currentDirection]
		dj, di := direction[0], direction[1]

		switch {
		case row+dj < 0 || row+dj >= len(matrix) || (dj != 0 && matrix[row+dj][column] == 0):
			currentDirection = (currentDirection + 1) % 4
		case column+di < 0 || column+di >= len(matrix[0]) || (dj != 0 && matrix[column][row+di] == 0):
			currentDirection = (currentDirection + 1) % 4
		}

		direction = directions[currentDirection]
		dj, di = direction[0], direction[1]

		row += dj
		column += di
	}

	return output
}
