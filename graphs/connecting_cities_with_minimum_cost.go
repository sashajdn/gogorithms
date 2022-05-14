package graphs

import "sort"

// ConnectingCitiesMinimumCost ...
//
// T -> O(elog(e) + elog(v))
// S -> O(v)
func ConnectingCitiesMinimumCost(n int, connections [][]int) int {
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	var (
		dsu               = NewCitiesDSU(n)
		mstNodes, mstCost int
	)
	for _, connection := range connections {
		if mstNodes == n-1 {
			break
		}

		from, to, cost := connection[0], connection[1], connection[2]

		if dsu.Union(from, to) {
			mstCost += cost
			mstNodes++
		}
	}

	if mstNodes < n-1 {
		return -1
	}

	return mstCost
}

type CitiesDSU struct {
	reps, ranks []int
}

func NewCitiesDSU(size int) *CitiesDSU {
	var reps, ranks = make([]int, size+1), make([]int, size+1)
	for i := 1; i < size+1; i++ {
		reps[i] = i
		ranks[i] = 1
	}

	return &CitiesDSU{
		reps:  reps,
		ranks: ranks,
	}
}

func (d *CitiesDSU) Find(group int) int {
	if d.reps[group] == group {
		return group
	}

	d.reps[group] = d.Find(d.reps[group])
	return d.reps[group]
}

func (d *CitiesDSU) Union(a, b int) bool {
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
