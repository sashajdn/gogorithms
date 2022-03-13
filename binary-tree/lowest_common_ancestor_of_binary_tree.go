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

// LowestCommonAncestorOfBinaryTreeParent ...
//
// T -> O(n)
// S -> O(n)
func LowestCommonAncestorOfBinaryTreeParent_Recursive(p, q *BinaryTreeParent) *BinaryTreeParent {
	return search(p, q, p, q)
}

func LowestCommonAncestorOfBinaryTreeParent_Iterative(p, q *BinaryTreeParent) *BinaryTreeParent {
	p1, p2 := p, q
	for p1 != p2 {
		switch p1.Parent {
		case nil:
			p1 = q
		default:
			p1 = p1.Parent
		}

		switch p2.Parent {
		case nil:
			p2 = p
		default:
			p2 = p2.Parent
		}
	}

	return p1
}

func search(p, q, p1, p2 *BinaryTreeParent) *BinaryTreeParent {
	if p1 == p2 {
		return p1
	}

	if p1.Parent != nil && p2.Parent != nil {
		return search(p, q, p1.Parent, p2.Parent)
	}

	if p1.Parent == nil {
		return search(p, q, q, p2.Parent)
	}

	return search(p, q, p1.Parent, q)
}
