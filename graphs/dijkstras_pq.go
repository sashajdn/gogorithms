package graphs

import (
	"container/heap"
	"math"
)

type PQItem struct {
	Value  int
	Vertex int
}

type DijkstrasPQ struct {
	h       []*PQItem
	indexes map[int]int
}

func (d DijkstrasPQ) Len() int           { return len(d.h) }
func (d DijkstrasPQ) Less(i, j int) bool { return d.h[i].Value < d.h[j].Value }
func (d DijkstrasPQ) Swap(i, j int) {
	d.h[i], d.h[j] = d.h[j], d.h[i]

	d.indexes[i] = d.indexes[j]
	d.indexes[j] = d.indexes[i]
}
func (d *DijkstrasPQ) Pop() interface{} {
	v := d.h[d.Len()-1]
	d.h = d.h[:d.Len()-1]
	return v
}
func (d *DijkstrasPQ) Push(v interface{}) {
	vv := v.(*PQItem)
	d.h = append(d.h, vv)
}

func (d *DijkstrasPQ) UpdateValueAtIndex(index, value int) {
	internalIndex := d.indexes[index]
	d.h[internalIndex].Value = value
	heap.Init(d)
}

// DijkstrasPq ...
//
// T -> O((v + e) * log(v))
// S -> O(v)
func DijkstrasPq(source int, edges [][][]int) []int {
	var (
		distances = make([]int, len(edges))
		h         = make([]*PQItem, len(edges))
		indexes   = make(map[int]int)
	)

	// T -> O(e)
	// S -> O(v)
	for i := 0; i < len(edges); i++ {
		distances[i] = math.MaxInt
		h[i] = &PQItem{
			Value:  math.MaxInt,
			Vertex: i,
		}
		indexes[i] = i
	}

	distances[source] = 0
	h[source].Value = 0

	// T -> O(vlog(v))
	pq := &DijkstrasPQ{
		h:       h,
		indexes: indexes,
	}
	heap.Init(pq)

	// T -> (e log(v))
	// S -> O(v)
	var visited = map[int]struct{}{}
	for len(visited) < len(edges) {
		next := heap.Pop(pq)
		from := next.(*PQItem)

		if from.Value == math.MaxInt {
			break
		}

		if _, ok := visited[from.Vertex]; ok {
			continue
		}
		visited[from.Vertex] = struct{}{}

		for _, edge := range edges[from.Vertex] {
			to, cost := edge[0], edge[1]

			if distances[from.Vertex]+cost < distances[to] {
				distances[to] = distances[from.Vertex] + cost
				pq.UpdateValueAtIndex(to, distances[to])
			}
		}
	}

	for i := 0; i < len(distances); i++ {
		if distances[i] == math.MaxInt {
			distances[i] = -1
		}
	}

	return distances
}
