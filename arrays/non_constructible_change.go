package arrays

import "sort"

func NonConstructibleChange(coins []int) int {
	if len(coins) == 0 {
		return 1
	}
	sort.Ints(coins)

	var currentChangeSoFar int
	for _, coin := range coins {
		if coin > currentChangeSoFar+1 {
			return currentChangeSoFar + 1
		}
		currentChangeSoFar += coin
	}
	return currentChangeSoFar + 1
}
