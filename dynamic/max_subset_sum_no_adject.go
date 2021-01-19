package dynamic

// MaxSubsetNoAdjacent O(t) -> O(n), O(s) -> O(1)
func MaxSubsetNoAdjacent(array []int) int {
	switch len(array) {
	case 0:
		return 0
	case 1:
		return array[0]
	}

	var current int
	first := max(array[0], array[1])
	second := array[0]

	for i := 2; i < len(array); i++ {
		current = max(first, second+array[i])
		second = first
		first = current
	}
	return first
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
