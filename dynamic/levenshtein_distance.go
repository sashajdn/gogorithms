package dynamic

func LevenshteinDistance(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}

	if len(b) == 0 {
		return len(a)
	}

	l, s := longestString(a, b)

	oddEdits := make([]int, len(l)+1)
	evenEdits := make([]int, len(l)+1)

	var currentEdits, previousEdits []int

	for i := 0; i < len(l)+1; i++ {
		evenEdits[i] = i
	}

	for j := 1; j < len(l)+1; j++ {
		if j%2 == 0 {
			currentEdits, previousEdits = evenEdits, oddEdits
		} else {
			currentEdits, previousEdits = oddEdits, evenEdits
		}

		currentEdits[0] = j

		for k := 1; k < len(s)+1; k++ {
			if l[k-1] == s[k-1] {
				currentEdits[k] = previousEdits[k-1]
				continue
			}
			currentEdits[k] = 1 + min(
				currentEdits[k-1],
				previousEdits[k-1],
				previousEdits[k],
			)
		}
	}
	return currentEdits[len(l)-1]
}

func longestString(a, b string) (string, string) {
	if len(a) >= len(b) {
		return a, b
	}
	return b, a
}

func min(vals ...int) int {
	c := vals[0]
	for _, v := range vals[1:] {
		if v < c {
			c = v
		}
	}
	return c
}
