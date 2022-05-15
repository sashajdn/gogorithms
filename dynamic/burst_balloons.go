package dynamic

import "fmt"

// BurstBalloonsBruteForce ...
//
// T -> O(n ** n)
// S -> O(n ** n)
func BurstBalloonsBruteForce(nums []int) int {
	memo := map[string]int{}

	var recurse func(nums []int) int
	recurse = func(nums []int) int {
		switch len(nums) {
		case 0:
			return 0
		case 1:
			return nums[0]
		case 2:
			v := nums[0] * nums[1]
			return v + max(nums[0], nums[1])
		}

		key := valuesToKey(nums...)
		if v, ok := memo[key]; ok {
			return v
		}

		var (
			maxSoFar     int
			extendedNums = []int{1}
		)
		extendedNums = append(extendedNums, nums...)
		extendedNums = append(extendedNums, 1)

		for i := 1; i < len(extendedNums)-1; i++ {
			total := extendedNums[i-1] * extendedNums[i] * extendedNums[i+1]

			j := i - 1
			slice := append([]int{}, nums[:j]...)
			if i < len(nums) {
				slice = append(slice, nums[j+1:]...)
			}

			maxSoFar = max(maxSoFar, total+recurse(slice))
		}

		memo[key] = maxSoFar
		return maxSoFar
	}

	return recurse(nums)
}

func valuesToKey(vv ...int) string {
	var s string
	for _, v := range vv {
		s += fmt.Sprintf("%d*", v)
	}
	return s
}

// BurstBalloonsDynamicBottomUp ...
//
// T -> O(n ** 3)
// S -> O(n ** 2)
func BurstBalloonsDynamicBottomUp(nums []int) int {
	type tuple [2]int

	var (
		extendedNums = append(append([]int{1}, nums...), 1)
		memo         = map[tuple]int{}
	)

	// T -> O(n ** 2)[dfs] * O(n)[linear search on dfs]
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l > r {
			return 0
		}

		var t = tuple{l, r}
		if v, ok := memo[t]; ok {
			return v
		}

		memo[t] = 0
		for i := l; i <= r; i++ {
			coins := (extendedNums[l-1] * extendedNums[i] * extendedNums[r+1]) + dfs(l, i-1) + dfs(i+1, r)
			memo[t] = max(memo[t], coins)
		}
		return memo[t]
	}

	return dfs(1, len(extendedNums)-2)
}
