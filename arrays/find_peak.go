package arrays

// FindPeakElement_Iterative ...
//
// T -> O(log(n))
// S -> O(1)
func FindPeakElement_Iterative(nums []int) int {
	var left, right = 0, len(nums) - 1
	for left < right {
		mid := (left + right) / 2

		if mid+1 > len(nums)-1 {
			return len(nums) - 1
		}

		if nums[mid+1] >= nums[mid] {
			left = mid + 1
			continue
		}

		right = mid
	}

	return left
}

// FindPeakElement_Recursive ...
//
// T -> O(log(n))
// S -> O(log(n))
func FindPeakElement_Recursive(nums []int) int {
	return searchPeakElement(nums, 0, len(nums)+1)
}

func searchPeakElement(nums []int, l, r int) int {
	if l == r {
		return l
	}

	mid := (l + r) / 2
	if nums[mid] > nums[mid+1] {
		return searchPeakElement(nums, l, mid)
	}

	return searchPeakElement(nums, mid+1, r)
}
