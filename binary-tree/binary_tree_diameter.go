package binarytree

// BinaryTree ...
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

type treeInfo struct {
	Height   int
	Diameter int
}

// BinaryTreeDiameter returns the diameter of a binary tree, where the diameter is defined as the length
// of the longest path, even if the path doesn't pass through the root of the tree.
// A path is a collection of connected nodes in a tree, where no node is connected to
// to more than two other nodes.
// The length of the path is the length between the paths first node & last node.
//
// O(n) => T(n) - due to the depth first search.
// O(n) => S(h lim-> n) - average case is h, height of the tree, but tends to n, as the tree becomes more unbalanced.
func BinaryTreeDiameter(tree *BinaryTree) int {
	return getTreeInfo(tree).Diameter
}

func getTreeInfo(tree *BinaryTree) *treeInfo {
	if tree == nil {
		return &treeInfo{}
	}

	// Depth first search.
	lhs := getTreeInfo(tree.Left)
	rhs := getTreeInfo(tree.Right)

	// Calculate the max diameter so far.
	maxDiameterSoFar := max(
		lhs.Diameter,
		rhs.Diameter,
		lhs.Height+rhs.Height,
	)

	// Calculate the current height.
	currentHeight := max(lhs.Height, rhs.Height) + 1

	return &treeInfo{
		Diameter: maxDiameterSoFar,
		Height:   currentHeight,
	}
}

func max(vv ...int) int {
	var max int
	for _, v := range vv {
		if v > max {
			max = v
		}
	}

	return max
}
