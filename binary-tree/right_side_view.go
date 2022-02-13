package binarytree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// RightSideView given a root node; return the right most node from each level in an ordered array.
//
// T -> O(n) where n is the number of nodes in a tree, uses BFS which typically is O(v + e), but since e is at bounded for all nodes, gives us O(v=n)
// S -> O(n)
func RightSideView(root *TreeNode) []int {
	var collection []int
	if root == nil {
		return collection
	}

	var queue = []*TreeNode{root}
	collectRightSideView(&queue, &collection)

	return collection
}

func collectRightSideView(queue *[]*TreeNode, collection *[]int) {
	var (
		rightMostNode *TreeNode
		length        = len(*queue)
	)
	for i := 0; i < length; i++ {
		next := (*queue)[0]
		*queue = (*queue)[1:]

		if next.Left != nil {
			*queue = append(*queue, next.Left)
		}

		if next.Right != nil {
			*queue = append(*queue, next.Right)
		}

		rightMostNode = next
	}

	if rightMostNode == nil {
		return
	}
	*collection = append(*collection, rightMostNode.Val)

	collectRightSideView(queue, collection)
}
