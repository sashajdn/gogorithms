package dynamic

import "sort"

// CombinationSumIV ...
//
// T -> O(n * t + nlog(n))
// S -> O(n)
func CombinationSumIV(numbers []int, target int) int {
	var memo = map[int]int{}
	sort.Ints(numbers)

	var combinations func(remaining int) int
	combinations = func(remaining int) int {
		if remaining <= 0 {
			return 1
		}

		if stored, ok := memo[remaining]; ok {
			return stored
		}

		var result int
		for _, number := range numbers {
			if remaining-number < 0 {
				continue
			}

			result += combinations(remaining - number)
		}

		memo[remaining] = result
		return result
	}

	return combinations(target)
}
