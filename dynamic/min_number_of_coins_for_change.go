package dynamic

import (
	"math"
)

// MinNumberOfCoinsForChange given an array of positive integers representing coin denoms
// and a single non-negative integer representing a target, returns the smallest number of coins
// needed to make change for that target.
// Assumes infinte amount of coins per denomination.
// Returns -1 if impossible to make change.
//
// S -> O(...)
// T -> O(...)
func MinNumberOfCoinsForChange(n int, denoms []int) int {
	// Initialize.
	var ways = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		ways[i] = math.MaxInt32
	}
	ways[0] = 0 // since zero coins required to make zero change.

	// Calculate the minimum number of ways using dynamic programming.
	for _, d := range denoms {
		for i := 1; i < n+1; i++ {
			if d <= i {
				ways[i] = min(ways[i], ways[i-d]+1)
			}
		}
	}

	// Validation that we can actually make change.
	switch o := ways[n]; {
	case o == math.MaxInt32:
		return -1
	default:
		return o
	}
}
