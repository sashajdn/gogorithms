package dynamic

import (
	"fmt"
	"strconv"
)

// NumberOfWaysOfDecoding_TopDown ...
//
// T -> O(n) where `n` is the length of the string, as we prune the total search space with memoization.
// S -> O(max(n, charset))
func NumberOfWaysOfDecoding_TopDown(s string) int {
	var encoding = map[string]struct{}{}
	for i := 1; i < 27; i++ {
		encoding[fmt.Sprintf("%d", i)] = struct{}{}
	}

	var memo = map[string]int{}
	var find func(s string) int
	find = func(s string) int {
		switch len(s) {
		case 0:
			return 1
		case 1:
			if _, ok := encoding[s]; ok {
				return 1
			}
			return 0
		}

		first, firstTail := string(s[0]), string(s[1:])
		second, secondTail := string(s[:2]), string(s[2:])

		var result int
		if _, ok := encoding[first]; ok {
			v, ok := memo[firstTail]
			switch {
			case ok:
				result += v
			default:
				result += find(firstTail)
			}
		}

		if _, ok := encoding[second]; ok {
			v, ok := memo[secondTail]
			switch {
			case ok:
				result += v
			default:
				result += find(secondTail)
			}
		}

		memo[s] = result
		return result
	}

	return find(s)
}

// NumberOfWaysOfDecoding_DPIterative ...
//
// T -> O(s)
// S -> O(s)
func NumberOfWaysOfDecoding_DPIterative(s string) int {
	if len(s) == 0 {
		return 0
	}

	var dp = make([]int, len(s)+1)

	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}

	for i := 2; i < len(s)+1; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}

		if isValidDoubleDigit(string(s[i-2 : i])) {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

// NumberOfWaysOfDecoding_TwoPointer ...
//
// T -> O(s)
// S -> O(1)
func NumberOfWaysOfDecoding_TwoPointer(s string) int {
	if len(s) == 0 {
		return 0
	}

	var twoBefore, oneBefore = 1, 0
	if s[0] != '0' {
		oneBefore = 1
	}

	for i := 2; i < len(s)+1; i++ {
		var current int
		if s[i-1] != '0' {
			current += oneBefore
		}

		if isValidDoubleDigit(string(s[i-2 : i])) {
			current += twoBefore
		}

		twoBefore, oneBefore = oneBefore, current
	}

	return oneBefore
}

func isValidDoubleDigit(s string) bool {
	asInt, _ := strconv.Atoi(s)
	return asInt >= 10 && asInt <= 26
}
