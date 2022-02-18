package arrays

func binarySearch(array []int, k int) bool {
	if len(array) == 0 {
		return false
	}

	mid := len(array) / 2

	if k > array[mid] {
		return binarySearch(array[mid+1:], k)
	}

	if k < array[mid] {
		return binarySearch(array[:mid], k)
	}

	return true
}
