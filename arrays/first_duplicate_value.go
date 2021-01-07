package arrays

// Find the first duplicate number for an array of length n,
// with {1, n} inclusive integers.

// FirstDuplicateValue finds the first duplicate value in the given array
// O(t) -> O(n) O(s) -> O(1)
func FirstDuplicateValue(array []int) int {
	for _, value := range array {
		absValue := func(a int) int {
			if a < 0 {
				return -a
			}
			return a
		}(value)
		if array[absValue-1] < 0 {
			return absValue
		}
		array[absValue-1] *= -1
	}
	return -1
}
