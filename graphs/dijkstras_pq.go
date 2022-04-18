package graphs

import (
	"container/heap"
	"math"
)

// DijkstrasPQ ...
// T -> O((V + E) * log(V)) where `V` is the number of vertices & `E` the number of edges in G(V, E).
// S -> O(V)
func DijkstrasPQ(source int, edges [][][]int) []int {
	var (
		distances = make([]int, len(edges))
		pq        = NewPQ(source, len(edges))
	)
	for i := 0; i < len(edges); i++ {
		distances[i] = math.MaxInt
	}
	distances[source] = 0

	// T -> O((V + E) * log(V))
	// Since there is a total of E edges in G(V, E), we have (V + E) here rather than (VE)
	for pq.Len() != 0 {
		next := heap.Pop(pq)
		item := next.(*PQItem)
		from := item.Index

		if item.Value == math.MaxInt {
			break
		}

		// T -> O(log(V) * e)  where `e` if the number of outbound edges at vertex, V.
		for _, edge := range edges[from] {
			to, distance := edge[0], edge[1]

			pathDistance := distances[from] + distance
			if pathDistance < distances[to] {
				distances[to] = pathDistance
				pq.UpdateValueAtIndex(pathDistance, to)
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

type PQItem struct {
	Value, Index int
}

func NewPQ(source, size int) *PQ {
	var (
		items      = make([]*PQItem, 0, size)
		references = make(map[int]int, size)
	)
	for i := 0; i < size; i++ {
		items = append(items, &PQItem{
			Value: math.MaxInt,
			Index: i,
		})
		references[i] = i
	}
	items[source].Value = 0

	pq := &PQ{
		items:      items,
		references: references,
	}
	heap.Init(pq)

	return pq
}

type PQ struct {
	items      []*PQItem
	references map[int]int
}

func (p PQ) Len() int           { return len(p.items) }
func (p PQ) Less(i, j int) bool { return p.items[i].Value < p.items[j].Value }
func (p PQ) Swap(i, j int) {
	p.references[p.items[j].Index] = i
	p.references[p.items[i].Index] = j
	p.items[i], p.items[j] = p.items[j], p.items[i]
}
func (p *PQ) Pop() interface{} {
	v := p.items[p.Len()-1]
	p.items = p.items[:p.Len()-1]
	return v
}
func (p *PQ) Push(value interface{}) {
	vv := value.(*PQItem)
	p.items = append(p.items, vv)
}
func (p *PQ) UpdateValueAtIndex(value, index int) {
	p.items[p.references[index]].Value = value
	heap.Init(p)
}
