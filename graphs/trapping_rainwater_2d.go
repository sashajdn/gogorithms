package graphs

import (
	"container/heap"
	"fmt"
)

type PriorityQueueItem struct {
	x, y   int
	height int
}

type PriorityQueue []*PriorityQueueItem

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].height < p[j].height
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PriorityQueue) Len() int { return len(p) }

func (p PriorityQueue) Push(value interface{}) {
	pqi := value.(*PriorityQueueItem)
	p = append(p, pqi)
}

func (p PriorityQueue) Pop() interface{} {
	v := p[0]
	p = p[:p.Len()]
	return v
}

// TrappingRainWater ...
// T -> O(nm * log(n + m))
// S -> O(nm)
func TrappingRainWater(heights [][]int) int {
	if len(heights) < 3 {
		return 0
	}

	var pq = make(PriorityQueue, 0, 2*len(heights)+2*len(heights[0]))
	var visited = map[string]struct{}{}

	// T -> O(n)
	// S -> O(n)
	for j := 0; j < len(heights); j++ {
		pq = append(pq, &PriorityQueueItem{
			x:      0,
			y:      j,
			height: heights[j][0],
		})
		visited[hash(0, j)] = struct{}{}

		pq = append(pq, &PriorityQueueItem{
			x:      len(heights[0]) - 1,
			y:      j,
			height: heights[j][len(heights[0])-1],
		})
		visited[hash(len(heights[0])-1, j)] = struct{}{}
	}
	for i := 1; i < len(heights)-1; i++ {
		pq = append(pq, &PriorityQueueItem{
			x:      i,
			y:      0,
			height: heights[0][i],
		})
		visited[hash(i, 0)] = struct{}{}

		pq = append(pq, &PriorityQueueItem{
			x:      i,
			y:      len(heights) - 1,
			height: heights[len(heights)-1][i],
		})
		visited[hash(i, len(heights)-1)] = struct{}{}
	}
	heap.Init(pq)

	var (
		dx   = []int{-1, 1, 0, 0}
		dy   = []int{0, 0, -1, +1}
		area int
	)

	// T -> O(nm)
	// S -> O(n + m)
	for pq.Len() > 0 {
		// T -> O(log(n + m))
		s := pq.Pop()

		smallestItem := s.(*PriorityQueueItem)

		// O(log(n + m))
		for k := 0; k < 4; k++ {
			nx, ny := smallestItem.x+dx[k], smallestItem.y+dy[k]

			if nx < 0 || ny < 0 || nx > len(heights[0])-1 || ny > len(heights)-1 {
				continue
			}

			if _, ok := visited[hash(nx, ny)]; ok {
				continue
			}

			area += max(0, smallestItem.height-heights[ny][nx])

			pq.Push(&PriorityQueueItem{
				x:      nx,
				y:      ny,
				height: max(smallestItem.height, heights[ny][nx]),
			})

			visited[hash(nx, ny)] = struct{}{}
		}
	}

	return 0
}

func hash(i, j int) string {
	return fmt.Sprintf("%d%d", i, j)
}
