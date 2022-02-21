package strings

import (
	"strconv"
)

// CalculateManyDigit_WithStack ...
//
// T -> O(n) where n is the number of chars in the string.
// S -> O(n)
func CalculateManyDigit_WithStack(s string) int {
	var (
		stack         []int
		currentNumber int
		operation     = "+"
	)

	for i := 0; i < len(s); i++ {
		if isDigit(string(s[i])) {
			currentNumber = (10 * currentNumber) + strToInt(string(s[i]))
		}

		if string(s[i]) == " " {
			continue
		}

		if isDigit(string(s[i])) && i != len(s)-1 {
			continue
		}

		switch operation {
		case "+":
			stack = append(stack, currentNumber)
		case "-":
			stack = append(stack, -currentNumber)
		case "*":
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, top*currentNumber)
		case "/":
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, top/currentNumber)
		}

		operation = string(s[i])
		currentNumber = 0
	}

	var result int
	for _, number := range stack {
		result += number
	}

	return result
}

func isDigit(c string) bool {
	switch c {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return true
	default:
		return false
	}
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
