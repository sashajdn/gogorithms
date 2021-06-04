package arrays

import "sort"

// ClassPhotos : O(T) -> O(nlog(n)), O(S) -> O(1)
func ClassPhotos(redShirtHeights, blueShirtHeights []int) bool {
	return recurse(rowSort(redShirtHeights, blueShirtHeights))
}

func rowSort(a, b []int) (frontrow, backrow []int) {
	sort.Ints(a)
	sort.Ints(b)
	if a[0] > b[0] {
		return b, a
	}
	return a, b
}

func recurse(frontrow, backrow []int) bool {
	if len(frontrow) == 0 {
		return true
	}
	if frontrow[0] < backrow[0] {
		return recurse(frontrow[1:], backrow[1:])
	}
	return false
}
