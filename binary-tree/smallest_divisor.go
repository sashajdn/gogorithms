package binarytree

import "math"

// FindSmallestDivisor ...
//
// T -> O(nlog(max(n))) where `n` is the number of items in array, max(n) is the max number in the numbers.
// S -> O(1)
func FindSmallestDivisor(numbers []int, threshold int) int {
	if len(numbers) == 0 {
		return 0
	}

	var max int
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	var (
		l, r = 1, max
	)
	for l < r {
		m := l + (r-l)/2

		var sum int
		for _, number := range numbers {
			sum += int(math.Ceil(float64(number) / float64(m)))
		}

		if sum <= threshold {
			r = m
			continue
		}
		l = m + 1
	}

	return l
}
