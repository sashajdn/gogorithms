package dynamic

import (
	"math"
)

// JumpGameII_Dynamic ...
//
// T -> O(n ** 2)
// S -> O(n)
func JumpGameII_Dynamic(nums []int) int {
	var dp = make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		dp[i] = math.MaxInt
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < min(len(nums), i+nums[i]+1); j++ {
			dp[j] = min(dp[i]+1, dp[j])
		}
	}

	return dp[len(dp)-1]
}

// JumpGameII_Greedy ...
//
// T -> O(n)
// S -> O(1)
func JumpGameII_Greedy(nums []int) int {
	var (
		jumps, maxReach, lastIndex int
	)
	for i := 0; i < len(nums)-1; i++ {
		maxReach = max(maxReach, i+nums[i])

		if i >= lastIndex {
			jumps++
			lastIndex = maxReach
		}
	}

	return jumps
}
