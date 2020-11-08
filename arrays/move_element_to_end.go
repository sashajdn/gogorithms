package arrays

import (
	"sort"
)

// MoveElementToEnd T -> O(nlogn), S -> O(1)
func MoveElementToEnd(array []int, toMove int) []int {
	sort.Ints(array)

	lidx := bisect(array, toMove)
	ridx := lidx

	if lidx == -1 {
		return array

	}

	for array[ridx] == array[lidx] {
		ridx++
	}

	for array[lidx] == array[ridx-1] {
		lidx--
	}

	r := append(array[ridx:], array[lidx+1:ridx]...)
	r = append(array[:lidx+1], r...)

	return r
}

// bisect binary search, returns the index at which the first value
// found in the array
func bisect(arr []int, toFind int) int {

	midx := len(arr) / 2

	for midx >= 0 && midx < len(arr) {
		if arr[midx] == toFind {
			return midx
		}

		if arr[midx] < toFind {
			midx = midx / 2
		} else {
			midx += midx / 2
		}
	}

	return -1
}
