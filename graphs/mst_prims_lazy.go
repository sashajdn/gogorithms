package graphs

import "container/heap"

// MSTPrimsLazy ...
//
// T -> O(E + V * log(e))
// S -> O(V + E)
func MSTPrimsLazy(n int, edges [][]int) (int, [][]int) {
	var (
		graph   = map[int][]*PrimsLazyEdge{}
		minEdge *PrimsLazyEdge
	)
	for _, edge := range edges {
		from, to, cost := edge[0], edge[1], edge[2]
		edge := &PrimsLazyEdge{
			From: from,
			To:   to,
			Cost: cost,
		}

		if _, ok := graph[from]; !ok {
			graph[from] = []*PrimsLazyEdge{}
		}
		graph[from] = append(graph[from], edge)

		if minEdge == nil {
			minEdge = edge
			continue
		}

		if cost < minEdge.Cost {
			minEdge = edge
		}
	}

	var (
		mstCost             int
		mst                 = make([][]int, 0, n-1)
		mstSet              = map[int]struct{}{}
		pq                  = &PrimsLazyPQ{}
		numberOfVistedNodes int
	)
	heap.Push(pq, minEdge)

	for pq.Len() > 0 {
		if numberOfVistedNodes == n-1 {
			break
		}

		next := heap.Pop(pq)
		edge := next.(*PrimsLazyEdge)

		if _, ok := mstSet[edge.To]; ok {
			continue
		}
		mstSet[edge.To] = struct{}{}

		mstCost += edge.Cost
		mst = append(mst, []int{edge.From, edge.To, edge.Cost})
		numberOfVistedNodes++

		for _, adjacentEdge := range graph[edge.To] {
			heap.Push(pq, adjacentEdge)
		}
	}

	return mstCost, mst
}

type PrimsLazyEdge struct {
	From, To int
	Cost     int
}

type PrimsLazyPQ []*PrimsLazyEdge

func (p PrimsLazyPQ) Len() int           { return len(p) }
func (p PrimsLazyPQ) Less(i, j int) bool { return p[i].Cost < p[j].Cost }
func (p PrimsLazyPQ) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *PrimsLazyPQ) Push(v interface{}) {
	vv := v.(*PrimsLazyEdge)
	*p = append(*p, vv)
}
func (p *PrimsLazyPQ) Pop() interface{} {
	v := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return v
}
