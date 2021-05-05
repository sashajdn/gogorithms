package arrays

import (
	"sort"
)

type Intervals [][]int

func (i Intervals) Len() int {
	return len(i)
}

func (i Intervals) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (i Intervals) Less(a, b int) bool {
	return i[a][0] < i[b][0]
}

// MergeOverlappingIntervals: T(N) -> O(n), S(N) -> O(n)
func MergeOverlappingIntervals(intervals Intervals) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Sort(intervals)
	currentInterval := intervals[0]
	merged := Intervals{currentInterval}

	for _, nextInterval := range intervals {
		if isOverlapping(currentInterval, nextInterval) {
			currentInterval = merge(currentInterval, nextInterval)
		} else {
			currentInterval = nextInterval
			merged = append(merged, currentInterval)
		}
	}
	return merged
}

// isOverlapping assumes that a, b are slices of len 2.
func isOverlapping(a, b []int) bool {
	return a[1] >= b[0]
}

// merge merges two intervals
func merge(a, b []int) []int {
	a[1] = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}(a[1], b[1])
	return a
}
