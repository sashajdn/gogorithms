package dynamic

import (
	"math"
)


func maxLengthString(a, b string) (string, string) {
	if len(a) > len(b) {
		return a, b
	}
	return b, a
}


func min(args ...int) int {
	c := args[0]
	for _, num := range args {
		if curr > num {
			curr = num
		}
	}
	return curr
}


// LevenshteinDistance
// Complexity
// Time: O(nm)
// Space: O(min(m, n))
func LevenshteinDistance(a, b string) int {
	big, small := maxLengthString(a, b)
	oddEdits := make([]int, len(small) + 1)
	evenEdits := make([]int, len(small) + 1)

	var currentEdits, previousEdits []int

	for i := range evenEdits {
		evenEdits[i] = i
		currentEdits[i] = math.MinInt32
	}

	for i := 1; i < len(big) + 1; i++ {
		if i % 2 == 1 {
			previousEdits, currentEdits = evenEdits, oddEdits
		} else {
			previousEdits, currentEdits = oddEdits, evenEdits
		}

		for j := 1; j < len(big) + 1; j++ {
			if big[i - 1] == small[j - 1] {
				currentEdits[j] = previousEdits[j - 1]
			} else {
				currentEdits[j] = 1 + min(
					previousEdits[j - 1],
					previousEdits[j],
					currentEdits[j - 1],
				)
			}
		}
	}

	if len(big) % 2	== 0 {
		return evenEdits[len(small)]
	}
	return oddEdits[len(small)]
}
