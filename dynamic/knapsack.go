package dynamic

// S (len(items) * capacity )
func Knapsack(items [][]int, capacity int) (int, [][]int) {
	var dp = make([][]int, 0, len(items)+1)
	for i := 0; i < len(items)+1; i++ {
		dp = append(dp, make([]int, capacity+1))
	}

	for j := 1; j < len(items)+1; j++ {
		value, weight := items[j-1][0], items[j-1][1]
		for i := 1; i < capacity+1; i++ {
			switch {
			case i < weight:
				dp[j][i] = dp[j-1][i]
			default:
				dp[j][i] = max(dp[j-1][i], dp[j-1][i-weight]+value)
			}
		}
	}

	var (
		includedItems   [][]int
		currentCapacity = capacity
	)
	for j := len(items); j >= 1; j-- {
		if currentCapacity == 0 {
			break
		}

		if dp[j][currentCapacity] == dp[j-1][currentCapacity] {
			continue
		}

		includedItems = append(includedItems, items[j-1])
		currentCapacity -= items[j-1][1]
	}

	var left, right = 0, len(includedItems) - 1
	for left < right {
		includedItems[left], includedItems[right] = includedItems[right], includedItems[left]
		left++
		right--
	}

	return dp[len(dp)-1][capacity], includedItems
}
