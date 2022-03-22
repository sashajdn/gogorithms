package binarytree

// FlattenPreOrderRecursive ...
//
// T -> O(n) where `n` is the number of nodes in the tree in the tree.
// S -> O(n) due to worst case of `n` recursive calls on the tree.
func FlattenPreOrderRecursive(root *BinaryTree) {
	flattenPreOrder(root)
}

func flattenPreOrder(node *BinaryTree) *BinaryTree {
	if node == nil {
		return nil
	}

	if node.Left == nil && node.Right == nil {
		return node
	}

	lt, rt := flattenPreOrder(node.Left), flattenPreOrder(node.Right)
	if lt != nil {
		lt.Right = node.Right
		node.Right = node.Left
		node.Left = nil
	}

	if rt != nil {
		return rt
	}

	return lt
}

// FlattenPreOrderMorris ...
//
// T -> O(n) where `n` is the number of nodes in the tree.
// S -> O(1)
func FlattenPreOrderMorris(root *BinaryTree) {
	if root == nil {
		return
	}

	var current = root
	for current != nil {
		if current.Left != nil {
			var rightmost = current.Left
			for rightmost.Right != nil {
				rightmost = rightmost.Right
			}

			rightmost.Right = current.Right
			current.Right = current.Left
			current.Left = nil
		}

		current = current.Right
	}
}
