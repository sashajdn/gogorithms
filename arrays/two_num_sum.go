package arrays

import (
	"sort"
)

// TwoNumberSum_Iterative ...
//
// T -> O(n ** 2)
// S -> O(1)
func TwoNumberSum_Iterative(array []int, target int) []int {
	for _, n := range array {
		for _, nPrime := range array {
			if n == nPrime {
				continue
			}

			if n+nPrime == target {
				return []int{n, nPrime}
			}
		}
	}

	return []int{}
}

// TwoNumberSum_Sorted ...
//
// T -> O(nlog(n))
// S -> O(1)
func TwoNumberSum_Sorted(array []int, target int) []int {
	sort.Ints(array)
	l, r := 0, len(array)-1
	for l < r {
		currentSum := array[l] + array[r]
		if currentSum == target {
			return []int{array[l], array[r]}
		} else if currentSum < target {
			l++
		} else {
			r--
		}
	}
	return make([]int, 0)
}

// TwoNumberSum_HashmapOnePass ...
//
// T -> O(n) where n is the number of numbers in numbers.
// S -> O(n)
func TwoNumberSum_HashmapOnePass(numbers []int, target int) []int {
	// Since we want to find a + b = target.
	// We can precompute (a - target) & store in a hashmap: {a-target: index of a}
	//
	// We can then in the same loop check if b exists - doing it in the same loop
	// means we don't have to check to i==j case (e.g sum against ourself)
	var hm = make(map[int]int)
	for i, number := range numbers {
		if j, ok := hm[number]; ok {
			return []int{i, j}
		}

		hm[number] = i
	}

	return []int{}
}
