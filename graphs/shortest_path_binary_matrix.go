package graphs

// ShortestPathBinaryMatrix ...
//
// T -> O(n) where n is the w*h of the grid.
// S -> O(n)
func ShortestPathBinaryMatrix(grid [][]int) int {
	var seen = make([][]bool, 0, len(grid))
	for i := 0; i < len(seen); i++ {
		seen = append(seen, make([]bool, len(seen[0])))
	}

	var (
		count   int
		minPath = -1
	)
	search(grid, seen, 0, 0, count, &minPath)
	return minPath
}

func search(grid [][]int, seen [][]bool, i, j, count int, minSumSoFar *int) {
	if j == len(grid)-1 && i == len(grid[0])-1 {
		count++
		*minSumSoFar = min(*minSumSoFar, count)
	}

	if seen[j][i] {
		return
	}
	seen[j][i] = true

	if grid[j][i] == 1 {
		return
	}
	if count >= *minSumSoFar {
		return
	}

	for _, neighbour := range getNeighbours(grid, i, j) {
		x, y := neighbour[0], neighbour[1]
		search(grid, seen, x, y, count+1, minSumSoFar)
	}
}

func getNeighbours(grid [][]int, i, j int) [][]int {
	var nn = [][]int{
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i, j - 1},
		{i, j + 1},
		{i + 1, j - 1},
		{i + 1, j},
		{i + 1, j + 1},
	}

	var validNeighbours [][]int
	for _, n := range nn {
		x, y := n[0], n[1]

		if x > 0 && x < len(grid[0])-1 && y > 0 && y < len(grid)-1 {
			validNeighbours = append(validNeighbours, n)
		}
	}
	return validNeighbours
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
