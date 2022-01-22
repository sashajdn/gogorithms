package binarytree

type treeInfoHeight struct {
	Height     int
	IsBalanced bool
}

// HeightBalancedBinaryTree given a tree, returns a boolean to indicate if the tree is indeed balanced.
// A balanced tree is defined as a tree whereby the heights of the left & right subtrees are within a
// distance of 1.
//
// T :-> O(n) - Since depth first search we must check all nodes.
// S :-> O(h) - Since we will have a `h` number of recursive function calls on the stack, in the average
//              case. This tends to `n` when the tree moves to degeneracy (linked list).
//              Where `n` is the number of nodes & `h` is the height of the tree.
func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	if tree == nil {
		return true
	}

	lhs, rhs := getHeight(tree.Left), getHeight(tree.Right)
	return isBalanced(lhs, rhs)
}

func getHeight(tree *BinaryTree) *treeInfoHeight {
	if tree == nil {
		return &treeInfoHeight{
			IsBalanced: true,
		}
	}

	lhs := getHeight(tree.Left)
	rhs := getHeight(tree.Right)

	return &treeInfoHeight{
		Height:     max(lhs.Height, rhs.Height) + 1,
		IsBalanced: isBalanced(lhs, rhs),
	}
}

func maxHeight(h1, h2 int) int {
	if h1 > h2 {
		return h1
	}
	return h2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isBalanced(a, b *treeInfoHeight) bool {
	if !a.IsBalanced || !b.IsBalanced {
		return false
	}

	if abs(a.Height-b.Height) > 1 {
		return false
	}

	return true
}
