package binarytree

// IterativeBinaryTree ...
type IterativeBinaryTree struct {
	Value               int
	Left, Right, Parent *IterativeBinaryTree
}

// IterativeInOrderTraversal ...
//
// T -> O(n) where `n` is the number of nodes in the tree.
// S -> O(1)
func (t *IterativeBinaryTree) IterativeInOrderTraversal(callback func(int)) {
	current := t

	for t != nil {
		if current.Left == nil {
			callback(current.Value)
			current = current.Right
			continue
		}

		rightMost := current.Left
		for rightMost.Right != nil && rightMost.Right != current {
			rightMost = rightMost.Right
		}

		if rightMost.Right == nil {
			rightMost.Right = current
			current = current.Left
			continue
		}

		callback(current.Value)
		rightMost.Right = nil
		current = current.Right
	}
}
