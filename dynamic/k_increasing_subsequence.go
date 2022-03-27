package dynamic

// KIncreasingSubsequence ...
//
// T -> O(nlog(n)) where `n` is the number of elements in the array
//                 here we do (k * b(i)log(b(i))) operations, but here the sum of bi (0 - k) is equal to `n`
//                 so we are left with O(nlog(n))
// S -> O(n)
func KIncreasingSubsequence(array []int, k int) int {
	buckets := partition(array, k)

	var numberOfOperationsRequired int
	for _, bucket := range buckets {
		numberOfOperationsRequired += longestIncreasingSequenceDiff(bucket)
	}

	return numberOfOperationsRequired
}

// T -> O(n)
// S -> O(n)
func partition(array []int, k int) [][]int {
	var buckets = make([][]int, 0, k)
	for i := 0; i < k; i++ {
		buckets = append(buckets, []int{})
	}

	for i := 0; i < len(array); i++ {
		bucket := i % k
		buckets[bucket] = append(buckets[bucket], array[i])
	}

	return buckets
}

// T -> O(nlog(n))
// S -> O(n)
func longestIncreasingSequenceDiff(array []int) int {
	if len(array) == 0 {
		return 0
	}

	// T -> O(n)
	// S -> O(n)
	var sequence = []int{array[0]}
	for i := 1; i < len(array); i++ {
		currentNumber := array[i]

		if currentNumber >= sequence[len(sequence)-1] {
			sequence = append(sequence, currentNumber)
			continue
		}

		// T -> O(log(n))
		firstIncreasingIndex := bisectLeft(sequence, currentNumber)
		if sequence[firstIncreasingIndex] == currentNumber {
			firstIncreasingIndex++
		}

		sequence[firstIncreasingIndex] = currentNumber
	}

	return len(array) - len(sequence)
}
