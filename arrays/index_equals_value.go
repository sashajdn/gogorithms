package arrays

func IndexEqualsValue(array []int) int {
	if len(array) == 0 {
		return -1
	}
	return searchIndexEqualsValue(array, 0, len(array))
}

func searchIndexEqualsValue(array []int, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2
	if array[mid] == mid && mid == 0 {
		return mid
	}

	if array[mid] == mid && array[mid-1] < mid-1 {
		return mid
	}

	if array[mid] < mid {
		return searchIndexEqualsValue(array, mid+1, right)
	}

	return searchIndexEqualsValue(array, left, mid-1)
}
