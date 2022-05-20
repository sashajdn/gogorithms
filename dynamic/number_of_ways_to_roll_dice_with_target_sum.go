package dynamic

import (
	"math"
)

// NumberOfWaysToRollDice_BottomUp ...
//
// T -> O(f ^ d) -> with memo O(f * d * t)
// S -> O() -> with memo O(f * d)
func NumberOfWaysToRollDice_BottomUp(d, f, target int) int {
	type tuple [2]int
	var memo = map[tuple]int{}

	var dfs func(rolls, targetLeft int) int
	dfs = func(rolls, targetLeft int) int {
		switch {
		case targetLeft == 0:
			return 1
		case targetLeft < 0:
			return 0
		case rolls == 0:
			return 0
		}

		var t = tuple{rolls, targetLeft}
		if howMany, ok := memo[t]; ok {
			return howMany
		}

		var numberOfWays int
		for i := 1; i <= f; i++ {
			numberOfWays += dfs(rolls-1, targetLeft-i)
		}

		memo[t] = numberOfWays
		return numberOfWays
	}

	return dfs(d, target) % (int(math.Pow(10, 9)) + 7)
}

// NumberOfWaysToRollDice_TopDown ...
//
// T -> O(d * f * target)
// S -> O(target)
func NumberOfWaysToRollDice_TopDown(d, f, target int) int {
	var dp = make([]int, target+1)
	dp[0] = 1

	var m = int(math.Pow(10, 9)) + 7
	for i := 0; i < d; i++ {
		var rollDp = make([]int, target+1)
		for k := 0; k <= target; k++ {
			for j := 1; j <= f; j++ {
				if j+k > target {
					continue
				}

				rollDp[k+j] = (rollDp[k+j] + dp[k]) % m
			}
		}

		dp = rollDp
	}

	return dp[target]
}
