package sorting

// MergeSort : T -> O(nlogn), S -> O(nlogn)
func MergeSort(array []int) []int{
	if len(array) < 2 {
		return array
	}

	return merge(
		make([]int, 0),
		MergeSort(array[:len(array) / 2]),
		MergeSort(array[len(array) / 2:]),
	)
}

// merge: T -> O(n), S -> O(nlogn)
func merge(arr []int, a []int, b []int) []int {
	if len(a) == 0 {
		arr = append(arr, b...)
		return arr
	}
	if len(b) == 0 {
		arr = append(arr, a...)
		return arr
	}

	if a[0] < b[0] {
		arr = append(arr, a[0])
		return merge(arr, a[1:], b)
	}

	arr = append(arr, b[0])
	return merge(arr, a, b[1:])
}
