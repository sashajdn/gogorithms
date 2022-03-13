package arrays

// NextPermutation ...
//
// T -> O(n) where `n` is the total number of permutations in the array.
// S -> O(1)
func NextPermutation(array []int) {
	var firstDecreasingIndex = -1
	for i := len(array) - 1; i >= 1; i-- {
		if array[i-1] < array[i] {
			firstDecreasingIndex = i - 1
			break
		}
	}

	if firstDecreasingIndex >= 0 {
		var firstIncreasingIndex int
		for j := len(array) - 1; j > firstDecreasingIndex; j-- {
			if array[firstDecreasingIndex] < array[j] {
				firstIncreasingIndex = j
				break
			}
		}
		swapPerm(array, firstDecreasingIndex, firstIncreasingIndex)
	}
	reversePerm(array, firstDecreasingIndex+1)
}

func reversePerm(array []int, start int) {
	i, j := start, len(array)-1
	for i < j {
		swap(array, i, j)
		i++
		j--
	}
}

func swapPerm(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
