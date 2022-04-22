package graphs

import (
	"container/heap"
	"math"
)

func NewFlightsPriorityQueue(source, size int) *FlightsPriorityQueue {
	var (
		items      = make([]*FlightHeapItem, 0, size)
		references = make(map[int]int, size)
	)

	for i := 0; i < size; i++ {
		items = append(items, &FlightHeapItem{
			Value: math.MaxInt,
			Index: i,
		})
		references[i] = i
	}
	items[source].Value = 0

	pq := &FlightsPriorityQueue{}
	heap.Init(pq)

	return pq
}

type FlightHeapItem struct {
	Value, Index int
}

type FlightsPriorityQueue struct {
	items      []*FlightHeapItem
	references map[int]int
}

func (f FlightsPriorityQueue) Len() int           { return len(f.items) }
func (f FlightsPriorityQueue) Less(i, j int) bool { return f.items[i].Value < f.items[j].Value }
func (f FlightsPriorityQueue) Swap(i, j int) {
	f.references[f.items[i].Index] = j
	f.references[f.items[j].Index] = i
	f.items[j], f.items[i] = f.items[i], f.items[j]
}

func (f *FlightsPriorityQueue) Push(item interface{}) {
	flightItem := item.(*FlightHeapItem)
	f.items = append(f.items, flightItem)
}

func (f *FlightsPriorityQueue) Pop() interface{} {
	v := f.items[f.Len()-1]
	f.items = f.items[:f.Len()-1]

	delete(f.references, v.Index)

	return v
}

func (f *FlightsPriorityQueue) UpdateValueAtIndex(value, index int) {
	pqIndex := f.references[index]
	f.items[pqIndex].Value = value
	heap.Init(f)
}

// FindCheapestPrice ...
//
// T -> O((v + e) * log(v))
// S -> O(v)
func FindCheapestPrice_Dijkstras(n int, flights [][]int, src int, dst int, k int) int {
	var graph = make(map[int][][]int, n)
	for _, flight := range flights {
		from, to, cost := flight[0], flight[1], flight[2]

		if _, ok := graph[from]; !ok {
			graph[from] = [][]int{}
		}
		graph[from] = append(graph[from], []int{to, cost})

		if _, ok := graph[to]; !ok {
			graph[to] = [][]int{}
		}
	}

	var (
		costs  = make([]int, n)
		counts = make([]int, n)
	)
	for i := 0; i < n; i++ {
		costs[i] = math.MaxInt
	}
	costs[src] = 0

	// T -> O(vlogv)
	// S -> O(v)
	pq := NewFlightsPriorityQueue(src, n)

	// T -> O((v + e) * logv)
	// S -> O(1)
	for pq.Len() > 0 {
		next := heap.Pop(pq)
		heapItem := next.(*FlightHeapItem)

		if counts[heapItem.Index] >= k {
			continue
		}

		if heapItem.Value == math.MaxInt {
			break
		}

		for _, edge := range graph[heapItem.Index] {
			to, cost := edge[0], edge[1]

			totalCost := costs[heapItem.Index] + cost

			switch {
			case totalCost == costs[to]:
				counts[to]++
				continue
			case totalCost < costs[to]:
				costs[to] = totalCost
				counts[to] = 0
				pq.UpdateValueAtIndex(totalCost, to)
			}
		}
	}

	if costs[dst] == math.MaxInt {
		return -1
	}

	return costs[dst]
}

func FindCheapestPrice_BFS(n int, flights [][]int, src int, dst int, k int) int {
	var graph = make(map[int][][]int, n)
	for _, flight := range flights {
		from, to, cost := flight[0], flight[1], flight[2]

		if _, ok := graph[from]; !ok {
			graph[from] = [][]int{}
		}
		graph[from] = append(graph[from], []int{to, cost})

		if _, ok := graph[to]; !ok {
			graph[to] = [][]int{}
		}
	}

	var costs = make([]int, n)
	for i := 0; i < n; i++ {
		costs[i] = math.MaxInt
	}
	costs[src] = 0

	var (
		queue = []int{src}
		level int
	)
	for len(queue) > 0 && level < k {
		l := len(queue)
		for i := 0; i < l; i++ {
			var from int
			from, queue = queue[0], queue[1:]

			for _, edge := range graph[from] {
				to, cost := edge[0], edge[1]

				costs[to] = min(costs[to], cost+costs[from])

				queue = append(queue, to)
			}
		}

		level++
	}

	if costs[dst] == math.MaxInt {
		return -1
	}

	return costs[dst]
}
