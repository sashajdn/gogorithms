package graphs

import (
	"math"
)

// DijkstraAlgorithm ...
//
// T -> O(log((v + e) * log(v))) where `v` is the number of vertices in the graph & `e` is the number of edges.
// S -> O(v)
func DijkstraAlgorithm(source int, edges [][][]int) []int {
	var (
		distances = make([]int, len(edges))
		heapArray = make([]int, len(edges))
	)

	for i := 0; i < len(edges); i++ {
		distances[i] = math.MaxInt
		heapArray[i] = math.MaxInt
	}

	distances[source], heapArray[source] = 0, 0
	heap := NewMinHeap(heapArray)

	visited := map[int]struct{}{}

	// T -> O(v)
	// TT -> O(v  * (elog(e))) -> O((v + e) * log(e))
	for len(visited) < len(edges) {
		// T -> O(log(v))
		node, ok := heap.Pop()
		if !ok {
			panic("Expecting heap to have at least one item; empty")
		}

		if _, ok := visited[node.Index]; ok {
			continue
		}

		if node.Value == math.MaxInt {
			break
		}

		visited[node.Index] = struct{}{}

		// T -> O(subset of e)
		for _, edge := range edges[node.Index] {
			to, cost := edge[0], edge[1]

			if node.Value+cost >= distances[to] {
				continue
			}

			// T -> O(log(v))
			distances[to] = node.Value + cost
			heap.UpdateValueAtIndex(to, distances[to])
		}
	}

	// T -> O(v)
	for i := range distances {
		if distances[i] == math.MaxInt {
			distances[i] = -1
		}
	}

	return distances
}

type HeapItem struct {
	Index, Value int
}

func NewMinHeap(array []int) *MinHeap {
	references := map[int]int{}
	heapArray := make([]*HeapItem, 0, len(array))
	for i, v := range array {
		references[i] = i

		heapArray = append(heapArray, &HeapItem{
			Index: i,
			Value: v,
		})
	}

	heap := &MinHeap{
		h: heapArray,
		r: references,
	}
	heap.Heapify()

	return heap
}

type MinHeap struct {
	h []*HeapItem
	r map[int]int
}

func (m MinHeap) Heapify() {
	rightMostParent := (len(m.h) - 1 - 1) / 2

	for k := rightMostParent; k >= 0; k-- {
		m.siftDown(k, len(m.h)-1)
	}
}

func (m MinHeap) Pop() (*HeapItem, bool) {
	if len(m.h) == 0 {
		return nil, false
	}

	m.swap(0, len(m.h)-1)

	var popped *HeapItem
	popped, m.h = m.h[len(m.h)-1], m.h[:len(m.h)-1]

	m.siftDown(0, len(m.h)-1)
	return popped, true
}

func (m MinHeap) Push(value int) {
	m.h = append(m.h, &HeapItem{
		Value: value,
		Index: len(m.h),
	})

	m.siftUp(len(m.h) - 1)
}

func (m MinHeap) UpdateValueAtIndex(index, value int) bool {
	latestIndex, ok := m.r[index]
	if !ok {
		return false
	}

	m.h[latestIndex].Value = value
	m.siftUp(latestIndex)
	return true
}

func (m MinHeap) siftUp(index int) {
	currentIndex := index
	for currentIndex >= 0 {
		parentIndex := (currentIndex - 1) / 2

		if m.h[parentIndex].Value < m.h[currentIndex].Value {
			return
		}

		m.swap(currentIndex, parentIndex)
		currentIndex = parentIndex
	}
}

func (m MinHeap) siftDown(index, endIndex int) {
	currentIndex := index
	for currentIndex <= len(m.h)-1 {
		leftChildIndex, rightChildIndex := 2*currentIndex+1, 2*currentIndex+2

		minIndex := m.indexWithMinValue(leftChildIndex, rightChildIndex)
		if minIndex == -1 {
			return
		}

		if m.h[currentIndex].Value < m.h[minIndex].Value {
			m.swap(currentIndex, minIndex)
			currentIndex = minIndex
			continue
		}

		return
	}
}

func (m MinHeap) swap(i, j int) {
	m.h[i], m.h[j] = m.h[j], m.h[i]
	m.r[i] = j
	m.r[j] = i
}

func (m MinHeap) indexWithMinValue(left, right int) int {
	if left > (len(m.h) - 1) {
		return -1
	}

	if right > (len(m.h) - 1) {
		return left
	}

	if m.h[left].Value < m.h[right].Value {
		return left
	}

	return right
}
