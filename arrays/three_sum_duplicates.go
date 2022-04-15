package arrays

import "sort"

// ThreeSumDuplicates ...
//
// T -> O(nlog(n)) where `n` is the number of elements in the array
// S -> O(1) since we can use inplace sort.
func ThreeSumDuplicates(array []int) [][]int {
	sort.Ints(array)

	var (
		output  = [][]int{}
		i, j, k = 0, 1, len(array) - 1
	)
	for j < k {
		var left, right = j, k

		for left < right {
			value := array[i] + array[left] + array[right]

			switch {
			case value == 0:
				output = append(output, []int{array[i], array[left], array[right]})

				var currentValue = array[left]
				for left < right && currentValue == array[left] {
					left++
				}

			case value > 0:
				right--
			default:
				left++
			}
		}

		var currentValue = array[i]
		for i < k-1 && currentValue == array[i] {
			i++
		}
		j = i + 1
	}

	return output
}
