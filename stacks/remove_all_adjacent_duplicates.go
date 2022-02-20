package stacks

import "strings"

// RemoveAllAdjacentDuplicates ...
//
// T -> O(n) - where n is the length of the input string.
// S -> O(n) - since we might place all chars of the string on the stack at once, if there are no duplicates.
func RemoveAllAdjacentDuplicates(s string) string {
	var stack = []string{string(s[0])}
	for i := 1; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, string(s[i]))
			continue
		}

		if string(s[i]) == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, string(s[i]))
	}

	return strings.Join(stack, "")
}
