package dynamic

// LongestCommonSubSequenceRecursive ...
//
// T -> O(m(n**2)) where `n` is the length of the first string & `m` is the length of the second string.
// S -> O(m * n ** 2)
func LongestCommonSubSequenceRecursive(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	firstOccurance := findFirstOccurance(rune(s1[0]), s2)

	var c1, c2 int
	c1 = LongestCommonSubSequenceRecursive(s1[1:], s2)
	if firstOccurance >= 0 && firstOccurance < len(s2) {
		c2 = 1 + LongestCommonSubSequenceRecursive(s1[1:], s2[firstOccurance+1:])
	}

	return max(c1, c2)
}

func findFirstOccurance(c rune, s string) int {
	for i, r := range s {
		if r == c {
			return i
		}
	}

	return -1
}

// LongestCommonSubSequenceRecursiveWithMemo ...
//
// T -> O(m * (n ** 2))
// S -> O(m * n)
func LongestCommonSubSequenceRecursiveWithMemo(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	// T -> O(n * m)
	// S -> O(n * m)
	var memo = make([][]int, 0, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		memo = append(memo, make([]int, len(s2)+1))
	}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			memo[i][j] = -1
		}
	}

	// T -> O(m * (n ** 2))
	// S -> O(1)
	var solve func(i, j int) int
	solve = func(i, j int) int {
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		var o1, o2 int
		o1 = solve(i+1, j) // [-1, -1, -1, -1, 0] when i = len(s1+1)

		// T -> O(n)
		firstOccurance := findFirstOccurance(rune(s1[i]), s2[j:])
		if firstOccurance != -1 {
			o2 = 1 + solve(i+1, firstOccurance+j+1)
		}

		memo[i][j] = max(o1, o2)
		return memo[i][j]
	}

	return solve(0, 0)
}

// LongestCommonSubSequenceDynamicBottomUp ...
//
// T -> O(n * m)
// S -> O(n * m)
func LongestCommonSubSequenceDynamicBottomUp(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	// T -> O(n * m)
	// S -> O(n * m)
	var dp = make([][]int, 0, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp = append(dp, make([]int, len(s2)+1))
	}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			dp[i][j] = -1
		}
	}

	// T -> O(n * m)
	// S -> O(n * m)
	for j := len(s1) - 1; j >= 0; j-- {
		for i := len(s2) - 1; i >= 0; i-- {
			if s1[j] == s2[i] {
				dp[j][i] = dp[j+1][i+1] + 1
				continue
			}

			dp[j][i] = max(dp[j+1][i], dp[j][i+1])
		}
	}

	return dp[0][0]
}

// LongestCommonSubSequenceDynamicBottomUpOptimized ...
// T -> O(n * m)
// S -> O(min(n, m))
func LongestCommonSubSequenceDynamicBottomUpOptimized(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}

	// Ensure that the first string is indeed the shortest.
	if len(s2) < len(s1) {
		return LongestCommonSubSequenceDynamicBottomUpOptimized(s2, s1)
	}

	// S -> O(2 * n) -> O(n)
	var (
		top    = make([]int, len(s1)+1)
		bottom = make([]int, len(s1)+1)
	)

	// T -> O(min(m, n))
	for j := len(s2) - 1; j >= 0; j-- {
		for i := len(s1) - 1; i >= 0; i-- {
			if s2[j] == s1[i] {
				top[i] = 1 + bottom[i+1]
				continue
			}

			top[i] = max(top[i+1], bottom[i])
		}

		// Store the previous in the temp variable as to retain the pointer.
		tmp := bottom
		bottom = top
		top = tmp
	}

	return bottom[0]
}
