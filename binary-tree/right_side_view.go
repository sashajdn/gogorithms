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

// RightSideViewDFS ...
//
// T -> O(n) where `n` is the total number of nodes in the binary tree.
// S -> O(log(n)) as we shall have in the worst case at most log(n) recursive calls on the stack as this is the depth of the tree.
func RightSideViewDFS(root *TreeNode) []int {
	var collection []int
	rightSideViewDFS(root, 0, &collection)
	return collection
}

func rightSideViewDFS(node *TreeNode, level int, collector *[]int) {
	if node == nil {
		return
	}

	switch {
	case len(*collector) >= level+1:
		(*collector)[level] = node.Val
	default:
		*collector = append(*collector, node.Val)
	}

	rightSideViewDFS(node.Left, level+1, collector)
	rightSideViewDFS(node.Right, level+1, collector) // Last write wins
}
