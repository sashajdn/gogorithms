package binarytree

// PreOrderMorrisTraversal ...
// Remember pre order is current, left, right
//
// T -> O(n)
// S -> O(1)
func PreOrderMorrisTraversal(root *BinaryTree) []int {
	var preorder []int

	for root != nil {
		if root.Left == nil {
			preorder = append(preorder, root.Value)
			root = root.Right
			continue
		}

		rightMost := root.Left
		for rightMost.Right != nil && rightMost.Right != root {
			rightMost = rightMost.Right
		}

		if rightMost.Right == nil {
			rightMost.Right = root
			preorder = append(preorder, root.Value)
			root = root.Left
			continue
		}

		rightMost.Right = nil
		root = root.Right
	}

	return preorder
}
