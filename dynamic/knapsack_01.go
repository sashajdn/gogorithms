package dynamic

import (
	"fmt"
	"strings"
)

// KnapsackO1_BottomUp ...
//
// T -> O(n ** 2)
// S -> O(n)
func KnapsackO1_BottomUp(items [][]int, capacity int) ([][]int, int) {
	type result struct {
		Profit int
		Items  [][]int
	}

	var memo = map[string]*result{}
	var dfs func(currentItems [][]int, capacityLeft int) ([][]int, int)
	dfs = func(currentItems [][]int, capacityLeft int) ([][]int, int) {
		if capacityLeft == 0 {
			return [][]int{}, 0
		}

		var key = hashItems(currentItems)
		if r, ok := memo[key]; ok {
			return r.Items, r.Profit
		}

		var (
			maxProfitSoFar int
			includedItems  = [][]int{}
		)
		for i, item := range currentItems {
			itemProfit, weight := item[0], item[1]
			if weight > capacityLeft {
				continue
			}

			newCurrentItems := append(append([][]int{}, currentItems[:i]...), currentItems[i+1:]...)
			itemSet, profit := dfs(newCurrentItems, capacityLeft-weight)

			if itemProfit+profit > maxProfitSoFar {
				includedItems = append(itemSet, item)
				maxProfitSoFar = itemProfit + profit
			}
		}

		memo[key] = &result{
			Profit: maxProfitSoFar,
			Items:  includedItems,
		}

		return includedItems, maxProfitSoFar
	}

	includedItems, capacity := dfs(items, capacity)

	var left, right = 0, len(includedItems) - 1
	for left < right {
		includedItems[left], includedItems[right] = includedItems[right], includedItems[left]
		left++
		right--
	}

	return includedItems, capacity
}

// KnapsackO1_TopDown ...
// T -> O(I * C) where `I` is the number of items, `C` is the capacity.
// S -> O(I * C)
func KnapsackO1_TopDown(items [][]int, capacity int) ([][]int, int) {
	var dp = make([][]int, 0, len(items)+1)
	for j := 0; j < len(items)+1; j++ {
		dp = append(dp, make([]int, capacity+1))
	}

	for j := 1; j < len(items)+1; j++ {
		value, weight := items[j-1][0], items[j-1][1]

		for i := 0; i < capacity+1; i++ {
			if weight > i {
				dp[j][i] = dp[j-1][i]
				continue
			}

			dp[j][i] = max(dp[j-1][i], value+dp[j-1][i-weight])
		}
	}

	var (
		includedItems   = [][]int{}
		currentCapacity = capacity
	)
	for k := len(items); k >= 1; k-- {
		if currentCapacity <= 0 {
			break
		}

		if dp[k][currentCapacity] == dp[k-1][currentCapacity] {
			continue
		}

		item := items[k-1]
		includedItems = append(includedItems, item)
		currentCapacity -= item[1]
	}

	var left, right = 0, len(includedItems) - 1
	for left < right {
		includedItems[left], includedItems[right] = includedItems[right], includedItems[left]
		left++
		right--
	}

	return includedItems, dp[len(items)][capacity]
}

func hashItems(items [][]int) string {
	var sb strings.Builder
	for _, item := range items {
		sb.WriteString(fmt.Sprintf("%d#%d*", item[0], item[1]))
	}

	return sb.String()
}
