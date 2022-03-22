package binarytree

// MaxDepthDFS ...
//
// T -> O(n) where `n` is the number of nodes in the binary tree.
// S -> O(n) since we can have in the worst `n` recursive calls on our stack.
func MaxDepthDFS(root *BinaryTree) int {
	if root == nil {
		return 0
	}

	return max(MaxDepthDFS(root.Left)+1, MaxDepthDFS(root.Right)+1)
}

// MaxDepthBFS ...

// T -> O(n) where `n` is the number of nodes in the binary tree.
// S -> O(n) Worst case when the tree is a linkedlist & O(log(n)) Best case.
func MaxDepthBFS(root *BinaryTree) int {
	if root == nil {
		return 0
	}

	var (
		maxDepth int
		q        = []*BinaryTree{root}
	)

	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			var next *BinaryTree
			next, q = q[0], q[1:]

			if next.Left != nil {
				q = append(q, next.Left)
			}
			if next.Right != nil {
				q = append(q, next.Right)
			}
		}
		maxDepth++
	}

	return maxDepth
}
