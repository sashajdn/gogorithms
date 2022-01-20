package binarytree

// BinaryTree ...
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// BinaryTreeDiameter returns the diameter of a binary tree, where the diameter is defined as the length
// of the longest path, even if the path doesn't pass through the root of the tree.
// A path is a collection of connected nodes in a tree, where no node is connected to
// to more than two other nodes.
// The length of the path is the length between the paths first node & last node.
//
// O(n) => T(H**2) where H is the height of the tree.
// O(n) => S(H) for H recursive calls on the stack.
func BinaryTreeDiameter(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}

	if tree.Left != nil && tree.Right != nil {
		return max(recurseAndSum(tree), max(recurseAndSum(tree.Left), recurseAndSum(tree.Right)))
	}

	if tree.Left == nil {
		return max(recurseAndSum(tree), recurseAndSum(tree.Right))
	}

	if tree.Right == nil {
		return max(recurseAndSum(tree), recurseAndSum(tree.Left))
	}

	return -1
}

func recurseAndSum(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}

	if tree.Left != nil && tree.Right != nil {
		return 1
	}

	if tree.Left != nil {
		return 1 + recurseAndSum(tree.Right)
	}

	if tree.Right != nil {
		return 1 + recurseAndSum(tree.Left)
	}

	return 1 + recurseAndSum(tree.Left) + recurseAndSum(tree.Right)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
