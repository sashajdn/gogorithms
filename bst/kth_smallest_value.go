package bst

// KthSmallest ...
//
// T -> O(n)
// S -> O(1)
func KthSmallest(root *BinaryTree, k int) int {
	var (
		count   = 1
		current = root
	)
	for current != nil {
		if current.Left == nil {
			if count == k || current.Right == nil {
				return current.Value
			}

			count++
			current = current.Right
			continue
		}

		var rightmost = current.Left
		for rightmost.Right != nil && rightmost.Right != current {
			rightmost = rightmost.Right
		}

		if rightmost.Right == nil {
			rightmost.Right = current
			current = current.Left
			continue
		}

		if count == k {
			return current.Value
		}

		count++
		rightmost.Right = nil
		current = current.Right
	}

	return -1
}
