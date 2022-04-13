package binarysearch

import "math"

// MedianSortedArrays ...
//
// T -> O(log(n + m))
// S -> O(1)
func MedianSortedArrays(smallest, largest []int) float64 {
	if len(smallest) > len(largest) {
		return MedianSortedArrays(largest, smallest)
	}

	total := len(smallest) + len(largest)
	half := total / 2

	var left, right = 0, len(smallest) - 1
	for {
		var (
			pSmallest = int(math.Floor(float64(left+right) / 2))
			pLargest  = half - pSmallest - 2
		)

		var leftSmallest, rightSmallest, leftLargest, rightLargest int
		switch {
		case pSmallest >= 0:
			leftSmallest = smallest[pSmallest]
		default:
			leftSmallest = math.MinInt
		}

		switch {
		case pSmallest < len(smallest)-1:
			rightSmallest = smallest[pSmallest+1]
		default:
			rightSmallest = math.MaxInt
		}

		switch {
		case pLargest >= 0:
			leftLargest = largest[pLargest]
		default:
			leftLargest = math.MinInt
		}

		switch {
		case pLargest < len(largest)-1:
			rightLargest = largest[pLargest+1]
		default:
			rightLargest = math.MaxInt
		}

		if leftSmallest <= rightLargest && leftLargest <= rightSmallest {
			if total%2 == 0 {
				return float64(max(leftSmallest, leftLargest)+min(rightSmallest, rightLargest)) / 2.0
			}

			return float64(min(rightSmallest, rightLargest))
		}

		if leftSmallest < rightLargest {
			right = pSmallest - 1
			continue
		}

		left = pSmallest + 1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
