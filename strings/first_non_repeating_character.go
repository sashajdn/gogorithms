package strings

type runeSlice []rune

func (r runeSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r runeSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runeSlice) Len() int {
	return len(r)
}

// FirstNonRepeatingCharacter takes in a string of lowercase English chars & returns the index of the first non repeating chars.
//
// T: O(n) -> n
// S: O(n) -> 26 -> 1 (since english chars only, limited number of possible chars)
func FirstNonRepeatingCharacter(s string) int {
	if len(s) == 0 {
		return -1
	}

	var m = map[rune]int{}
	for _, r := range s {
		if _, ok := m[r]; !ok {
			m[r] = 1
			continue
		}

		m[r]++
	}

	for i, r := range s {
		v, ok := m[r]
		if !ok {
			return -1 // This is actually a failure & shouldn't happen.
		}
		if v == 1 {
			return i
		}
	}

	return -1
}
