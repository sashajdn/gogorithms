package arrays

// FindMinimumSortedRotatedArray ...
//
// T -> O(log(n)) where `n` is the number of values in the array.
// S -> O(1)
func FindMinimumSortedRotatedArray(array []int) int {
	if len(array) == 0 {
		return 0
	}
	if len(array) == 1 {
		return array[0]
	}

	if array[len(array)-1] > array[0] {
		return array[0]
	}

	// Binary search.
	// T -> O(log(n))
	l, r := 0, len(array)-1
	for l < r {
		m := l + (r-l)/2

		if array[m+1] < array[m] {
			return array[m+1]
		}

		if array[m-1] > array[m] {
			return array[m]
		}

		if array[m] > array[0] {
			l = m + 1
			continue
		}

		r = m
	}

	return -1
}
