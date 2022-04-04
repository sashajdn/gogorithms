package binarytree

import "math"

// MaxPathWithNegatives ...
//
// T -> O(n) where `n` is the number of nodes in the binary tree.
// S -> O(H) where `H` is the height of the binary tree.
func MaxPathWithNegatives(root *BinaryTree) int {
	var maxPath = math.MinInt

	var find func(node *BinaryTree) int
	find = func(node *BinaryTree) int {
		if node == nil {
			return 0
		}

		left := max(find(node.Left), 0)
		right := max(find(node.Right), 0)

		maxPath = max(maxPath, node.Value+left+right)

		return node.Value + max(left, right)
	}

	find(root)

	return maxPath
}
