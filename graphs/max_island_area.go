package graphs

// MaxIslandArea ...
//
// T -> O(n * m)
// S -> O(n * m)
func MaxIslandArea(grid [][]int) int {
	var directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var graph = map[int][]int{}
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] == 0 {
				continue
			}

			key := gridHash(len(grid[0]), j, i)

			graph[key] = []int{}
			for _, direction := range directions {
				dj, di := j+direction[0], i+direction[1]

				if dj < 0 || dj > len(grid)-1 {
					continue
				}
				if di < 0 || di > len(grid[0])-1 {
					continue
				}
				if grid[dj][di] == 0 {
					continue
				}

				graph[key] = append(graph[key], gridHash(len(grid[0]), dj, di))
			}
		}
	}

	if len(graph) == 0 {
		return 0
	}

	var dsu = NewMaxIslandDSU(len(grid) * len(grid[0]))
	for node, edges := range graph {
		for _, edge := range edges {
			dsu.Union(node, edge)
		}
	}

	return dsu.MaxRank()
}

type MaxIslandDSU struct {
	reps, ranks []int
	maxRank     int
}

func NewMaxIslandDSU(size int) *MaxIslandDSU {
	var reps, ranks = make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		reps[i] = i
		ranks[i] = 1
	}

	return &MaxIslandDSU{
		reps:    reps,
		ranks:   ranks,
		maxRank: 1,
	}
}

func (d *MaxIslandDSU) MaxRank() int {
	return d.maxRank
}

func (d *MaxIslandDSU) Find(group int) int {
	if d.reps[group] == group {
		return group
	}

	// Path Compression.
	d.reps[group] = d.Find(d.reps[group])
	return d.reps[group]
}

func (d *MaxIslandDSU) Union(a, b int) {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return
	}

	switch {
	case d.ranks[ra] >= d.ranks[rb]:
		d.reps[rb] = d.reps[ra]
		d.ranks[ra] += d.ranks[rb]
	default:
		d.reps[ra] = d.reps[rb]
		d.ranks[rb] += d.ranks[ra]
	}

	d.maxRank = max(d.maxRank, max(d.ranks[ra], d.ranks[rb]))

	return
}

func gridHash(size, j, i int) int {
	return j*size + i
}
