package graphs

type Point struct {
	x, y int
}

// ShortestPathBinaryMatrix ...
//
// T -> O(n) where n is the w*h of the grid.
// S -> O(n)
func ShortestPathBinaryMatrix(grid [][]int) int {
	var seen = make([][]bool, 0, len(grid))
	for i := 0; i < len(grid); i++ {
		seen = append(seen, make([]bool, len(grid[0])))
	}

	var (
		count   = 1
		minPath = -1
		queue   = []*Point{{0, 0}}
	)
	search(grid, seen, count, &minPath, &queue)
	return minPath
}

func search(grid [][]int, seen [][]bool, count int, minSumSoFar *int, queue *[]*Point) {
	if len(*queue) == 0 {
		return
	}

	l := len(*queue)
	for i := 0; i < l; i++ {
		nextPair := (*queue)[0]
		*queue = (*queue)[1:]

		if nextPair.y == len(grid)-1 && nextPair.x == len(grid[0])-1 && grid[nextPair.y][nextPair.x] == 0 {
			switch *minSumSoFar {
			case -1:
				*minSumSoFar = count
			default:
				*minSumSoFar = min(*minSumSoFar, count)
			}
		}

		if seen[nextPair.y][nextPair.x] {
			continue
		}
		seen[nextPair.y][nextPair.x] = true

		if grid[nextPair.y][nextPair.x] == 1 {
			continue
		}

		for _, neighbour := range getNeighbours(grid, nextPair.x, nextPair.y, seen) {
			x, y := neighbour[0], neighbour[1]
			if seen[y][x] {
				continue
			}

			*queue = append(*queue, &Point{x, y})
		}
	}

	search(grid, seen, count+1, minSumSoFar, queue)
}

func getNeighbours(grid [][]int, i, j int, seen [][]bool) [][]int {
	var nn = [][]int{
		{i - 1, j - 1}, {i, j}, {i + 1, j - 1},
		{i - 1, j}, {i + 1, j},
		{i - 1, j + 1}, {i, j + 1}, {i + 1, j + 1},
	}

	var validNeighbours [][]int
	for _, n := range nn {
		x, y := n[0], n[1]

		if x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid) && !seen[y][x] {
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
