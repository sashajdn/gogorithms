package graphs

import "container/heap"

// SwimInWater ...
//
// T -> O(N**2 * log(N))
// S -> O(N ** 2)
func SwimInWater_Dijkstras(grid [][]int) int {
	var g = map[int][][]int{}
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			key := j*len(grid[0]) + i

			g[key] = [][]int{}
			for _, neighbour := range fetchSwimNeighbours(grid, j, i) {
				dj, di := neighbour[0], neighbour[1]
				neighbourKey := dj*len(grid[0]) + di
				g[key] = append(g[key], []int{neighbourKey, grid[dj][di]})
			}
		}
	}

	var pq = &PQ{}
	heap.Push(pq, &PQSwimItem{
		Index: 0,
		Value: grid[0][0],
		Bound: grid[0][0],
	})

	var (
		visited = map[int]struct{}{}
		dst     = len(grid)*len(grid[0]) - 1
	)
	for pq.Len() > 0 {
		next := heap.Pop(pq)
		from := next.(*PQSwimItem)

		if from.Index == dst {
			return from.Bound
		}

		if _, ok := visited[from.Index]; ok {
			continue
		}
		visited[from.Index] = struct{}{}

		for _, edge := range g[from.Index] {
			to, cost := edge[0], edge[1]

			heap.Push(pq, &PQSwimItem{
				Index: to,
				Value: cost,
				Bound: max(from.Bound, cost),
			})
		}
	}

	return -1
}

func fetchSwimNeighbours(grid [][]int, j, i int) [][]int {
	var directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var validNeighbours = make([][]int, 0, 4)
	for _, d := range directions {
		dj, di := j+d[0], i+d[1]

		if dj < 0 || dj > len(grid)-1 {
			continue
		}
		if di < 0 || di > len(grid[0])-1 {
			continue
		}

		validNeighbours = append(validNeighbours, []int{dj, di})
	}

	return validNeighbours
}

type PQSwimItem struct {
	Index, Value, Bound int
}

type PQSwim []*PQItem

func (p PQSwim) Len() int           { return len(p) }
func (p PQSwim) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PQSwim) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *PQSwim) Pop() interface{} {
	v := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return v
}
func (p *PQSwim) Push(v interface{}) {
	vv := v.(*PQItem)
	*p = append(*p, vv)
}
