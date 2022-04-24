package strings

// LongestNonRepeatingSubString ...
//
// T -> O(s) where `s` is the lenght of the input string
// S -> O(min(s, charset)) where `charset` is the length of the possible charset.
func LongestNonRepeatingSubString(s string) int {
	var (
		hm             = make(map[rune]int, 128)
		maxSoFar, left int
	)

	for right := 0; right < len(s); right++ {
		if last, ok := hm[rune(s[right])]; ok {
			left = max(left, last+1)
		}

		hm[rune(s[right])] = right
		maxSoFar = max(maxSoFar, right-left+1)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
