package arrays

// ProductExceptSelf ...
//
// T -> O(n) where `n` is the number of items in the array of numbers.
// S -> S(1)
func ProductExceptSelf(numbers []int) []int {
	var output = make([]int, len(numbers))

	output[0] = 1
	for i := 1; i < len(numbers); i++ {
		output[i] = output[i-1] * numbers[i-1]
	}

	var r = 1
	for j := len(numbers) - 2; j >= 0; j-- {
		r *= numbers[j+1]
		output[j] *= r
	}

	return output
}
