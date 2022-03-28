package intervals

import "sort"

// NonOverlappingIntervals ...
//
// T -> O(log(n))
// S -> O(1)
func NonOverlappingIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	var (
		count        = 1
		lastInterval = intervals[0]
	)
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]

		if current[0] >= lastInterval[1] {
			count++
			lastInterval = current
		}
	}

	return len(intervals) - count
}
