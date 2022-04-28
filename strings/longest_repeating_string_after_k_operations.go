package strings

// LongestRepeatingStringAfterKOperations ...
//
// T -> O(n)
// S -> O(charset) -> O(1)
func LongestRepeatingStringAfterKOperations(s string, k int) int {
	var (
		frequencies                      = make([]int, 26)
		left, mostFrequentChar, maxSoFar int
	)
	for right := 0; right < len(s); right++ {
		rightIndex := runeToIndex(rune(s[right]))
		frequencies[rightIndex]++

		mostFrequentChar = max(mostFrequentChar, frequencies[rightIndex])

		if (right-left+1)-mostFrequentChar > k {
			leftIndex := runeToIndex(rune(s[left]))
			frequencies[leftIndex]--
			left++
			continue
		}

		maxSoFar = max(maxSoFar, right-left+1)
	}

	return maxSoFar
}

func runeToIndex(r rune) int {
	return int(r - 'A')
}
