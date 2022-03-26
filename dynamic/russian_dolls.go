package dynamic

import (
	"sort"
)

// MaxRussianDolls ...
//
// T -> O(nlog(n))
// S -> O(n)
func MaxRussianDolls(dolls [][]int) int {
	if len(dolls) == 0 {
		return 0
	}

	// T -> O(nlog(n))
	sort.Slice(dolls, func(i, j int) bool {
		if dolls[i][0] == dolls[j][0] {
			return dolls[i][1] > dolls[j][1]
		}

		return dolls[i][0] < dolls[j][0]
	})

	var (
		increasingSeq []int
		heights       = make([]int, 0, len(dolls))
	)

	// T -> O(n)
	// S -> O(n)
	for i := 0; i < len(dolls); i++ {
		heights = append(heights, dolls[i][1])
	}

	// Longest increasing subsequence problem.
	for i := 0; i < len(heights); i++ {
		idx := bisectLeft(increasingSeq, heights[i])
		if idx == len(increasingSeq) {
			increasingSeq = append(increasingSeq, heights[i])
			continue
		}

		increasingSeq[idx] = heights[i]
	}

	return len(increasingSeq)
}

func bisectLeft(array []int, target int) int {
	l, r := 0, len(array)
	for l < r {
		m := (l + r) / 2

		if target == array[m] {
			return m
		}

		if target < array[m] {
			r = m
		}

		l = m + 1
	}

	return l
}
