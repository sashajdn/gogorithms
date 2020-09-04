package arrays

func ValidateSubsequence(array []int, sequence[]int) bool {
	if len(sequence) == 0 {
		return true
	}

	if len(array) == 0 {
		return false
	}

	if array[0] == sequence[0] {
		return ValidateSubsequence(array[1:], sequence[1:])
	}

	return ValidateSubsequence(array[1:], sequence)
}
