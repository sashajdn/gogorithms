package arrays

import (
	"math"
	"sort"
)

// SmallestDifference T -> O(nlog(n) + mlog(m)), S -> O(1)
func SmallestDifference(a, b []int) []int {
	r := []int{}

	minSoFar := math.MaxInt32
	var curr int

	idxA, idxB := 0, 0

	sort.Ints(a)
	sort.Ints(b)

	for idxA < len(a) && idxB < len(b) {
		f, s := a[idxA], b[idxB]
		curr = absDiff(f, s)

		if curr == 0 {
			r = []int{f, s}
			return r
		}

		if curr < minSoFar {
			minSoFar = curr
			r = []int{f, s}
		}

		if f < s {
			idxA++
		} else {
			idxB++
		}
	}

	return r
}

func absDiff(a, b int) int {
	if a < b {
		return b - a

	}
	return a - b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
