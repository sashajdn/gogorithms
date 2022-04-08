package arrays

// MergeKSortedArrays ...
//
// T -> O(nlog(n))
// S -> O(n)
func MergeKSortedArrays(arrays [][]int) []int {
	switch len(arrays) {
	case 0:
		return []int{}
	case 1:
		return arrays[0]
	}

	var end = len(arrays) - 1
	for end > 0 {
		var left, right = 0, end
		for left < right {
			arrays[left] = mergeTwoArrays(arrays[left], arrays[right])
			left++
			right--
		}

		end /= 2
	}

	return arrays[0]
}

func mergeTwoArrays(a, b []int) []int {
	var (
		output      []int
		left, right = 0, 0
	)

	for left < len(a) && right < len(b) {
		if a[left] < b[right] {
			output = append(output, a[left])
			left++
			continue
		}

		output = append(output, b[right])
		right++
	}

	for left < len(a) {
		output = append(output, a[left])
		left++
	}
	for right < len(b) {
		output = append(output, b[right])
		right++
	}

	return output
}
