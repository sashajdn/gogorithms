package arrays

// SumOfSubArraysBruteForce ...
//
// T -> O(n ** 3) where `n` is the number of numbers. Where we have n**2 to create all of the sub arrays;
//                and for each subarray, we must add up all of the elements in the subarray (n + (n-1) + (n-2) + ... + (n - (n - 1))
//                which tends to O(n): thus in total this gives us O(n ** 3)
// S -> O(1) - linear space complexity since we don't have to createa any auxiliary space.

func SumOfSubArraysBruteForce(nums []int, k int) int {
	var count int
	for start := 0; start < len(nums); start++ {
		for end := start + 1; end <= len(nums); end++ {
			var sum int
			for j := start; j < end; j++ {
				sum += nums[j]
			}

			if sum == k {
				count++
			}
		}
	}
	return count
}

// SumOfSubArraysCumSum ...
//
// T -> O(n ** 2)
// S -> O(n)
func SumOfSubArraysCumSum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	var cumSum = make([]int, len(nums)+1)
	cumSum[0] = 0
	for i := 1; i < len(nums)+1; i++ {
		cumSum[i] = cumSum[i-1] + nums[i-1]
	}

	var count int
	for start := 0; start < len(nums); start++ {
		for end := start + 1; end <= len(nums); end++ {
			if cumSum[end]-cumSum[start] == k {
				count++
			}
		}
	}
	return count
}

// SumOfSubArraysWithoutSpace ...
//
// T -> O(n ** 2)
// S -> O(1)
func SumOfSubArraysWithoutSpace(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	var count int
	for start := 0; start < len(nums); start++ {
		var sum int
		for end := start; end < len(nums); end++ {
			sum += nums[end]

			if sum == k {
				count++
			}
		}
	}
	return count
}

// SumOfSubArraysHashMap ...
//
// T -> O(n) where n is the number of numbers.
// S -> O(n) due to the auxilary hashmap we create.
func SumOfSubArraysHashMap(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		count, cumSum int
		hm            = map[int]int{
			0: 1,
		}
	)
	for _, v := range nums {
		cumSum += v
		if howMany, ok := hm[cumSum-k]; ok {
			count += howMany
		}

		hm[cumSum] += 1
	}

	return count
}
