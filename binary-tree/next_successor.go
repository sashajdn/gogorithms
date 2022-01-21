package binarytree

// BinaryTreeParent ...
type BinaryTreeParent struct {
	Value  int
	Left   *BinaryTreeParent
	Right  *BinaryTreeParent
	Parent *BinaryTreeParent
}

// FindSuccessor given a tree & a node in that tree, returns the next successor assuming
// in order traversal.
//
// T -> O(h)
// S -> O(1)
func FindSuccessor(tree *BinaryTreeParent, node *BinaryTreeParent) *BinaryTreeParent {
	if node.Right != nil {
		// Find left most child.
		return leftMostChild(node.Right)
	}

	// Find right most parent.
	return rightMostParent(node)
}

func leftMostChild(node *BinaryTreeParent) *BinaryTreeParent {
	if node.Left == nil {
		return node
	}

	return leftMostChild(node.Left)
}

func rightMostParent(node *BinaryTreeParent) *BinaryTreeParent {
	if node.Parent == nil {
		return nil
	}

	if node.Parent.Right != node {
		return node.Parent
	}

	return rightMostParent(node.Parent)
}
