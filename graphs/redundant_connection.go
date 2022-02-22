package graphs

import "fmt"

// FindRedundantConnection ...
//
// T: O(e) where e is the number of edges.
// S: O(v) where v is the number of vertices
func FindRedundantConnection(edges [][]int) []int {
	nodes := map[int]int{}

	for _, edge := range edges {
		fro, to := edge[0], edge[1]

		if _, ok := nodes[fro]; !ok {
			nodes[fro] = len(nodes)
		}

		if _, ok := nodes[to]; !ok {
			nodes[to] = len(nodes)
		}

	}

	var redundantEdge []int
	dsu := BuildRedundantDSU(len(nodes))
	for _, edge := range edges {
		fro, to := edge[0], edge[1]
		froID, toID := nodes[fro], nodes[to]

		isRedundantConnection := dsu.Union(froID, toID)
		if isRedundantConnection {
			redundantEdge = []int{fro, to}
		}
	}

	return redundantEdge
}

type RedundantDSU struct {
	representations []int
	sizes           []int
}

func BuildRedundantDSU(size int) *RedundantDSU {
	var (
		r = make([]int, size)
		s = make([]int, size)
	)

	for i := 0; i < size; i++ {
		r[i] = i
		s[i] = 1
	}

	fmt.Println(r, s)

	return &RedundantDSU{
		representations: r,
		sizes:           s,
	}
}

func (r *RedundantDSU) Find(group int) int {
	if r.representations[group] == group {
		return group
	}

	r.representations[group] = r.Find(r.representations[group])
	return r.representations[group]
}

func (r *RedundantDSU) Union(groupA, groupB int) bool {
	ra, rb := r.Find(groupA), r.Find(groupB)
	if ra == rb {
		return true
	}

	if r.sizes[ra] >= r.sizes[rb] {
		r.representations[rb] = r.representations[ra]
		r.sizes[ra] += r.sizes[rb]
		return false
	}

	r.representations[ra] = r.representations[rb]
	r.sizes[rb] += r.sizes[ra]
	return false
}
