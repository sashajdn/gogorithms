package dynamic

import "math"

// HouseRobber_Optimized ...
//
// T -> O(n) where `n` is the number of houses.
// S -> O(1)
func HouseRobber_Optimized(houses []int) int {
	var (
		maxSoFar                             = math.MinInt
		adjacent, previous, previousPrevious int
	)
	for i := 0; i < len(houses); i++ {
		var current int
		switch {
		case i-3 >= 0:
			current = max(previous, previousPrevious) + houses[i]
		case i-2 >= 0:
			current = previous + houses[i]
		default:
			current = houses[i]
		}

		maxSoFar = max(maxSoFar, current)
		adjacent, previous, previousPrevious = current, adjacent, previous
	}

	return maxSoFar
}

// HouseRobber ...
//
// T -> O(n) where `n` is the number of houses.
// S -> S(n)
func HouseRobber(houses []int) int {
	var (
		dp       = make([]int, len(houses))
		maxSoFar int
	)
	for i := 0; i < len(houses); i++ {
		switch {
		case i-3 >= 0:
			dp[i] = max(dp[i-3], dp[i-2]) + houses[i]
		case i-2 >= 0:
			dp[i] = dp[i-2] + houses[i]
		default:
			dp[i] = houses[i]
		}

		maxSoFar = max(maxSoFar, dp[i])
	}

	return maxSoFar
}
