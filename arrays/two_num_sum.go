package arrays

import  (
	"sort"
)

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	l, r := 0, len(array) - 1
	for l < r {
		currentSum := array[l] + array[r]
		if currentSum == target{
			return []int{array[l], array[r]}
		} else if currentSum < target {
			l++
		} else {
			r--
		}
	}
	return make([]int, 0)
}
