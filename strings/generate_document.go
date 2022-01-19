package strings

// GenerateDocument given a string of available characters & a string representing a document,
// returns a boolean based on whether the document can be generated from the available characters.
//
// T(n, m) -> O(n+m)
// S(n, m) -> O(min(n, k))  where k is the total set of chars possible.
func GenerateDocument(characters, document string) bool {
	if len(document) == 0 {
		return true
	}

	if len(characters) < len(document) {
		return false
	}

	var count = map[rune]int{}
	for _, r := range characters {
		count[r]++
	}

	for _, r := range document {
		v, ok := count[r]
		if !ok {
			return false
		}

		if v < 1 {
			return false
		}

		count[r]--
	}

	return true
}
