package arrays

func findPeakElement(nums []int) int {
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
