package binarysearch

import (
	"fmt"
	"math"
)

// MaximumCandies ...
//
// T -> O(nlog(n))
// S -> O(1)
func MaximumCandies(candies []int, k int64) int {
	var min, max = 0, math.MinInt
	for _, c := range candies {
		if c > max {
			max = c
		}
	}

	var left, right = min, max
	for left < right {
		mid := (left + right + 1) / 2

		if maximumCandiesFeasible(candies, mid, int(k)) {
			left = mid
			continue
		}

		right = mid - 1
	}

	fmt.Println(left, right)

	return right
}

func maximumCandiesFeasible(candies []int, threshold int, k int) bool {
	var count int

	for _, c := range candies {
		count += (c / threshold)
	}

	if count < k {
		return false
	}

	return true
}
