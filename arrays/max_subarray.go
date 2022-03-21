package arrays

import "math"

// MaxSubArray ...
//
// T -> O(n) where `n` is the number of elements in the array.
// S -> O(1)
func MaxSubArray(nums []int) int {
	var (
		currentSubArraySum int
		maxSubArraySum     = math.MinInt
	)

	for _, num := range nums {
		maxSubArraySum = max(maxSubArraySum, num)

		currentSubArraySum += num
		if currentSubArraySum < 0 {
			currentSubArraySum = 0
			continue
		}

		maxSubArraySum = max(maxSubArraySum, currentSubArraySum)
	}

	return maxSubArraySum
}
