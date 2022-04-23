package strings

// ValidParentheses ...
//
// T -> O(s) where `s` is the length of the input string.
// S -> O(s)
func ValidParentheses(s string) bool {
	var stack []rune
	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}

			var top rune
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]

			if top != matching(r) {
				return false
			}
		}
	}

	return len(stack) == 0
}

func matching(r rune) rune {
	switch r {
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	default:
		return 'u'
	}
}
