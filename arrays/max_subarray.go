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

// MaxSubArray_Recursive ...
//
// T -> O(nlog(n)) where `n` is the number of elements in the input array.
// S -> O(log(n)) since we in the worst case will have log(n) calls in the stack as we are recursing.
func MaxSubArray_Recursive(nums []int) int {
	return maxSumRecurse(nums, 0, len(nums)-1)
}

func maxSumRecurse(nums []int, l, r int) int {
	if l > r {
		return math.MinInt
	}

	var (
		mid                                = (l + r) / 2
		current, bestLeftSum, bestRightSum int
	)

	for i := mid - 1; i >= l; i-- {
		current += nums[i]
		bestLeftSum = max(bestLeftSum, current)
	}

	current = 0
	for j := mid + 1; j <= r; j++ {
		current += nums[j]
		bestRightSum = max(bestRightSum, current)
	}

	combination := nums[mid] + bestLeftSum + bestRightSum

	return max(combination, max(maxSumRecurse(nums, l, mid-1), maxSumRecurse(nums, mid+1, r)))
}
