package binarytree

// SplitArray ...
//
// T -> O()
// S -> O(n)
func SplitArray(numbers []int, m int) int {
	var sum, max int
	for _, number := range numbers {
		if number > max {
			max = number
			sum += number
		}
	}

	// T -> O(log(n))
	var left, right = max, sum
	for left < right {
		mid := left + (right-left)/2

		// T -> O(n) where `n` is the length of numbers.
		if feasible(numbers, mid, m) {
			right = mid
			continue
		}

		left = mid + 1
	}

	return left
}

func feasible(numbers []int, threshold, m int) bool {
	var (
		total          int
		currentPacking = 1
	)

	for _, num := range numbers {
		if num+total > threshold {
			currentPacking++
			if currentPacking > m {
				return false
			}

			total = 0
		}

		total += num
	}

	return true
}
