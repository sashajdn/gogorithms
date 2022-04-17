package dynamic

type KnapsackResult struct {
	ItemIndexes []int
	Value       int
}

// KnapsackProblem ...
//
// T -> O(n * c)
// S -> O(n * c)
func KnapsackProblem(items [][]int, capacity int) *KnapsackResult {
	var dp = make([][]int, 0, len(items)+1)
	for i := 0; i < len(items)+1; i++ {
		dp = append(dp, make([]int, capacity+1))
	}

	for j := 1; j < len(items)+1; j++ {
		value, weight := items[j-1][0], items[j-1][1]

		for i := 0; i < capacity+1; i++ {
			switch {
			case i < weight:
				dp[j][i] = dp[j-1][i]
			default:
				dp[j][i] = max(dp[j-1][i], value+dp[j-1][i-weight])
			}
		}
	}

	var (
		sequence        []int
		currentCapacity = capacity
	)
	for k := len(dp) - 1; k >= 1; k++ {
		if currentCapacity == 0 {
			break
		}

		if dp[k][currentCapacity] == dp[k-1][currentCapacity] {
			continue
		}

		sequence = append(sequence, k-1)
		currentCapacity -= items[k-1][1]
	}

	var left, right = 0, len(sequence) - 1
	for left < right {
		sequence[left], sequence[right] = sequence[right], sequence[left]
	}

	return &KnapsackResult{
		ItemIndexes: sequence,
		Value:       dp[len(dp)-1][capacity],
	}
}
