package graphs

// NumberOfIslandsII ...
//
// T -> O(k * log(m * n))
// S -> O(max(k, m * n))
func NumberOfIslandsII(m, n int, positions [][]int) []int {
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var (
		land            = map[int]struct{}{}
		dsu             = NewIslandDSU(m * n)
		islands         = make([]int, len(positions))
		numberOfIslands int
	)
	for k, position := range positions {
		j, i := position[0], position[1]

		key := reduceDimension(j, i, n)
		if _, ok := land[key]; ok {
			islands[k] = numberOfIslands
			continue
		}
		numberOfIslands++

		for _, direction := range directions {
			dj, di := j+direction[0], i+direction[1]

			if dj < 0 || dj > m-1 {
				continue
			}
			if di < 0 || di > n-1 {
				continue
			}

			neighbouringKey := reduceDimension(dj, di, n)
			if _, ok := land[neighbouringKey]; !ok {
				continue
			}

			if dsu.Union(key, neighbouringKey) {
				numberOfIslands--
			}
		}

		islands[k] = numberOfIslands
		land[key] = struct{}{}
	}

	return islands
}

type IslandDSU struct {
	reps, ranks []int
}

func NewIslandDSU(size int) *IslandDSU {
	var reps, ranks = make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		reps[i], ranks[i] = i, 1
	}

	return &IslandDSU{
		reps:  reps,
		ranks: ranks,
	}
}

func (d *IslandDSU) Find(group int) int {
	if d.reps[group] == group {
		return group
	}

	d.reps[group] = d.Find(d.reps[group])
	return d.reps[group]
}

func (d *IslandDSU) Union(a, b int) bool {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return false
	}

	switch {
	case d.ranks[ra] >= d.ranks[rb]:
		d.reps[rb] = d.reps[ra]
		d.ranks[ra] += d.ranks[rb]
	default:
		d.reps[ra] = d.reps[rb]
		d.ranks[rb] += d.ranks[ra]
	}

	return true
}

func reduceDimension(j, i, n int) int {
	return (j * n) + i
}
