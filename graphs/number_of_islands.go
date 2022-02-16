package graphs

// NumberOfIslands ...
//
// T -> O(nm)
// S -> O(nm)
func NumberOfIslands(grid [][]rune) int {
	if len(grid) == 0 {
		return 0
	}

	var count int
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] != 1 {
				continue
			}

			count++
			numIslandsDFS(i, j, &grid)
		}
	}

	return count
}

func numIslandsDFS(i, j int, grid *[][]rune) {
	if i < 0 || j < 0 || j >= len(*grid) || i >= len((*grid)[0]) || (*grid)[j][i] != 1 {
		return
	}

	(*grid)[j][i] = '#'

	numIslandsDFS(i+1, j, grid)
	numIslandsDFS(i, j+1, grid)
	numIslandsDFS(i-1, j, grid)
	numIslandsDFS(i, j-1, grid)
}
