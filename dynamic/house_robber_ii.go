package dynamic

// HouseRobberII ...
//
// T -> O(n) where `n` the number of houses.
// S -> O(n)
func HouseRobberII_Dynamic(houses []int) int {
	switch len(houses) {
	case 0:
		return 0
	case 1:
		return houses[0]
	}

	return max(
		robHouseDynamic(houses, 0, len(houses)-2),
		robHouseDynamic(houses, 1, len(houses)-1),
	)
}

func robHouseDynamic(houses []int, start, end int) int {
	var (
		dp       = make([]int, len(houses))
		maxSoFar int
	)

	for i := start; i <= end; i++ {
		switch {
		case i > 2:
			dp[i] = max(dp[i-3], dp[i-2]) + houses[i]
		case i > 1:
			dp[i] = dp[i-2] + houses[i]
		default:
			dp[i] = houses[i]
		}

		maxSoFar = max(maxSoFar, dp[i])
	}

	return maxSoFar
}

// HouseRobberII_Pointers ...
//
// T -> O(n)
// S -> O(1)
func HouseRobberII_Pointers(houses []int) int {
	switch len(houses) {
	case 0:
		return 0
	case 1:
		return houses[0]
	}

	return max(
		robHousePointers(houses, 0, len(houses)-2),
		robHousePointers(houses, 1, len(houses)-1),
	)
}

func robHousePointers(houses []int, start, end int) int {
	var (
		maxSoFar                             int
		previousPrevious, previous, adjacent int
	)

	for i := start; i <= end; i++ {
		var current int
		switch {
		case i > 2:
			current = max(previousPrevious, previous) + houses[i]
		case i > 1:
			current = previous + houses[i]
		default:
			current = houses[i]
		}

		maxSoFar = max(maxSoFar, current)
		previousPrevious, previous, adjacent = previous, adjacent, current
	}

	return maxSoFar
}
