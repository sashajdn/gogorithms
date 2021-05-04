package arrays

// SortedSquareArray T(N) -> O(n), S(N) -> O(1)
func SortedSquareArray(array []int) []int {
	if array == nil {
		return nil
	}
	l, r := 0, len(array)-1
	sortedSquares := make([]int, len(array))
	for i := len(array) - 1; i >= 0; i-- {
		if abs(array[l]) > abs(array[r]) {
			sortedSquares[i] = array[l] * array[l]
			l++
		} else {
			sortedSquares[i] = array[r] * array[r]
			r--
		}
	}
	return sortedSquares
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
