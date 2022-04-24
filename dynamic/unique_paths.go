package dynamic

// UniquePaths ...
// T -> O(m * n)
// S -> O(min(n, m))
func UniquePaths(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if n > m {
		return UniquePaths(n, m)
	}

	var dp = make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			dp[i] += dp[i-1]
		}
	}

	return dp[n-1]
}
