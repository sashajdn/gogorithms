package dynamic

import (
	"math"
)

// NumberOfWaysToRollDice ...
//
// T -> O(d * f * target)
// S -> O(target)
func NumberOfWaysToRollDice(d, f, target int) int {
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
