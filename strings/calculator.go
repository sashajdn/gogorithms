package strings

import (
	"strconv"
)

func Calculate(s string) int {
	var array []string
	for _, r := range s {
		if r != ' ' {
			array = append(array, string(r))
		}
	}

	return calculate(array)
}

func calculate(array []string) int {
	if len(array) == 1 {
		v, _ := strconv.Atoi(array[0])
		return v
	}

	mid := len(array) / 2

	switch array[mid] {
	case "+":
		return calculate(array[:mid]) + calculate(array[mid+1:])
	case "*":
		return calculate(array[:mid]) * calculate(array[mid+1:])
	case "-":
		return calculate(array[:mid]) - calculate(array[mid+1:])
	case "/":
		return calculate(array[:mid]) / calculate(array[mid+1:])
	}

	operationOrder := orderOfOperation(array, mid-1, mid+1)
	lop, rop := array[mid-1], array[mid+1]

	var sum, _ = strconv.Atoi(array[mid])
	switch operationOrder {
	case OperationOrderLeft:
		switch lop {
		case "+":
			sum += calculate(array[0 : mid-1])
		case "-":
			sum -= calculate(array[0 : mid-1])
		case "*":
			sum *= calculate(array[0 : mid-1])
		case "/":
			sum /= calculate(array[0 : mid-1])
		}
		switch rop {
		case "+":
			sum += calculate(array[mid+2:])
		case "-":
			sum -= calculate(array[mid-2:])
		case "*":
			sum *= calculate(array[mid-2:])
		case "/":
			sum /= calculate(array[mid-2:])
		}
	case OperationOrderRight:
		switch rop {
		case "+":
			sum += calculate(array[mid+2:])
		case "-":
			sum -= calculate(array[mid+2:])
		case "*":
			sum *= calculate(array[mid+2:])
		case "/":
			sum /= calculate(array[mid+2:])
		}
		switch lop {
		case "+":
			sum += calculate(array[0 : mid-1])
		case "-":
			sum -= calculate(array[0 : mid-1])
		case "*":
			sum *= calculate(array[0 : mid-1])
		case "/":
			sum /= calculate(array[0 : mid-1])
		}
	}

	return sum
}

type OperationOrder int

const (
	OperationOrderLeft OperationOrder = iota + 1
	OperationOrderRight
)

func orderOfOperation(array []string, l, r int) OperationOrder {
	switch array[r] {
	case "*", "/":
		return OperationOrderRight
	}

	switch array[l] {
	case "*", "/":
		return OperationOrderLeft
	}

	return OperationOrderRight
}
