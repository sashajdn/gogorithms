package intervals

import (
	"container/heap"
	"sort"
)

// MinNumberOfMeetingRooms ...
//
// T -> O(nlog(n)) where `n` is the number of meetings; we have to sort them.
// S -> O(n) where `n` is the number of intervals; in the worst case when all meetings are overlapping.
func MinNumberOfMeetingRooms(intervals [][]int) int {
	if len(intervals) < 2 {
		return len(intervals)
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	pq := &PriorityQueue{}
	heap.Push(pq, intervals[0][1])

	var maxConcurrentMeetings = 1
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]

		if current[0] < pq.Peek() {
			heap.Push(pq, current[1])
			maxConcurrentMeetings = max(maxConcurrentMeetings, pq.Len())
			continue
		}

		heap.Pop(pq)
		heap.Push(pq, current[1])
	}

	return maxConcurrentMeetings
}

type PriorityQueue []int

func (p PriorityQueue) Len() int           { return len(p) }
func (p PriorityQueue) Less(i, j int) bool { return p[i] < p[j] }
func (p PriorityQueue) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *PriorityQueue) Pop() interface{} {
	if p.Len() == 0 {
		return nil
	}

	v := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return v
}
func (p *PriorityQueue) Push(value interface{}) {
	vi := value.(int)
	*p = append(*p, vi)
}
func (p *PriorityQueue) Peek() int {
	if p.Len() == 0 {
		return 0
	}

	return (*p)[0]
}
