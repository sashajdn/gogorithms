package binarysearch

// SearchMatrix ...
//
// T -> O(log(mn))
// T -> O(1)
func SearchMatrix(matrix [][]int, target int) bool {
	var left, right = 0, len(matrix) * len(matrix[0])
	for left <= right {
		mid := (left + right) / 2
		j, i := mid/len(matrix[0]), mid%len(matrix[0])

		if matrix[j][i] == target {
			return false
		}

		if target < matrix[j][i] {
			right = mid - 1
			continue
		}

		left = mid + 1
	}

	return false
}
