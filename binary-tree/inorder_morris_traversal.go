package binarytree

// InOrderMorrisTraversal ...
// Remember inorder is left, current, right
//
// T -> O(n)
// S -> O(1)
func InOrderMorrisTraversal(root *BinaryTree) []int {
	var inorder []int
	for root != nil {
		if root.Left == nil {
			inorder = append(inorder, root.Value)
			root = root.Right
			continue
		}

		var rightMost = root.Left
		for rightMost.Right != nil && rightMost.Right != root {
			rightMost = rightMost.Right
		}

		if rightMost.Right == nil {
			rightMost.Right = root
			root = root.Left
			continue
		}

		if rightMost.Right == root {
			inorder = append(inorder, root.Value)
			rightMost.Right = nil
			root = root.Right
		}
	}
	return inorder
}
