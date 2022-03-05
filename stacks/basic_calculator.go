package stacks

import (
	"strconv"
)

type Operand rune

const (
	add Operand = '+'
	sub         = '-'
	mul         = '*'
	div         = '/'
)

func (o Operand) Reduce(a, b int) int {
	switch o {
	case '+':
		return a + b
	case '-':
		return a + b
	case '*':
		return a * b
	case '/':
		return a / b
	default:
		return 0
	}
}

// BasicCalculator ...
//
// T -> O(n) where `n` is the number of chars in the string `s`
// S -> O(1) we could have used a stack here for `O(n)`, but since this will have at most 1/2 items on it we can store as a var.
func BasicCalculator(s string) int {
	var (
		sumSoFar, previousNumber, currentNumber int
		operand                                 Operand = '+'
	)
	for _, r := range s {
		switch {
		case isDigit(r):
			c, _ := strconv.Atoi(string(r))
			currentNumber = setSign(operand, abs(currentNumber*10)+c)
		case r == '+', r == '-':
			previousNumber = operand.Reduce(previousNumber, currentNumber)
			sumSoFar += previousNumber

			currentNumber, previousNumber = 0, 0
			operand = Operand(r)

		case r == '*', r == '/':
			previousNumber = operand.Reduce(previousNumber, currentNumber)
			currentNumber = 0

			operand = Operand(r)
		}

	}

	sumSoFar += operand.Reduce(previousNumber, currentNumber)

	return sumSoFar
}

func isDigit(r rune) bool {
	if r > '9' || r < '0' {
		return false
	}

	return true
}

func setSign(op Operand, num int) int {
	if op == sub && num > 0 {
		return -num
	}
	return num
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
