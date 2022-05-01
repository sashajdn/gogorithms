package arrays

import (
	"math"
)

// SlidingWindowMaximum_BruteForce ...
//
// T -> O(nk)
// S -> O(k)
func SlidingWindowMaximum_BruteForce(nums []int, k int) []int {
	if len(nums)*k == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}

	var (
		output        = make([]int, 0, len(nums)-k+1)
		slidingWindow = make([]int, 0, k)
	)

	var maxInWindow = math.MinInt
	for i := 0; i < k; i++ {
		slidingWindow = append(slidingWindow, nums[i])
		if nums[i] > maxInWindow {
			maxInWindow = nums[i]
		}
	}
	output = append(output, maxInWindow)

	// T -> O(n * k)
	// S -> O(k)
	for j := k; j < len(nums); j++ {
		var newMax = math.MinInt
		for i := 1; i < k; i++ {
			newMax = max(newMax, slidingWindow[i])
			slidingWindow[i-1], slidingWindow[i] = slidingWindow[i], slidingWindow[i-1]
		}

		newMax = max(newMax, nums[j])
		slidingWindow[k-1] = nums[j]

		output = append(output, newMax)
	}

	return output
}
