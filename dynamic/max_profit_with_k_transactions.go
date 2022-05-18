package dynamic

import (
	"math"
)

// MaxProfitWithKTransactions_BottomUp ...
//
// T -> O(k * (p**p))
// S -> O(k * (p**p))
func MaxProfitWithKTransactions_BottomUp(prices []int, k int) int {
	type tuple [3]int
	var memo = map[tuple]int{}

	var dfs func(l, r, k int) int
	dfs = func(l, r, k int) int {
		if l >= r || k == 0 {
			return 0
		}

		var t = tuple{l, r, k}
		if profit, ok := memo[t]; ok {
			return profit
		}

		var maxSoFar = math.MinInt
		for i := l; i < r; i++ {
			for j := i + 1; j <= r; j++ {
				currentTransationProfit := prices[j] - prices[i]
				profit := max(currentTransationProfit+dfs(j+1, r, k-1), currentTransationProfit)

				maxSoFar = max(maxSoFar, profit)
			}
		}

		memo[t] = maxSoFar
		return maxSoFar
	}

	return dfs(0, len(prices)-1, k)
}

// MaxProfitWithKTransactions_TopDown2D ...
//
// T -> O(n * k) where `n` is the number of prices & `k` is the number of transactions.
// S -> O(n * k)
func MaxProfitWithKTransactions_TopDown2D(prices []int, k int) int {
	if len(prices) == 0 {
		return 0
	}

	var dp = make([][]int, 0, k+1)
	for i := 0; i < k+1; i++ {
		dp = append(dp, make([]int, len(prices)))
	}

	for j := 1; j < k+1; j++ {
		var maxSoFar = math.MinInt
		for i := 1; i < len(prices); i++ {
			maxSoFar = max(maxSoFar, dp[j-1][i-1]-prices[i-1])
			dp[j][i] = max(dp[j][i-1], maxSoFar+prices[i])
		}
	}

	return dp[k][len(prices)-1]
}
