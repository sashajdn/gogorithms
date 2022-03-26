package dynamic

// MaxSumIncreasingSubsequence ...
//
// T -> O(n ** 2) where `n` is the number of elements in the array.
//   		  Here, since we have to iteratively check all the values for j < i, where
//                in the worst case i == n, we have a time complexity of O(n ** 2)
// S -> O(n) since we have to build a adjacency array to maximize sums & to store the last increasing
//           sequence index.
func MaxSumIncreasingSubsequence(array []int) (int, []int) {
	if len(array) == 0 {
		return 0, []int{}
	}

	var (
		sums                  []int
		maxIncreasingSequence []int
	)
	for i := 0; i < len(array); i++ {
		sums[i] = array[i]
		maxIncreasingSequence[i] = -1
	}

	var maxSumIndex int
	for i := 0; i < len(array); i++ {
		currentNumber := array[i]
		for j := 0; j < i; j++ {
			if currentNumber <= array[j] {
				continue
			}

			if sums[j]+currentNumber > sums[i] {
				sums[i] = sums[j] + currentNumber
				maxIncreasingSequence[i] = j
			}
		}

		if sums[i] > sums[maxSumIndex] {
			maxSumIndex = i
		}
	}

	return sums[maxSumIndex], buildSequence(array, maxIncreasingSequence, maxSumIndex)
}

func buildSequence(array []int, sequence []int, index int) []int {
	var increasingSequence []int
	for index > 0 {
		increasingSequence = append(increasingSequence, array[index])
		index = sequence[index]
	}

	i, j := 0, len(increasingSequence)-1
	for i < j {
		increasingSequence[i], increasingSequence[j] = increasingSequence[j], increasingSequence[i]
	}

	return increasingSequence
}
