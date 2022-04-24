package dynamic

// MinimumASCIIDeleteTwoStrings ...
//
// T -> O(len(s1) * len(s2))
// S -> O(len(s1) * len(s2))
func MinimumASCIIDeleteTwoStrings(s1, s2 string) int {
	var dp = make([][]int, 0, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp = append(dp, make([]int, len(s2)+1))
	}

	for j := len(s1) - 1; j >= 0; j-- {
		dp[j][len(s2)-1] = dp[j+1][len(s2)-1] + int(s2[j])
	}

	for i := len(s2) - 1; i >= 0; i-- {
		dp[len(s1)-1][i] = dp[len(s1)-1][i+1] + int(s1[i])
	}

	for j := len(s1) - 1; j >= 0; j-- {
		for i := len(s2) - 1; i >= 0; i-- {
			if s1[j] == s2[i] {
				dp[j][i] = dp[j+1][i+1]
				continue
			}

			dp[j][i] = min(dp[j+1][i]+int(s1[j]), dp[j][i+1]+int(s2[i]))
		}
	}

	return dp[0][0]
}
