package searching

// Quickselect ...
//
// T -> Best: O(n), Worst: O(n ** 2), Avg: O(n), where n is the number of elements in the array.
//      The worst case scenario is when the the pivot divides the array into sub-arrays of length = 1 & length = length(array)-1
//      for every pivot selection.
// S -> O(1)
func Quickselect(array []int, k int) int {
	if k >= len(array) {
		return -1
	}

	quickselect(array, 0, len(array)-1, k-1)

	return array[k-1]
}

func quickselect(array []int, start, end int, k int) {
	if start > end {
		return
	}

	pivot := start
	l, r := start+1, end

	for l <= r {
		if array[l] > array[pivot] && array[r] < array[pivot] {
			quickselectSwap(array, l, r)
		}
		if array[l] <= array[pivot] {
			l++
		}
		if array[r] >= array[pivot] {
			r--
		}
	}
	quickselectSwap(array, pivot, r)

	if r == k {
		return
	}

	if r < k {
		quickselect(array, r+1, end, k)
		return
	}

	quickselect(array, start, r-1, k)
	return
}

func quickselectSwap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
