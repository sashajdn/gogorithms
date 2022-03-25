package dynamic

// MinNumberOfWaysToMakeChange ...
//
// T -> O(len(coins) * amount)
// S -> O(amount)
func MinNumberOfWaysToMakeChange(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}

	for j := 0; j < len(coins); j++ {
		coin := coins[j]
		for i := 1; i < amount+1; i++ {
			if coin > i {
				continue
			}

			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
