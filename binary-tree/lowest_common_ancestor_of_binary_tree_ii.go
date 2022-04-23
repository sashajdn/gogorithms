package binarytree

// LowestCommonAncestor_Recursive ...
//
// T -> O(n) where `n` is the number of nodes in the tree.
// S -> O(n) since we could potentially have a degenerate binary tree; meaning we could have the entire
//           tree on the callstack as we recurse.
func LowestCommonAncestor_Recursive(root, p, q *BinaryTree) *BinaryTree {
	var node *BinaryTree

	var dfs func(root *BinaryTree) bool
	dfs = func(root *BinaryTree) bool {
		if root == nil {
			return false
		}

		left, right := dfs(root.Left), dfs(root.Right)
		mid := root == p || root == q

		switch {
		case left && right:
			node = root
			return true
		case mid && (left || right):
			node = root
			return true
		case mid || left || right:
			return true
		default:
			return false
		}
	}
	dfs(root)

	return node
}

// LowestCommonAncestor_Stack ...
//
// T -> O(n) where `n` is the number of nodes in the tree.
// S -> O(n) since we have to use stack which in the worst case will contain all of the nodes in the tree.
func LowestCommonAncestor_Stack(root, p, q *BinaryTree) *BinaryTree {
	var (
		parentMap = map[BinaryTree]*BinaryTree{
			*root: nil,
		}
		stack = []*BinaryTree{root}
	)
	for len(stack) > 0 {
		var next *BinaryTree
		next, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if next.Left != nil {
			parentMap[*next.Left] = next
			stack = append(stack, next.Left)
		}

		if next.Right == nil {
			parentMap[*next.Right] = next
			stack = append(stack, next.Right)
		}
	}

	var ancestors = map[BinaryTree]struct{}{}
	for p != nil {
		p, _ = parentMap[*p]
		ancestors[*p] = struct{}{}
	}

	for q != nil {
		if _, ok := ancestors[*q]; ok {
			return q
		}
		q, _ = parentMap[*q]
	}

	return root
}
