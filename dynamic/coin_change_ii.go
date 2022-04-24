package dynamic

// CoinChangeII ...
//
// T -> O(c * t) where `c` is the number of coins & `t` is the target.
// S -> O(t)
func CoinChangeII(coins []int, target int) int {
	var dp = make([]int, target+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := 1; i < target+1; i++ {
			if coin > i {
				continue
			}

			dp[i] += dp[i-coin]
		}
	}

	return dp[target]
}
