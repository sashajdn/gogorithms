package graphs

// MakeALargeIsland ...
//
// T -> O(n * m) where `m` is the height of the grid, and `n` is the width.
// S -> O(n * m)
func MakeALargeIsland(grid [][]int) int {
	// Use DSU data structure to find connected components with size.
	dsu := NewGridDSU(len(grid) * len(grid[0]))

	// Build DSU. Connect current node to it's neighbours.
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] == 0 {
				continue
			}

			currentGroup := (j * len(grid[0])) + i
			for _, neighbour := range fetchGridNeighbours(j, i, grid) {
				dj, di := neighbour[0], neighbour[1]
				if grid[dj][di] == 0 {
					continue
				}

				neighbourGroup := (dj * len(grid[0])) + di

				dsu.Union(currentGroup, neighbourGroup)
			}
		}
	}

	// Find largest component if we flip a `0` to a `1`.
	var maxSoFar = dsu.MaxRank()
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] == 1 {
				continue
			}

			// Fetch all neighbour components, using a set to remove duplicates.
			// Add to current size.
			var (
				currentIslandSize = 1
				set               = map[int]struct{}{}
			)
			for _, neighbour := range fetchGridNeighbours(j, i, grid) {
				dj, di := neighbour[0], neighbour[1]
				if grid[dj][di] == 0 {
					continue
				}

				neighbourID := (dj * len(grid[0])) + di

				representative := dsu.Find(neighbourID)
				if _, ok := set[representative]; ok {
					continue
				}
				set[representative] = struct{}{}

				currentIslandSize += dsu.RankByGroup(neighbourID)
			}

			maxSoFar = max(maxSoFar, currentIslandSize)
		}
	}

	return maxSoFar
}

func fetchGridNeighbours(j, i int, grid [][]int) [][]int {
	var (
		neighbours = make([][]int, 0, 4)
		directions = [][]int{
			{0, 1},
			{0, -1},
			{1, 0},
			{-1, 0},
		}
	)
	for _, direction := range directions {
		dj, di := j+direction[0], i+direction[1]

		if dj < 0 || dj > len(grid)-1 {
			continue
		}

		if di < 0 || di > len(grid[0])-1 {
			continue
		}

		neighbours = append(neighbours, []int{dj, di})
	}

	return neighbours
}

type GridDSU struct {
	reps, ranks []int
	maxRank     int
}

func NewGridDSU(size int) *GridDSU {
	var (
		reps  = make([]int, size)
		ranks = make([]int, size)
	)
	for i := 0; i < size; i++ {
		reps[i] = i
		ranks[i] = 1
	}

	return &GridDSU{
		reps:    reps,
		ranks:   ranks,
		maxRank: 1,
	}
}

func (d *GridDSU) MaxRank() int {
	return d.maxRank
}

func (d *GridDSU) Find(group int) int {
	if d.reps[group] == group {
		return group
	}

	d.reps[group] = d.Find(d.reps[group])
	return d.reps[group]
}

func (d *GridDSU) Union(a, b int) {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return
	}

	if d.ranks[ra] >= d.ranks[rb] {
		d.reps[rb] = d.reps[ra]
		d.ranks[ra] += d.ranks[rb]
		d.maxRank = max(d.maxRank, d.ranks[ra])
		return
	}

	d.reps[ra] = d.reps[rb]
	d.ranks[rb] += d.ranks[ra]
	d.maxRank = max(d.maxRank, d.ranks[rb])
}

func (d *GridDSU) RankByGroup(group int) int {
	return d.ranks[d.Find(group)]
}
