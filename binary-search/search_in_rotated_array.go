package binarysearch

// SearchInRotatedArray ...
//
// T -> O(log(n)) where `n` is the number of elements in the input numbers.
// S -> O(1)
func SearchInRotatedArray(numbers []int, target int) int {
	var left, right = 0, len(numbers) - 1
	for left <= right {
		mid := left + (right-left)/2

		if target == numbers[mid] {
			return mid
		}

		if numbers[mid] >= numbers[left] {
			if target >= numbers[left] && target < numbers[mid] {
				right = mid - 1
				continue
			}

			left = mid + 1
			continue
		}

		if target <= numbers[right] && target > numbers[mid] {
			if target <= numbers[right] && target > numbers[mid] {
				left = mid + 1
				continue
			}

			right = mid - 1
		}
	}

	return -1
}
