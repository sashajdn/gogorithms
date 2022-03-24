package arrays

// MaximumProduct ...
//
// T -> O(n)
// S -> O(1)
func MaximumProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		result   = nums[0]
		maxSoFar = nums[0]
		minSoFar = nums[0]
	)

	for _, n := range nums[1:] {
		if n == 0 {
			result = max(result, 0)
			minSoFar, maxSoFar = 0, 0
		}

		tempMax := max(maxSoFar, max(maxSoFar*n, minSoFar*n))
		minSoFar = min(minSoFar, min(minSoFar*n, maxSoFar*n))
		maxSoFar = tempMax

		result = max(result, maxSoFar)
	}

	return result
}
