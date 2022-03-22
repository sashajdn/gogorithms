package binarytree

// FlattenBinaryTree ...
//
// T -> O(n) where `n` is the number of nodes in the binary tree; here we have to visit every node in order to flatten.
// S -> O(d) at most we'll have the depth of the binary tree on the stack as we recurse, worst case this will be O(n) when the tree is degenerate.
func FlattenBinaryTree(root *BinaryTree) *BinaryTree {
	if root == nil {
		return nil
	}

	var first, last *BinaryTree
	var dfs func(node *BinaryTree)
	dfs = func(node *BinaryTree) {
		if node == nil {
			return
		}

		dfs(node.Left)

		switch {
		case last != nil:
			last.Right = node
			node.Left = last
		default:
			first = node
		}

		last = node

		dfs(node.Right)
	}

	dfs(root)

	return first
}
