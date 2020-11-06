package bst

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// BranchSums T -> O(n), S -> O(log(n))
func BranchSums(root *BinaryTree) []int {
	sums := make([]int, 0)
	branchSums(root, 0, &sums)
	return sums
}

func branchSums(root *BinaryTree, currentValue int, sums *[]int) {
	currentValue += root.Value

	if root.Left != nil {
		branchSums(root.Left, currentValue, sums)
	}
	if root.Right != nil {
		branchSums(root.Right, currentValue, sums)
	}

	if root.Right == nil && root.Left == nil {
		*sums = append(*sums, currentValue)
	}
}
