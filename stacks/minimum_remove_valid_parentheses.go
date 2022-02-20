package stacks

import (
	"strings"
)

// MinRemoveToMakeValidParentheses ...
func MinRemoveToMakeValidParentheses(s string) string {
	var count int
	var stack []string

	for i := 0; i < len(s); i++ {
		switch rune(s[i]) {
		case ')':
			if count < 1 {
				continue
			}

			count--
			stack = append(stack, string(s[i]))
		case '(':
			count++
			stack = append(stack, string(s[i]))
		default:
			stack = append(stack, string(s[i]))
		}
	}

	if count == 0 {
		return strings.Join(stack, "")
	}

	var popped []string
	var l = len(stack) - 1
	for i := (l); i >= 0; i-- {
		if count < 1 {
			break
		}
		if len(stack) == 0 {
			break
		}

		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if v == "(" {
			count--
			continue
		}
		popped = append(popped, v)
	}

	for j := len(popped) - 1; j >= 0; j-- {
		stack = append(stack, popped[j])
	}

	return strings.Join(stack, "")
}
