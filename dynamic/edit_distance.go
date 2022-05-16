package dynamic

// EditDistance ...
//
// T -> O(n * m) where `n` is the length of word1 & `m` is the length of word2.
// S -> O(min(n, m))
func EditDistance(word1, word2 string) int {
	if len(word1)*len(word2) == 0 {
		return len(word1) + len(word2)
	}

	if len(word2) > len(word1) {
		return EditDistance(word2, word1)
	}

	var top, bottom = make([]int, len(word2)+1), make([]int, len(word2)+1)
	for i := 1; i < len(word2)+1; i++ {
		top[i] = i
	}
	bottom[0] = 1

	for j := 1; j < len(word1)+1; j++ {
		for i := 1; i < len(word2)+1; i++ {
			if word1[j-1] == word2[i-1] {
				bottom[i] = top[i-1]
				continue
			}

			bottom[i] = min(min(top[i], bottom[i-1]), top[i-1]) + 1
		}

		top, bottom = bottom, make([]int, len(word2)+1)
		bottom[0] = j + 1
	}

	return top[len(word2)]
}
