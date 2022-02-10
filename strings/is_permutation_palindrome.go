package strings

func isPermutationPalindrome(s string) bool {
	set := make(map[rune]struct{})

	for _, r := range s {
		if _, ok := set[r]; ok {
			delete(set, r)
			continue
		}

		set[r] = struct{}{}
	}

	return len(set) < 2
}
