package arrays

// KRadiusSubarrayAverage ...
//
// T -> O(n)
// S -> O(1)
func KRadiusSubarrayAverage(nums []int, k int) []int {
	if k == 0 {
		return nums
	}

	var output = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		output[i] = -1
	}

	var sum int
	for j := 0; j < len(nums); j++ {
		sum += nums[j]

		if j >= (2*k)+1 {
			sum -= nums[j-((2*k)+1)]
		}

		if j >= (2 * k) {
			output[j-k] = sum / ((2 * k) + 1)
		}
	}

	return output
}
