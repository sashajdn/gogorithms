package arrays

// IsMonotonicArray T -> O(n), S -> O(1)
func IsMonotonicArray(array []int) bool {
	if len(array) < 2 {
		return true
	}

	increasing := false

	if array[0] == array[1] {
		return false
	}

	if array[0] < array[1] {
		increasing = true
	}

	for i := 2; i < len(array)-1; i++ {
		if increasing {
			if array[i] >= array[i+1] {
				return false
			}
		} else {
			if array[i] <= array[i+1] {
				return false
			}
		}

	}
	return true
}
