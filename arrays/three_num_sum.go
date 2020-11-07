package arrays

import (
	"sort"
)

// ThreeNumberSum T -> O(n**2), S -> O(n)
func ThreeNumberSum(array []int, target int) [][]int {
	triplets := [][]int{}
	sort.Ints(array)

	for i := 0; i < len(array)-2; i++ {

		l, r := i+1, len(array)-1

		for l < r {
			total := sum(array[i], array[l], array[r])
			if total == target {
				triplets = append(triplets, []int{array[i], array[l], array[r]})
				l++
				r--
			} else if total < target {
				l++
			} else {
				r--
			}
		}
	}
	return triplets
}

func sum(vals ...int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}
