package arrays

import "sort"

// MinimumWaitingTime O(T) -> O(nlogn), O(S) -> O(1)
func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)
	var waitingTime int
	for i, q := range queries {
		queriesLeft := len(queries) - (i + 1)
		waitingTime += q * queriesLeft
	}
	return waitingTime
}
