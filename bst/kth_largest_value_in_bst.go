package bst

type treeInfo struct {
	numberOfNodesVisited int
	lastVisitedNodeValue int
}

// FindKthLargestValueImBST T -> O(h + k), S -> O(h)
func FindKthLargestValueInBST(tree *BST, k int) int {
	treeInfo := treeInfo{0, -1}
	reverseOrderTraverse(tree, k, &treeInfo)
	return treeInfo.lastVisitedNodeValue
}

func reverseOrderTraverse(node *BST, k int, treeInfo *treeInfo) {
	if node == nil || treeInfo.numberOfNodesVisited >= k {
		return
	}

	reverseOrderTraverse(node.right, k, treeInfo)
	if treeInfo.numberOfNodesVisited < k {
		treeInfo.numberOfNodesVisited++
		treeInfo.lastVisitedNodeValue = node.value
		reverseOrderTraverse(node.left, k, treeInfo)
	}
}
