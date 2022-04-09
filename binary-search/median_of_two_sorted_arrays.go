package binarysearch

import "math"

// MedianOfTwoSortedArrays ...
func MedianOfTwoSortedArrays(a, b []int) int {
	if len(a) > len(b) {
		return MedianOfTwoSortedArrays(b, a)
	}

	total := len(a) + len(b)
	half := total / 2

	var left, right = 0, len(a)
	for {
		i := (left + right) / 2
		j := half - i - 2

		var leftA, rightA, leftB, rightB int
		switch {
		case i < 0:
			leftA = math.MinInt
		default:
			leftA = a[i]
		}

		switch {
		case i+1 >= len(a):
			rightA = math.MaxInt
		default:
			rightA = a[i+1]
		}

		switch {
		case j < 0:
			leftB = math.MinInt
		default:
			leftB = b[j]
		}

		switch {
		case j+1 >= len(b):
			rightB = math.MaxInt
		default:
			rightB = b[j+1]
		}

		if leftA <= rightB && rightA <= leftB {
			if total%2 == 0 {
				return max(rightA, rightB)
			}

			return min(leftA, leftB) + max(rightA, rightB)
		}

		if leftA > rightB {
			right = i - 1
			continue
		}

		left = i + 1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
