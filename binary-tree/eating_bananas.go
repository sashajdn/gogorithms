package binarytree

// EatingBananas ...
//
// T -> O(nlog(sum(n))
// S -> O(1)
func EatingBananas(piles []int, h int) int {
	// T -> O(n)
	var sum int
	for _, p := range piles {
		sum += p
	}

	// T -> O(nlog(sum(n)))
	// S -> O(1)
	var left, right int
	for left < right {
		mid := (left + right) / 2

		// T -> O(n)
		if isSpeedFeasible(piles, mid, h) {
			right = mid
			continue
		}

		left = mid + 1
	}

	return left
}

func isSpeedFeasible(piles []int, threshold, h int) bool {
	var currentTimeSpent int
	for _, p := range piles {
		if p < threshold {
			currentTimeSpent++
			continue
		}

		currentTimeSpent += (p / threshold)
		if remainder := p % threshold; remainder > 0 {
			currentTimeSpent++
		}
	}

	if currentTimeSpent > h {
		return false
	}

	return true
}
