package graphs

import "container/heap"

// MinimumCostToReachDestination ...
//
// T -> O((e + v) * log(v))
// S -> O(v + e)
func MinimumCostToReachDestination(maxTime int, edges [][]int, passingFees []int) int {
	var graph = make(map[int][][]int, len(passingFees))
	for _, edge := range edges {
		from, to, time := edge[0], edge[1], edge[2]

		if _, ok := graph[from]; !ok {
			graph[from] = [][]int{}
		}
		graph[from] = append(graph[from], []int{to, time})

		if _, ok := graph[to]; !ok {
			graph[to] = [][]int{}
		}
		graph[to] = append(graph[to], []int{from, time})
	}

	var pq = &TimePQ{}
	heap.Push(pq, &TimePQItem{
		Index:     0,
		Cost:      passingFees[0],
		TimeSoFar: 0,
	})

	var visited = make(map[int]int, len(passingFees))
	for pq.Len() > 0 {
		next := heap.Pop(pq)
		from := next.(*TimePQItem)

		if from.TimeSoFar > maxTime {
			continue
		}

		if from.Index == len(passingFees)-1 {
			return from.Cost
		}

		if minTime, ok := visited[from.Index]; ok && from.TimeSoFar >= minTime {
			continue
		}
		visited[from.Index] = from.TimeSoFar

		for _, edge := range graph[from.Index] {
			to, time := edge[0], edge[1]
			cost := passingFees[to]

			heap.Push(pq, &TimePQItem{
				Index:     to,
				TimeSoFar: from.TimeSoFar + time,
				Cost:      from.Cost + cost,
			})
		}
	}

	return -1
}

type TimePQItem struct {
	Index, TimeSoFar, Cost int
}

type TimePQ []*TimePQItem

func (t TimePQ) Len() int           { return len(t) }
func (t TimePQ) Less(i, j int) bool { return t[i].Cost < t[j].Cost }
func (t TimePQ) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t *TimePQ) Pop() interface{} {
	v := (*t)[t.Len()-1]
	*t = (*t)[:t.Len()-1]
	return v
}
func (t *TimePQ) Push(v interface{}) {
	vv := v.(*TimePQItem)
	*t = append(*t, vv)
}
