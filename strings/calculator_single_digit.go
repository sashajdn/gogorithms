package strings

import (
	"strconv"
)

func CalculateSingleDigit(s string) int {
	var array []string
	for _, r := range s {
		if r != ' ' {
			array = append(array, string(r))
		}
	}

	return calculateSingleDigit(array)
}

func calculateSingleDigit(array []string) int {
	if len(array) == 1 {
		v, _ := strconv.Atoi(array[0])
		return v
	}

	mid := len(array) / 2

	switch array[mid] {
	case "+":
		return calculateSingleDigit(array[:mid]) + calculateSingleDigit(array[mid+1:])
	case "*":
		return calculateSingleDigit(array[:mid]) * calculateSingleDigit(array[mid+1:])
	case "-":
		return calculateSingleDigit(array[:mid]) - calculateSingleDigit(array[mid+1:])
	case "/":
		return calculateSingleDigit(array[:mid]) / calculateSingleDigit(array[mid+1:])
	}

	operationOrder := orderOfOperation(array, mid-1, mid+1)
	lop, rop := array[mid-1], array[mid+1]

	var sum, _ = strconv.Atoi(array[mid])
	switch operationOrder {
	case OperationOrderLeft:
		switch lop {
		case "+":
			sum += calculateSingleDigit(array[0 : mid-1])
		case "-":
			sum -= calculateSingleDigit(array[0 : mid-1])
		case "*":
			sum *= calculateSingleDigit(array[0 : mid-1])
		case "/":
			sum /= calculateSingleDigit(array[0 : mid-1])
		}
		switch rop {
		case "+":
			sum += calculateSingleDigit(array[mid+2:])
		case "-":
			sum -= calculateSingleDigit(array[mid-2:])
		case "*":
			sum *= calculateSingleDigit(array[mid-2:])
		case "/":
			sum /= calculateSingleDigit(array[mid-2:])
		}
	case OperationOrderRight:
		switch rop {
		case "+":
			sum += calculateSingleDigit(array[mid+2:])
		case "-":
			sum -= calculateSingleDigit(array[mid+2:])
		case "*":
			sum *= calculateSingleDigit(array[mid+2:])
		case "/":
			sum /= calculateSingleDigit(array[mid+2:])
		}
		switch lop {
		case "+":
			sum += calculateSingleDigit(array[0 : mid-1])
		case "-":
			sum -= calculateSingleDigit(array[0 : mid-1])
		case "*":
			sum *= calculateSingleDigit(array[0 : mid-1])
		case "/":
			sum /= calculateSingleDigit(array[0 : mid-1])
		}
	}

	return sum
}
