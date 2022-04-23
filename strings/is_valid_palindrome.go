package strings

import "strings"

// ValidPalindrome ...
//
// T -> O(s) where `s` is the length of the input string.
// S -> O(1)
func ValidPalindrome(s string) bool {
	var l, r = 0, len(s) - 1
	for l < r {
		if !isAlphanumeric(rune(s[l])) {
			l++
			continue
		}
		if !isAlphanumeric(rune(s[r])) {
			r--
			continue
		}

		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}

		l++
		r--
	}

	return true
}

func isAlphanumeric(r rune) bool {
	return (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
