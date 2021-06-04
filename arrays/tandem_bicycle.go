package arrays

import "sort"

// TandemBicycle : O(T) -> O(nlogn), O(S) -> O(1)
func TandemBicycle(redShirtSpeeds, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	var s = sort.Ints
	if fastest {
		s = reverseSort
	}
	s(blueShirtSpeeds)
	var total int
	for i, v := range redShirtSpeeds {
		total += max(v, blueShirtSpeeds[i])
	}
	return total
}

func reverseSort(a []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
}
