package arrays

// LongestPeak T -> O(n), S -> O(1)
func LongestPeak(array []int) int {
	if len(array) < 3 {
		return 0
	}

	maxSoFar, curr := 0, 0
	increasing := true

	for i := 0; i < len(array)-1; i++ {
		if increasing {
			if array[i] < array[i+1] {
				curr++
				continue
			} else if array[i] > array[i+1] && curr != 0 {
				increasing = false
				curr++
			} else {
				curr = 0
				continue
			}
		} else {
			if array[i] < array[i+1] {
				increasing = true
				maxSoFar = max(maxSoFar, curr)
				curr = 1
			} else if array[i] > array[i+1] {
				curr++
				continue
			} else {
				maxSoFar = max(maxSoFar, curr)
				curr = 0
				continue
			}
		}

	}

	if maxSoFar == 0 {
		return maxSoFar
	}

	// since the iteration counts steps, not items - thus items = steps + 1
	return maxSoFar + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
