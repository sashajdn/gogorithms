package bst

// MinHeightBST accepts a sorted array and returns a BST under a min height constraint
func MinHeightBST(array []int) *BST {
	if len(array) == 0 {
		return nil
	}
	t := &BST{
		value: array[midIndex(array)],
	}
	if len(array) == 1 {
		return t
	}
	t.left = MinHeightBST(array[:midIndex(array)])
	t.right = MinHeightBST(array[midIndex(array)+1:])
	return t
}

func midIndex(array []int) int {
	if len(array) == 1 {
		return 0
	}
	return (len(array) / 2)
}
