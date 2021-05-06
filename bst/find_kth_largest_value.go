package bst

type treeInfoReverse struct {
	numberOfNodesVisited   int
	latestVisitedNodeValue int
}

func FindKthLargestValue(node *BST, k int) int {
	treeInfo := &treeInfoReverse{0, -1}
	reverseOrderTraversal(node, k, treeInfo)
	return treeInfo.latestVisitedNodeValue
}

func reverseOrderTraversal(node *BST, k int, treeInfo *treeInfoReverse) {
	if node == nil || treeInfo.latestVisitedNodeValue >= k {
		return
	}
	reverseOrderTraversal(node.right, k, treeInfo)
	if treeInfo.numberOfNodesVisited < k {
		treeInfo.numberOfNodesVisited++
		treeInfo.latestVisitedNodeValue = node.value
		reverseOrderTraversal(node.left, k, treeInfo)
	}
}
