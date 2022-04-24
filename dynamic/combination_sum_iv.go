package dynamic

import "sort"

// CombinationSumIV_TopDown ...
//
// T -> O(n * t + nlog(n))
// S -> O(t)
func CombinationSumIV_TopDown(numbers []int, target int) int {
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

// CombinationSumIV_BottomUp ...
//
// T -> O(n * t + nlog(n))
// S -> O(t)
func CombinationSumIV_BottomUp(numbers []int, target int) int {
	sort.Ints(numbers)

	var dp = make([]int, target+1)
	dp[0] = 1

	for i := 1; i < target+1; i++ {
		for _, number := range numbers {
			if number > i {
				continue
			}

			dp[i] += dp[i-number]
		}
	}

	return dp[target]
}
