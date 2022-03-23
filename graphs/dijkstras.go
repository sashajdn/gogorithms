package graphs

import (
	"container/heap"
	"math"
)

type heapItem struct {
	Node, Distance int
}
type heapItems struct {
	items []*heapItem
	ref   map[int]int
}

func (h heapItems) Len() int      { return len(h.items) }
func (h heapItems) Swap(i, j int) { h.items[i], h.items[j] = h.items[j], h.items[i] }
func (h heapItems) Less(i, j int) bool {
	if h.items[i].Distance == h.items[j].Distance {
		return h.items[i].Node < h.items[j].Node
	}

	return h.items[i].Distance < h.items[j].Distance
}
func (h heapItems) Pop() any {
	var output *heapItem
	output, h.items = h.items[h.Len()-1], h.items[:h.Len()-1]
	delete(h.ref, output.Node)
	return output
}
func (h heapItems) Push(item any) {
	v, ok := item.(*heapItem)
	if ok {
		h.items = append(h.items, v)
	}
}
func (h heapItems) SetDistance(node, distance int) {
	nodeIdx := h.ref[node]
	h.items[nodeIdx].Distance = distance
	h.Heapify() // Awkward, but one way to rebuild the heap.
}
func (h heapItems) Heapify() {
	heap.Init(h)
}

// DijkstrasAlgorithmHeap ...
//
// T -> O((e + v) * log(v))
// S -> O(v)
func DijkstrasAlgorithmHeap(start int, edges [][][]int) []int {
	// T -> O(v)
	var (
		distances = make([]int, len(edges))
		items     = make([]*heapItem, len(edges))
	)
	for i := 0; i < len(edges); i++ {
		distances[i] = math.MaxInt
		items[i] = &heapItem{
			Distance: math.MaxInt,
			Node:     i,
		}
	}
	items[start].Distance = 0

	// T -> O(vlog(v)), S -> O(v)
	var h = heapItems{
		items: items,
		ref:   map[int]int{},
	}
	h.Heapify()

	// T -> O(v), S -> O(1)
	var visited = map[int]struct{}{}
	for len(visited) < len(edges) {
		// T -> O(log(v)), S -> O(1)
		a := h.Pop()
		item, ok := a.(*heapItem)
		if !ok {
			// Should panic here.
			continue
		}

		if _, ok := visited[item.Node]; ok {
			continue
		}

		if item.Distance == math.MaxInt {
			break
		}

		visited[item.Node] = struct{}{}

		// T -> O(e), S -> O(1)
		for _, edge := range edges[item.Node] {
			destinationNode, distanceToDestination := edge[0], edge[1]
			if distanceToDestination+item.Distance < distances[destinationNode] {
				distances[destinationNode] = distanceToDestination + item.Distance
				// T -> O(log(v)) if we can directly siftup, we rebuild here which is O(nlog(v))
				h.SetDistance(destinationNode, distanceToDestination+item.Distance)
			}
		}
	}

	// T -> O(v), S -> O(1)
	for i, d := range distances {
		if d == math.MaxInt {
			distances[i] = -1
		}
	}

	return distances
}

// DijkstrasAlgorithmArray ...
//
// T -> O(v**2 + e)
// S -> O()
func DijkstrasAlgorithmArray(start int, edges [][][]int) []int {
	// T -> O(v), S -> (v)
	var distances = make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		distances[i] = math.MaxInt
	}
	distances[start] = 0

	// T -> O(v), S -> O(v)
	var visited = map[int]struct{}{}
	for len(visited) < len(edges) {
		// T -> O(v)
		node, minDistanceToNode := findMinDistanceNode(distances, visited)
		if minDistanceToNode == math.MaxInt {
			break
		}

		visited[node] = struct{}{}

		// T -> O(e)
		for _, edge := range edges[node] {
			destinationNode, distanceToDestination := edge[0], edge[1]
			if _, ok := visited[destinationNode]; ok {
				continue
			}

			distances[destinationNode] = min(distances[destinationNode], distanceToDestination+minDistanceToNode)
		}
	}

	// T -> O(v)
	for i, d := range distances {
		if d == math.MaxInt {
			distances[i] = -1
		}
	}

	return distances
}

func findMinDistanceNode(distances []int, visited map[int]struct{}) (int, int) {
	var (
		minSoFar = math.MaxInt
		node     int
	)
	for i, d := range distances {
		if _, ok := visited[i]; ok {
			continue
		}

		if d < minSoFar {
			minSoFar, node = d, i
		}
	}

	return node, distances[node]
}
