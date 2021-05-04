package arrays

// SortedSquareArray T(N) -> O(n), S(N) -> O(1)
func SortedSquareArray(array []int) []int {
	if array == nil {
		return nil
	}
	for i := 0; i < len(array); i++ {
		array[i] *= array[i]
	}
	return array
}
