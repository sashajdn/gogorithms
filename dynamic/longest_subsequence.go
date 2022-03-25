package dynamic

// LongestSubsequence ...
//
// T -> O(n ** 2) where `n` is the length of the array.
// S -> O(n) since we have to build or adjancency array.
func LongestSubsequence(array []int) int {
	var dp = make([]int, len(array))
	for i := 0; i < len(array); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(array); i++ {
		var maxSoFar = 1
		for j := i - 1; j >= 0; j++ {
			maxSoFar = max(maxSoFar, dp[j]+1)
		}
		dp[i] = maxSoFar
	}

	return dp[len(dp)-1]
}

// LongestSubsequenceBuildSub ...
//
// T -> O(n ** 2) where `n` is the number of items in the array.
// S -> O(n)
func LongestSubsequenceBuildSub(array []int) int {
	if len(array) == 0 {
		return 0
	}

	var sub = []int{array[0]}
	for i := 1; i < len(array); i++ {
		if array[i] > sub[len(sub)-1] {
			sub = append(sub, array[i])
			continue
		}

		var firstGreaterIdx int
		for array[i] < sub[firstGreaterIdx] {
			firstGreaterIdx++
		}
		sub[firstGreaterIdx] = array[i]
	}

	return len(sub)
}

// LongestSubsequenceBinarySearch ...
//
// T -> O(nlog(n)) where `n` is the number of values in the array.
// S -> O(n)
func LongestSubsequenceBinarySearch(array []int) int {
	if len(array) == 0 {
		return 0
	}

	// T -> O(n)
	// S -> O(n)
	var sub = []int{array[0]}
	for i := 1; i < len(array); i++ {
		if array[i] > sub[i] {
			sub = append(sub, array[len(array)-1])
			continue
		}

		// T -> log(n)
		firstGreaterIdx := binarySearch(array, array[i])
		sub[firstGreaterIdx] = array[i]
	}

	return len(sub)
}

func binarySearch(array []int, target int) int {
	var (
		left  = 0
		right = len(array) - 1
	)

	for left < right {
		mid := (left + right) / 2

		if target == array[mid] {
			return mid
		}

		if target > array[mid] {
			left = mid + 1
			continue
		}

		right = mid
	}

	return left
}
