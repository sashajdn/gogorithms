package dynamic

import (
	"strconv"
)

// DecodeWays_BottomUp ...
//
// T -> O(n)
// S -> O(n)
func DecodeWays_BottomUp(s string) int {
	switch len(s) {
	case 0:
		return 0
	case 1:
		if s[len(s)-1] == '0' {
			return 0
		}

		return 1
	}

	var mapping = map[int]struct{}{}
	for i := 1; i <= 26; i++ {
		mapping[i] = struct{}{}
	}

	var memo = map[int]int{}

	var dfs func(index int) int
	dfs = func(index int) int {
		if ways, ok := memo[index]; ok {
			return ways
		}

		switch {
		case index >= len(s):
			return 1
		case s[index] == '0':
			return 0
		case index >= len(s)-1:
			return 1
		}

		var total int
		if _, ok := mapping[stringToInt(string(s[index]))]; ok {
			total += dfs(index + 1)
		}

		if _, ok := mapping[stringToInt(string(s[index:index+2]))]; ok {
			total += dfs(index + 2)
		}

		memo[index] = total
		return memo[index]
	}

	return dfs(0)
}

func stringToInt(s string) int {
	sAsInt, _ := strconv.Atoi(s)
	return sAsInt
}

// DecodeWays_TopDown ...
//
// T -> O()
// S -> O()
func DecodeWays_TopDown(s string) int {
	switch len(s) {
	case 0:
		return 0
	case 1:
		if s[0] == '0' {
			return 0
		}
		return 1
	}

	var mapping = map[int]struct{}{}
	for i := 1; i <= 26; i++ {
		mapping[i] = struct{}{}
	}

	var dp = make([]int, len(s)+1)
	dp[0] = 1

	first := stringToInt(string(s[0]))
	if _, ok := mapping[first]; ok {
		dp[1] = 1
	}

	for i := 2; i < len(s)+1; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		second := stringToInt(string(s[i-2 : i]))
		if _, ok := mapping[second]; ok {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

// DecodeWays_TopDownTwoPointer ...
//
// T -> O(n)
// S -> O(1)
func DecodeWays_TopDownTwoPointer(s string) int {
	if len(s) == 0 {
		return 0
	}

	var mapping = map[int]struct{}{}
	for i := 0; i <= 26; i++ {
		mapping[i] = struct{}{}
	}

	var first, second = 1, 0
	if s[0] != '0' {
		second = 1
	}

	for i := 2; i < len(s)+1; i++ {
		var total int
		if s[i-1] != '0' {
			total += second
		}

		if _, ok := mapping[stringToInt(s[i-2:i])]; ok {
			total += first
		}

		first, second = second, total
	}

	return second
}
