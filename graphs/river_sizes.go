package graphs

// RiverSizes accepts a matrix of 0s and 1s; where 0s represent land & 1s represent water.
// It returns an array of integers that describe the length of all rivers in no particular ordrer.
//
// O(T) -> O(xy), O(S) -> O(xy); where y := height of the matrix, and x := the width of the matrix
func RiverSizes(matrix [][]int) []int {
	sizes := []int{}
	if len(matrix) == 0 {
		return sizes
	}

	seen := make([][]bool, 0, len(matrix))
	for i := 0; i < len(matrix); i++ {
		seen = append(seen, make([]bool, len(matrix[0])))
	}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if seen[y][x] {
				continue
			}

			if river := dps(matrix, seen, x, y); river > 0 {
				sizes = append(sizes, river)
			}
		}
	}

	return sizes
}

func dps(matrix [][]int, seen [][]bool, x, y int) int {
	if seen[y][x] {
		return 0
	}
	seen[y][x] = true

	if matrix[y][x] == 0 {
		return 0
	}

	// Includes the current node; which has now been marked as seen.
	var sizeSoFar = 1
	neighbours := fetchNeighbours(matrix, seen, x, y)
	for _, n := range neighbours {
		branchLength := dps(matrix, seen, n[0], n[1])
		sizeSoFar += branchLength
	}
	return sizeSoFar
}

func fetchNeighbours(matrix [][]int, seen [][]bool, x, y int) [][]int {
	neighbours := [][]int{}

	if y > 0 && !seen[y-1][x] {
		neighbours = append(neighbours, []int{x, y - 1})
	}
	if x > 0 && !seen[y][x-1] {
		neighbours = append(neighbours, []int{x - 1, y})
	}
	if y < len(matrix)-1 && !seen[y+1][x] {
		neighbours = append(neighbours, []int{x, y + 1})
	}
	if x < len(matrix[0])-1 && !seen[y][x+1] {
		neighbours = append(neighbours, []int{x + 1, y})
	}

	return neighbours
}
