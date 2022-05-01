package graphs

// RottingOranges ...
//
// T -> O(n ** 2)
// S -> O(1)
func RottingOranges(grid [][]int) int {
	var (
		currentTerm = 2
		directions  = [][]int{
			{0, 1},
			{0, -1},
			{1, 0},
			{-1, 0},
		}
	)

	var bfs func() bool
	bfs = func() bool {
		var canContinue bool
		for j := 0; j < len(grid); j++ {
			for i := 0; i < len(grid); i++ {
				if grid[j][i] != currentTerm {
					continue
				}

				for _, direction := range directions {
					dj, di := j+direction[0], i+direction[1]

					if dj < 0 || dj > len(grid)-1 {
						continue
					}

					if di < 0 || di > len(grid[0])-1 {
						continue
					}

					if grid[dj][di] != 1 {
						continue
					}

					grid[dj][di] = currentTerm + 1
					canContinue = true
				}
			}
		}
		return canContinue
	}

	for bfs() {
		currentTerm++
	}

	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] == 1 {
				return -1
			}
		}
	}

	return currentTerm - 2
}
