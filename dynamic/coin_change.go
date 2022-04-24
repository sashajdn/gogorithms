package dynamic

// CoinChange ...
//
// T -> O(c * t) where `c` is the number of coins & `t` is the target amount.
// S -> O(t)
func CoinChange(coins []int, target int) int {
	var dp = make([]int, target+1)
	for i := 1; i < target+1; i++ {
		dp[i] = target + 1
	}

	for _, coin := range coins {
		for i := 1; i < target+1; i++ {
			if coin > i {
				continue
			}

			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	return dp[target]
}
