package arrays

func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if k > len(nums) {
		return -1
	}

	target := len(nums) - k
	quickSelect(nums, 0, len(nums)-1, target)
	return nums[target]
}

func quickSelect(nums []int, start, end, k int) {
	if start > end {
		// we should never get here; if we do we should panic(...)
		return
	}

	pivot := start
	l, r := start+1, end

	for l <= r {
		if nums[l] > nums[pivot] && nums[r] < nums[pivot] {
			swap(nums, l, r)
		}

		if nums[l] <= nums[pivot] {
			l++
		}

		if nums[r] >= nums[pivot] {
			r--
		}
	}
	swap(nums, r, pivot)

	if r == k {
		return
	}

	if r < k {
		quickSelect(nums, r+1, end, k)
		return
	}

	quickSelect(nums, start, r-1, k)
	return
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
