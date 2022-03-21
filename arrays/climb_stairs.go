package arrays

// ClimbStairsWithMemo ...
//
// T -> O(n) where `n` is the total number of stairs to climb; since we use memoization we prune the branches.
// S -> O(n) due to recursing; since we branch the max depth we can reach is `n`
func ClimbStairsWithMemo(n int) int {
	var memo = make([]int, n+1)
	return climbStairsRecurse(0, n, &memo)
}

func climbStairsRecurse(i, n int, memo *[]int) int {
	switch {
	case i > n:
		return 0
	case i == n:
		return 1
	}

	if (*memo)[i] > 0 {
		return (*memo)[i]
	}

	(*memo)[i] = climbStairsRecurse(i+1, n, memo) + climbStairsRecurse(i+2, n, memo)
	return (*memo)[i]
}

// ClimbStairsWithDynamic ...
// T -> O(n)
// S -> O(n)
func ClimbStairsWithDynamic(n int) int {
	var dp = make([]int, n+1)
	dp[0], dp[1] = 1, 2

	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// ClimbStairsFibonacci ...
// T -> O(n)
// S -> O(1)
func ClimbStairsFibonacci(n int) int {
	if n == 1 {
		return 1
	}

	first, second := 1, 2
	for i := 3; i < n+1; i++ {
		third := second + first
		first, second = second, third
	}
	return second
}
