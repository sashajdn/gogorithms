package graphs

import (
	"container/heap"
	"sort"
)

// MinCostConnectAllPoints_PrimsLazy ...
//
// T -> O(Elog(E)) ~> O((n ** 2) * log(n))
// S -> O(E) -> O(n ** 2)
func MinCostConnectAllPoints_PrimsLazy(points [][]int) int {
	var (
		visited = make([]bool, len(points))
		pq      = &PointEdgePQ{}
	)
	heap.Push(pq, &PointEdge{
		From: 0,
		To:   0,
		Cost: 0,
	})

	// T -> O(Elog(E)) -> O((n**2)log(n ** 2)) ~> O((n ** 2) log(n))
	// S -> O(E) -> O(n ** 2)
	var edgeCount, mstCost int
	for edgeCount < len(points) {
		next := heap.Pop(pq)
		edge := next.(*PointEdge)

		if visited[edge.To] {
			continue
		}
		visited[edge.To] = true

		edgeCount++
		mstCost += edge.Cost

		for i := 0; i < len(points); i++ {
			if visited[i] {
				continue
			}

			p1, p2 := points[edge.To], points[i]
			cost := intAbs(p1[0]-p2[2]) + intAbs(p1[1]-p2[1])

			heap.Push(pq, &PointEdge{
				To:   i,
				Cost: cost,
			})
		}
	}

	return mstCost
}

type PointEdgePQ []*PointEdge

func (p PointEdgePQ) Len() int           { return len(p) }
func (p PointEdgePQ) Less(i, j int) bool { return p[i].Cost < p[j].Cost }
func (p PointEdgePQ) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *PointEdgePQ) Push(v interface{}) {
	vv := v.(*PointEdge)
	*p = append(*p, vv)
}
func (p *PointEdgePQ) Pop() interface{} {
	v := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return v
}

// MinCostConnectAllPoints_Kruskals ...
//
// T -> O(elog(e) + elog(v)) where e is the number of edges and is equal to (n ** 2) since we have a complete graph.
// S -> O(n ** 2)
func MinCostConnectAllPoints_Kruskals(points [][]int) int {
	// T -> O(n ** 2)
	// S -> O(n ** 2)
	var edgeList []*PointEdge
	for j := 0; j < len(points); j++ {
		for i := j + 1; i < len(points); i++ {
			edgeList = append(edgeList, &PointEdge{
				From: j,
				To:   i,
				Cost: intAbs(points[j][0]-points[i][0]) + intAbs(points[j][1]-points[i][1]),
			})
		}
	}

	// T -> O(elog(e))
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i].Cost < edgeList[j].Cost
	})

	// T -> O(elog(v))
	// S -> O(1)
	var (
		dsu                     = NewPointDSU(len(points))
		totalCost, nodesVisited int
	)
	for _, edge := range edgeList {
		if nodesVisited == len(points)-1 {
			break
		}

		if !dsu.Union(edge.From, edge.To) {
			continue
		}

		totalCost += edge.Cost
		nodesVisited++
	}

	return totalCost
}

type PointEdge struct {
	From, To int
	Cost     int
}

type PointDSU struct {
	reps, ranks []int
}

func NewPointDSU(size int) *PointDSU {
	var ranks, reps = make([]int, size), make([]int, size)
	for i := 0; i < size; i++ {
		ranks[i], reps[i] = 1, i
	}

	return &PointDSU{
		reps:  reps,
		ranks: ranks,
	}
}

func (p *PointDSU) Find(group int) int {
	if p.reps[group] == group {
		return group
	}

	p.ranks[group] = p.Find(p.ranks[group])
	return p.ranks[group]
}

func (p *PointDSU) Union(a, b int) bool {
	ra, rb := p.Find(a), p.Find(b)
	if ra == rb {
		return false
	}

	if p.ranks[ra] >= p.ranks[rb] {
		p.reps[rb] = p.reps[ra]
		p.ranks[ra] += p.ranks[rb]
		return true
	}

	p.reps[ra] = p.reps[rb]
	p.ranks[rb] += p.ranks[ra]
	return true
}

func intAbs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}
