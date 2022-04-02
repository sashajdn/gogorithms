package binarytree

// CuttingRibbons ...
//
// T -> O(n log(max(n)))
// S -> O(1)
func CuttingRibbons(ribbons []int, k int) int {
	// T -> O(n)
	var max int
	for _, r := range ribbons {
		if r > max {
			max = r
		}
	}

	// T -> O(n * log(max(n)))
	// S -> O(1)
	left, right := 0, max
	for left < right {
		mid := (left + right + 1) / 2

		// -> T -> O(log(max(n)))
		if feasible(ribbons, mid, k) {
			left = mid
			continue
		}

		right = mid - 1
	}

	return right
}

func isRibbonsFeasible(ribbons []int, ribbonLength, minNumberOfRibbons int) bool {
	var count int
	for _, r := range ribbons {
		count += r / ribbonLength
	}

	if count >= minNumberOfRibbons {
		return true
	}

	return false
}
