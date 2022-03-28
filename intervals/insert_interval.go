package intervals

// InsertIntervals ...
//
// T -> O(n)
// S -> O(1)
func InsertIntervals(intervals [][]int, newInterval []int) [][]int {
	var (
		output       = make([][]int, 0, len(intervals)+1)
		currentIndex int
	)

	// Find first inbetween or after.
	for currentIndex < len(intervals) {
		currentInterval := intervals[currentIndex]

		if currentInterval[0] < newInterval[0] {
			output = append(output, currentInterval)
			currentIndex++
			continue
		}

		break
	}

	// Append new interval or merge.
	switch {
	case len(output) == 0:
		output = append(output, newInterval)
	case newInterval[0] > output[len(output)-1][1]:
		output = append(output, newInterval)
	default:
		output[len(output)-1][1] = max(output[len(output)-1][1], newInterval[1])
	}

	// Append or merge rest of intervals.
	for currentIndex < len(intervals) {
		currentInterval := intervals[currentIndex]

		switch {
		case currentInterval[0] > output[len(output)-1][1]:
			output = append(output, currentInterval)
		default:
			output[len(output)-1][1] = max(output[len(output)-1][1], currentInterval[1])
		}

		currentIndex++
	}

	return output
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
