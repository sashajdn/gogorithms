package arrays

// BinarySearchIndex ...
//
// T -> O(log(n)) where n is the totat number of items in prefix values.
// S -> O(log(n)) (at most log(n) recursive calls on stack at any given time)
func BinarySearchIndexRecursive(prefixValues []int, target int) int {
	return binarySearchIndexRecursive(prefixValues, target, 0, len(prefixValues)-1)
}

func binarySearchIndexRecursive(prefixValues []int, target, l, r int) int {
	if l >= r {
		return l
	}

	mid := l + ((r - l) / 2)

	switch {
	case target > prefixValues[mid]:
		return binarySearchIndexRecursive(prefixValues, target, mid+1, r)
	default:
		return binarySearchIndexRecursive(prefixValues, target, l, mid)
	}
}

// BinarySearchIndexIterative ...
//
// T -> O(log(n)) where n is the number of values in the array.
// S -> O(1)
func BinarySearchIndexIterative(prefixValues []int, target int) int {
	l, r := 0, len(prefixValues)-1

	for r > l {
		mid := l + ((r - l) / 2)

		switch {
		case target > prefixValues[mid]:
			l = mid + 1
		default:
			r = mid
		}
	}

	return l
}
