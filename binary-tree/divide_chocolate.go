package binarytree

import "math"

// MaximiseSweetness ...
//
// T -> O(nlog(sum(n)/k)) where `n` is the number of items in the input array.
// S -> O(1)
func MaximiseSweetness(sweetness []int, k int) int {
	var (
		sum int
		min = math.MaxInt
	)

	// T -> O(n)
	for _, n := range sweetness {
		sum += n
		if n < min {
			min = n
		}
	}
	sum /= (k + 1)

	// O(n * log(sum(n)/k))
	var left, right = min, sum
	for left < right {
		mid := (left + right + 1) / 2

		// O(log(sum(n)/k))
		if sweetnessFeasible(sweetness, mid, k+1) {
			left = mid
			continue
		}

		right = mid - 1
	}

	return right
}

func sweetnessFeasible(sweetness []int, threshold, splitsRequired int) bool {
	var (
		currentSplits int
		total         int
	)
	for _, n := range sweetness {
		if total+n >= threshold {
			currentSplits++
			total = 0
			continue
		}

		total += n
	}

	if currentSplits < splitsRequired {
		return false
	}

	return true
}
