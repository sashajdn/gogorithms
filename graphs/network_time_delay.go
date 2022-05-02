package graphs

import (
	"container/heap"
	"math"
)

// NetworkDelayTime_Dijkstras ...
//
// T -> O((v + e)* log(v))
// S -> O(v + e)
func NetworkDelayTime_Dijkstras(times [][]int, n int, k int) int {
	// T -> O(e)
	// S -> O(v + e)
	var graph = make(map[int][][]int, n)
	for _, time := range times {
		from, to, cost := time[0], time[1], time[2]

		if _, ok := graph[from]; !ok {
			graph[from] = [][]int{}
		}
		graph[from] = append(graph[from], []int{to, cost})
	}

	var costs = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		costs[i] = math.MaxInt
	}
	costs[k] = 0

	var pq = &NetworkPQ{}
	heap.Push(pq, &NetworkPQItem{
		Index: k,
		Cost:  0,
	})

	// T -> O(elog(v) + vlog(v))
	// S -> O(v)
	var visited = map[int]struct{}{}
	for pq.Len() > 0 {
		next := heap.Pop(pq)
		from := next.(*NetworkPQItem)

		if _, ok := visited[from.Index]; ok {
			continue
		}
		visited[from.Index] = struct{}{}

		// elogv
		for _, edge := range graph[from.Index] {
			to, cost := edge[0], edge[1]

			if from.Cost+cost < costs[to] {
				costs[to] = from.Cost + cost

				heap.Push(pq, &NetworkPQItem{
					Index: to,
					Cost:  costs[to],
				})
			}
		}
	}

	// T -> O(n)
	var maxTime int
	for i := 1; i < n+1; i++ {
		if costs[i] == math.MaxInt {
			return -1
		}

		maxTime = max(maxTime, costs[i])
	}

	return maxTime
}

type NetworkPQItem struct {
	Index, Cost int
}

type NetworkPQ []*NetworkPQItem

func (p NetworkPQ) Len() int           { return len(p) }
func (p NetworkPQ) Less(i, j int) bool { return p[i].Cost < p[j].Cost }
func (p NetworkPQ) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *NetworkPQ) Pop() interface{} {
	v := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return v
}
func (p *NetworkPQ) Push(v interface{}) {
	vv := v.(*NetworkPQItem)
	*p = append(*p, vv)
}
