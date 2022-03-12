package binarytree

// LowestCommonAncestorOfBinaryTree ...
//
// T -> O(n) where `n` is the total number of leaves in the tree at any given time.
// S -> O(n) since we can have at most `n` recursive calls on the stack at any given time, worst case is if the tree is degenerate.
func LowestCommonAncestorOfBinaryTree(root, a, b *BinaryTree) *BinaryTree {
	switch root {
	case a, b, nil:
		return root
	}

	left, right := LowestCommonAncestorOfBinaryTree(root.Left, a, b), LowestCommonAncestorOfBinaryTree(root.Right, a, b)
	if left == nil && right == nil {
		return nil
	}

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}
