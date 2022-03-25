package stacks

import "strconv"

// BasicCalculator ...
//
// T -> O(n) where `n` is the number of chars in the string `s`
// S -> O(1) we could have used a stack here for `O(n)`, but since this will have at most 1/2 items on it we can store as a var.
func BasicCalculator(s string) int {
	var (
		currentNumber int
		currentValue  int
		total         int
		operator      rune
	)

	for _, r := range s {
		switch {
		case isDigit(r):
			currentNumber = 10*currentNumber + runeToInt(r)
			continue
		case r == ' ':
			continue
		}

		switch r {
		case '+', '-':
			total += calculate(operator, currentValue, currentNumber)
			currentNumber, currentValue = 0, 0
			operator = r
		case '*', '/':
			currentValue = calculate(operator, currentValue, currentNumber)
			currentNumber = 0
			operator = r
		}
	}

	total += calculate(operator, currentValue, currentNumber)
	return total
}

func runeToInt(r rune) int {
	i, _ := strconv.Atoi(string(r))
	return i
}

func isDigit(r rune) bool {
	if r < '0' || r > '9' {
		return false
	}
	return true
}

func calculate(op rune, a, b int) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '/':
		return a / b
	case '*':
		return a * b
	}

	return 0
}
