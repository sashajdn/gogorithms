package greedy

import "sort"

// Knapsack ...
//
// T -> O(nlog(n))
// S -> O(1)
func Knapsack(items [][]float64, capacity float64) (float64, [][]float64) {
	sort.Slice(items, func(i, j int) bool {
		a, b := items[i], items[j]

		return a[0]/a[1] > b[0]/b[1]
	})

	var (
		capacityLeft  = capacity
		includedItems [][]float64
		totalProfit   float64
	)
	for _, item := range items {
		if capacityLeft <= 0 {
			break
		}

		profit, weight := item[0], item[1]
		if weight <= capacityLeft {
			includedItems = append(includedItems, item)
			totalProfit += profit
			capacityLeft -= weight
			continue
		}

		ratio := profit / weight
		diff := capacityLeft - weight
		toTake := weight - (diff / ratio)

		totalProfit += (toTake * ratio)
		capacityLeft -= toTake
		includedItems = append(includedItems, []float64{toTake * ratio, toTake})
	}

	return totalProfit, includedItems
}
