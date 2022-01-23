package dynamic

// NumberOfWaysToMakeChange given an array of distinct positive integers representing coin
// denominations and a single non-negative integer `n` representing a target amount, returns
// the number of possible ways to make change for that target amount using the denominations.
// Assumes unlimited coins.
//
// T -> O(nd) where n is the target amount & d is the number of denominations.
// S -> O(nd) where n is the target amount & d is the number of denominations.
func NumberOfWaysToMakeChange(n int, denoms []int) int {
	var ways = make([]int, n+1)
	ways[0] = 1

	for _, d := range denoms {
		for i := 1; i < n+1; i++ {
			if d > i {
				continue
			}

			ways[i] += ways[i-d]
		}
	}

	return ways[n]
}
