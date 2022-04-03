package binarytree

// MinimizedMaximum ...
//
// T -> O(nlog(max(n)))
// S -> O(1)
func MinimizedMaximum(quantities []int, numberOfStores int) int {
	var max int
	for _, q := range quantities {
		if q > max {
			max = q
		}
	}

	var left, right = 1, max
	for left < right {
		mid := (left + right) / 2

		if isDistributionFeasible(quantities, mid, numberOfStores) {
			right = mid
			continue
		}

		left = mid + 1
	}

	return left
}

func isDistributionFeasible(quantities []int, threshold, numberOfStores int) bool {
	var currentCount int
	for _, q := range quantities {
		currentCount += q / threshold

		if q%threshold > 0 {
			currentCount++
		}
	}

	if currentCount > numberOfStores {
		return false
	}

	return true
}
