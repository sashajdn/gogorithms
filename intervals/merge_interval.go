package intervals

import (
	"sort"
)

// MergeInterval ...
//
// T -> O(nlog(n)) where `n` is the number of intervals in the array - since we have to sort it initially.
// S -> O(1)
func MergeInterval(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// T -> O(nlog(n))
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// T -> O(n)
	var output = [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		currentInterval := intervals[i]
		switch {
		case currentInterval[0] > output[len(output)-1][1]:
			output = append(output, currentInterval)
		default:
			output[len(output)-1][1] = max(output[len(output)-1][1], currentInterval[1])
		}
	}

	return output
}
