package binarytree

// LeftSideView ...
//
// T -> O(n) where n is the total number of nodes in the tree
// S -> O(log(n)) since we have recursed at most depth = log(n) times, meaning this is the max recursive calls we have on our stack.
func LeftSideView(node *Node) []int {
	var collection []int
	leftSideView(node, 0, &collection)
	return collection
}

func leftSideView(node *Node, level int, collection *[]int) {
	if node == nil {
		return
	}

	switch {
	case len(*collection) >= level+1:
	default:
		*collection = append(*collection, node.Value)
	}

	leftSideView(node.Left, level+1, collection)
	leftSideView(node.Right, level+1, collection)
}
