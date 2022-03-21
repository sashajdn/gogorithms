package binarytree

// MergeBinaryTree ...
//
// T -> O(min(n, m)) where `n` is the number of nodes in the tree for `root1`, and `m` for `root2`.
// S -> O(min(n, m))
func MergeBinaryTree(root1, root2 *Node) *Node {
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	merge(root1, root2)

	return root1
}

func merge(nodeA, nodeB *Node) {
	if nodeA == nil || nodeB == nil {
		return
	}

	nodeA.Value += nodeB.Value

	switch {
	case nodeA.Left != nil && nodeB.Left != nil:
		merge(nodeA.Left, nodeB.Left)
	case nodeA == nil:
		nodeA.Left = nodeB.Left
	}

	switch {
	case nodeA.Right != nil && nodeB.Right != nil:
		merge(nodeA.Right, nodeB.Right)
	case nodeA == nil:
		nodeA.Right = nodeB.Right
	}
}
