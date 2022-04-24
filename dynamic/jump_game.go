package dynamic

// JumpGameBottomUp ...
//
// T -> O(j ** 2) where `j`
// S -> O(j)
func JumpGameBottomUp(jumps []int) bool {
	if len(jumps) == 0 {
		return true
	}

	var dp = make([]bool, len(jumps))
	dp[0] = true

	for i := 0; i < len(jumps)-2; i++ {
		if dp[i] == false {
			return false
		}

		var (
			jump         = jumps[i]
			furthestJump = min(len(jumps)-1, i+jump)
		)
		for k := i; k <= furthestJump; k++ {
			dp[k] = true
		}
	}

	return dp[len(jumps)-1]
}

// JumpGameGreedy ...
//
// T -> O(j) where `j` is the number of jumps.
// S -> O(1)
func JumpGameGreedy(jumps []int) bool {
	var maxReachableIndex int
	for i, jump := range jumps {
		if i > maxReachableIndex {
			return false
		}

		maxReachableIndex = max(maxReachableIndex, i+jump)
	}

	return true
}
