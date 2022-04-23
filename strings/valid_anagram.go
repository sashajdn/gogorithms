package strings

// ValidAnagram ...
//
// T -> O(n) where `n` is the length of the word
// S -> O(n) where `n` is the length of the word
func ValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var hm = make(map[rune]int, len(s))
	for _, r := range s {
		hm[r]++
	}

	for _, r := range t {
		v, ok := hm[r]

		switch {
		case !ok:
			return false
		case v == 0:
			return false
		case v == 1:
			delete(hm, r)
		default:
			hm[r]--
		}
	}

	return len(hm) == 0
}
