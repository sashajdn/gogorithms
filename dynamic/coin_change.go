package dynamic

// CoinChange_Recursive ...
//
// T -> O(C ** A) where `C` is the number of coins and `A` is the amount.
// S -> O(A) where in the worst case we'll have `A` recursive calls on the stack at once.
func CoinChange_BottomUp(coins []int, amount int) int {
	var memo = map[int]int{}
	var search func(amountLeft int) int
	search = func(amountLeft int) int {
		if amountLeft == 0 {
			return 0
		}
		if amountLeft < 0 {
			return -1
		}
		if v, ok := memo[amountLeft]; ok {
			return v
		}

		for _, coin := range coins {
			if coin > amountLeft {
				continue
			}

			remaining := amountLeft - coin
			if v := search(remaining); v >= 0 {
				if _, ok := memo[amountLeft]; !ok {
					memo[amountLeft] = v + 1
					continue
				}

				memo[amountLeft] = min(memo[amountLeft], v+1)
			}
		}

		if v, ok := memo[amountLeft]; ok {
			return v
		}

		return -1
	}

	return search(amount)
}

// CoinChange_TopDown2D ...
//
// T -> O(c * a)
// S -> O(c * a)
func CoinChange_TopDown2D(coins []int, amount int) int {
	var dp = make([][]int, 0, len(coins)+1)
	for j := 0; j < len(coins)+1; j++ {
		row := make([]int, amount+1)
		for j := 1; j < amount+1; j++ {
			row[j] = amount + 1
		}

		dp = append(dp, row)
	}

	for j := 1; j < len(coins)+1; j++ {
		coin := coins[j-1]
		for i := 1; i < amount+1; i++ {
			if coin > i {
				dp[j][i] = dp[j-1][i]
				continue
			}

			dp[j][i] = min(dp[j-1][i-coin], dp[j][i-coin]) + 1
		}
	}

	if dp[len(coins)][amount] > amount {
		return -1
	}

	return dp[len(coins)][amount]
}

// CoinChange_TopDown1D ...
//
// T -> O(c * a)
// S -> O(a)
func CoinChange_TopDown1D(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}

	for _, coin := range coins {
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
