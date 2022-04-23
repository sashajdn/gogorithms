package graphs

import (
	"container/heap"
	"math"
)

type FlightHeapItem struct {
	Cost, Index, Stops int
}

type FlightsPriorityQueue []*FlightHeapItem

func (f FlightsPriorityQueue) Len() int           { return len(f) }
func (f FlightsPriorityQueue) Less(i, j int) bool { return f[i].Cost < f[j].Cost }
func (f FlightsPriorityQueue) Swap(i, j int) {
	f[j], f[i] = f[i], f[j]
}

func (f *FlightsPriorityQueue) Push(item interface{}) {
	flightItem := item.(*FlightHeapItem)
	*f = append(*f, flightItem)
}

func (f *FlightsPriorityQueue) Pop() interface{} {
	v := (*f)[f.Len()-1]
	*f = (*f)[:f.Len()-1]
	return v
}

// FindCheapestPrice ...
//
// T -> O((k + e) * log(k))
// S -> O(k + e)
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
		costs = make([]int, n)
		stops = make([]int, n)
	)
	for i := 0; i < n; i++ {
		costs[i] = math.MaxInt
		stops[i] = math.MaxInt
	}
	costs[src] = 0
	stops[src] = 0

	// T -> O(vlogv)
	// S -> O(v)
	pq := &FlightsPriorityQueue{
		{
			Cost:  0,
			Index: src,
			Stops: 0,
		},
	}

	// T -> O((v + e) * logv)
	// S -> O(1)
	for pq.Len() > 0 {
		next := heap.Pop(pq)
		from := next.(*FlightHeapItem)

		if from.Index == dst {
			break
		}

		if from.Stops >= k+1 {
			continue
		}

		for _, edge := range graph[from.Index] {
			to, cost := edge[0], edge[1]

			nextCost := from.Cost + cost
			nextStop := from.Stops + 1

			if nextCost < costs[to] {
				costs[to] = nextCost
				stops[to] = nextStop
				heap.Push(pq, &FlightHeapItem{
					Cost:  nextCost,
					Stops: nextStop,
					Index: to,
				})
				continue
			}

			if nextStop < stops[to] {
				heap.Push(pq, &FlightHeapItem{
					Cost:  nextCost,
					Stops: nextStop,
					Index: to,
				})
			}
		}
	}

	if costs[dst] == math.MaxInt {
		return -1
	}

	return costs[dst]
}

// FindCheapestPrice_BFS ...
// T -> O(v * k)
// S -> O(n)
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

	var (
		queue = []*FlightHeapItem{
			{
				Cost:  0,
				Index: src,
			},
		}
		level    int
		minPrice = math.MaxInt
	)
	for len(queue) > 0 && level != k+1 {
		l := len(queue)
		for i := 0; i < l; i++ {
			var from *FlightHeapItem
			from, queue = queue[0], queue[1:]

			if from.Index == dst {
				continue
			}

			for _, edge := range graph[from.Index] {
				to, cost := edge[0], edge[1]

				minPrice = min(minPrice, cost+from.Cost)

				queue = append(queue, &FlightHeapItem{
					Cost:  cost + from.Cost,
					Index: to,
				})
			}
		}

		level++
	}

	if minPrice == math.MaxInt {
		return -1
	}

	return minPrice
}

// FindCheapestPrice_BellmanFord ...
//
// T -> O(k * e)
// S -> O(v)
func FindCheapestPrice_BellmanFord(n int, flights [][]int, src int, dst int, k int) int {
	// S -> O(v)
	var costs = make([]int, n)
	for i := 0; i < n; i++ {
		costs[i] = math.MaxInt
	}
	costs[src] = 0

	// T -> O(k)
	for i := 0; i < k+1; i++ {
		tmp := make([]int, n)
		copy(tmp, costs)

		// T -> O(e)
		for _, flight := range flights {
			from, to, cost := flight[0], flight[1], flight[2]

			if costs[from] == math.MaxInt {
				continue
			}

			if costs[from]+cost < tmp[to] {
				tmp[to] = costs[from] + cost
			}
		}

		costs = tmp
	}

	if costs[dst] == math.MaxInt {
		return -1
	}

	return costs[dst]
}
