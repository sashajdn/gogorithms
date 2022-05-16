package dynamic

// EditDistanceII_TopDown2D ...
//
// T -> O(ab)
// S -> O(ab)
func EditDistanceII_TopDown2D(a, b string) int {
	if len(a)*len(b) == 0 {
		return max(len(a), len(b))
	}

	var dp = make([][]int, 0, len(a)+1)
	for i := 0; i < len(a)+1; i++ {
		dp = append(dp, make([]int, len(b)+1))
	}

	for j := len(a) - 1; j >= 0; j-- {
		dp[j][len(b)] = len(a) - j - 1
	}
	for i := len(b) - 1; i >= 0; i-- {
		dp[len(a)][i] = len(b) - i - 1
	}

	for j := len(a) - 1; j >= 0; j-- {
		for i := len(b) - 1; i >= 0; i-- {
			if a[j] == b[i] {
				dp[j][i] = dp[j+1][i+1]
				continue
			}

			dp[j][i] = min(dp[j+1][i], dp[j][i+1]) + 1
		}
	}

	return dp[0][0]
}

// EditDistancesII_TopDown1D ...
//
// T -> O(ab)
// S -> O(min(a, b))
func EditDistanceII_TopDown1D(a, b string) int {
	if len(a)*len(b) == 0 {
		return max(len(a), len(b))
	}
	if len(b) > len(a) {
		return EditDistanceII_TopDown1D(b, a)
	}

	var top, bottom = make([]int, len(b)+1), make([]int, len(b)+1)
	for i := len(b) - 1; i >= 0; i-- {
		bottom[i] = len(b) - i
	}
	top[len(b)] = 1

	for j := len(a) - 1; j >= 0; j-- {
		for i := len(b) - 1; i >= 0; i-- {
			if a[j] == b[i] {
				top[i] = bottom[i+1]
				continue
			}

			top[i] = min(bottom[i+1], min(bottom[i], top[i+1])) + 1
		}

		bottom, top = top, make([]int, len(b)+1)
		top[len(b)] = len(a) - j + 1
	}

	return bottom[0]
}
