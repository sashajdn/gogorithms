package dynamic

// InterweavingStrings_Recursive ...
//
// T -> O(2 ** n)
// S -> O(2 ** n)
func InterweavingStrings_Recursive(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	var recurse func(p1, p2 int) bool
	recurse = func(p1, p2 int) bool {
		var p3 = p1 + p2
		if p3 >= len(s3) {
			return true
		}

		if p1 >= len(s1) {
			if s2[p2] == s3[p3] {
				return recurse(p1, p2+1)
			}

			return false
		}

		if p2 >= len(s2) {
			if s1[p1] == s3[p3] {
				return recurse(p1+1, p2)
			}

			return false
		}

		if s1[p1] == s3[p3] && s2[p2] == s3[p3] {
			return recurse(p1+1, p2) || recurse(p1, p2+1)
		}

		if s1[p1] == s3[p3] {
			return recurse(p1+1, p2)
		}

		if s2[p2] == s3[p3] {
			return recurse(p1, p2+1)
		}

		return false
	}

	return recurse(0, 0)
}

// InterweavingStrings_TopDown ...
//
// T -> O(n * m)
// S -> O(n * m)
func InterweavingStrings_TopDown(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	type tuple [2]int
	var memo = map[tuple]bool{}

	var recurse func(p1, p2 int) bool
	recurse = func(p1, p2 int) bool {
		var (
			p3 = p1 + p2
			t  = tuple{p1, p2}
		)
		if p3 >= len(s3) {
			return true
		}

		if ok, found := memo[t]; found {
			return ok
		}

		if p1 < len(s1) && s1[p1] == s3[p3] && recurse(p1+1, p2) {
			return true
		}
		if p2 < len(s2) && s2[p2] == s3[p3] && recurse(p1, p2+1) {
			return true
		}

		memo[t] = false
		return memo[t]
	}

	return recurse(0, 0)
}

// InterweavingStrings_BottomUp2D ...
//
// T -> O(n * m)
// S -> O(n * m)
func InterweavingStrings_BottomUp2D(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	var dp = make([][]bool, 0, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp = append(dp, make([]bool, len(s2)+1))
	}
	dp[len(s1)][len(s2)] = true

	for j := len(s1) - 1; j >= 0; j-- {
		dp[j][len(s2)] = dp[j+1][len(s2)] && s1[j] == s3[j+len(s2)-1]
	}

	for i := len(s2) - 1; i >= 0; i-- {
		dp[len(s1)][i] = dp[len(s1)][i+1] && s2[i] == s3[i+len(s1)-1]
	}

	for j := len(s1) - 1; j >= 0; j-- {
		for i := len(s2) - 1; i >= 0; i-- {
			if s1[j] == s3[j+i] && s2[i] == s3[j+i] {
				dp[j][i] = dp[j+1][i] || dp[j][i+1]
				continue
			}

			if s1[j] == s3[j+i] {
				dp[j][i] = dp[j+1][i]
				continue
			}

			if s2[i] == s3[j+i] {
				dp[j][i] = dp[j][i+1]
				continue
			}
		}
	}

	return dp[0][0]
}

// InterweavingStrings_BottomUp1D ...
//
// T -> O(n * m)
// S -> O(min(n, m))
func InterweavingStrings_BottomUp1D(s1, s2, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1)+len(s2) == 0 {
		return true
	}

	if len(s2) > len(s1) {
		return InterweavingStrings_BottomUp1D(s2, s1, s3)
	}

	var top, bottom = make([]bool, len(s2)+1), make([]bool, len(s2)+1)

	bottom[len(s2)] = true
	for i := len(s2) - 1; i >= 0; i-- {
		bottom[i] = s2[i] == s3[i+len(s1)] && bottom[i+1]
	}
	top[len(s2)] = s1[len(s1)-1] == s3[len(s1)-1+len(s2)] && bottom[len(s2)]

	for j := len(s1) - 1; j >= 0; j-- {
		for i := len(s2) - 1; i >= 0; i-- {
			if s1[j] == s3[i+j] && s2[i] == s3[i+j] {
				top[i] = top[i+1] || bottom[i]
				continue
			}

			if s1[j] == s3[i+j] {
				top[i] = bottom[i]
				continue
			}

			if s2[i] == s3[i+j] {
				top[i] = top[i+1]
			}

			top[i] = false
		}

		bottom, top = top, bottom
		top[len(s2)] = s1[j] == s3[j+len(s2)] && bottom[len(s2)]
	}

	return bottom[0]
}
