package dynamic

import (
	"strconv"
)

// MahjongTiles ...
//
// T -> O(2 ** n) -> O(s)
// S -> O(s)
func MahjongTiles_BottomUp(tiles string) bool {
	var memo = map[string]bool{}

	var dfs func(tiles string) bool
	dfs = func(tiles string) bool {
		switch len(tiles) {
		case 0:
			return true
		case 1:
			return false
		case 2:
			return isPair(tiles)
		}

		if ok, found := memo[tiles]; found {
			return ok
		}

		var (
			pair    = string(tiles[:2])
			triplet = string(tiles[:3])
		)
		isValidPair := isPair(pair) && dfs(tiles[2:])
		isValidTriplet := (isIncreasingSequence(triplet) || (isTriplet(triplet))) && dfs(tiles[3:])

		memo[tiles] = isValidPair || isValidTriplet
		return memo[tiles]
	}

	return dfs(tiles)
}

func isPair(s string) bool {
	if len(s) != 2 {
		return false
	}

	if s[0] == s[1] {
		return true
	}

	return false
}

func isIncreasingSequence(s string) bool {
	if len(s) != 3 {
		return false
	}

	var previous = -1
	for _, r := range s {
		current, _ := strconv.Atoi(string(r))

		if current <= previous {
			return false
		}

		previous = current
	}

	return true
}

func isTriplet(s string) bool {
	if len(s) != 3 {
		return false
	}

	if s[0] == s[1] && s[0] == s[2] {
		return true
	}

	return false
}
