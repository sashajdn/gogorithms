package graphs

import "sort"

// MSTKruskals ...
//
// T -> O(e * (log(v) + log(e)))
// S -> O(v)
func MSTKruskals(n int, edges [][]int) (int, [][]int) {
	// T -> O(elog(e))
	// S -> O(1)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[i][2]
	})

	// T -> O(elog(v))
	// S -> O(v)
	var (
		mstCost, nodesVisited int
		mstSet                = make([][]int, 0, n-1)
		dsu                   = NewKruskalsDSU(n)
	)
	for _, edge := range edges {
		if nodesVisited == n {
			break
		}

		from, to, cost := edge[0], edge[1], edge[2]

		if dsu.Union(from, to) {
			mstCost += cost
			mstSet = append(mstSet, edge)
			nodesVisited++
		}
	}

	return mstCost, mstSet
}

type KruskalsDSU struct {
	reps, ranks []int
}

func NewKruskalsDSU(size int) *KruskalsDSU {
	var reps, ranks = make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		reps[i], ranks[i] = i, 1
	}

	return &KruskalsDSU{
		reps:  reps,
		ranks: ranks,
	}
}

func (k *KruskalsDSU) Find(group int) int {
	if k.reps[group] == group {
		return group
	}

	k.reps[group] = k.Find(k.reps[group])
	return k.reps[group]
}

func (k *KruskalsDSU) Union(a, b int) bool {
	ra, rb := k.Find(a), k.Find(b)
	if ra == rb {
		return false
	}

	if k.ranks[ra] >= k.ranks[rb] {
		k.reps[rb] = k.reps[ra]
		k.ranks[ra] += k.ranks[rb]
		return true
	}

	k.reps[ra] = k.reps[rb]
	k.ranks[rb] += k.ranks[ra]
	return true
}
