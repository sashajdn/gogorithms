package bst

type treeInfo struct {
	height     int
	isBalanced bool
}

// HeightBalancedBinaryTree T(N) -> O(n), S(N) -> O(h)
func HeightBalancedBinaryTree(tree *BST) bool {
	treeInfo := getTreeInfo(tree)
	return treeInfo.isBalanced
}

func getTreeInfo(node *BST) *treeInfo {
	if node == nil {
		return &treeInfo{
			height:     -1,
			isBalanced: true,
		}
	}
	l, r := getTreeInfo(node.left), getTreeInfo(node.right)
	return &treeInfo{
		isBalanced: and(l.isBalanced, r.isBalanced, abs(l.height-r.height) <= 2),
		height:     max(l.height, r.height) + 1,
	}
}

func and(conditions ...bool) bool {
	for _, c := range conditions {
		if !c {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
