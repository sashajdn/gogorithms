package dynamic

// Fibonacci_BottomUp ...
//
// T -> O(n)
// S -> O(n)
func Fibonacci_BottomUp(n int) int {
	var memo = map[int]int{}

	var dfs func(n int) int
	dfs = func(n int) int {
		switch n {
		case 0:
			return 0
		case 1, 2:
			return 1
		}

		if v, ok := memo[n]; ok {
			return v
		}

		memo[n] = dfs(n-1) + dfs(n-2)
		return memo[n]
	}

	return dfs(n)
}

// Fibonacci_TopDown ...
//
// T -> O(n)
// S -> O(n)
func Fibonacci_TopDown(n int) int {
	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	}

	var dp = make([]int, n+1)
	dp[1], dp[2] = 1, 1

	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func Fibonacci_TopDownPointers(n int) int {
	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	}

	var first, second = 1, 1
	for i := 2; i < n+1; i++ {
		first, second = second, first+second
	}

	return first
}
