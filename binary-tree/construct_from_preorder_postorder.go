package binarytree

// ConstructFromInOrderPostOrder ...
//
// T -> O(n) where `n` is the number of nodes in the tree.
// S -> O(h) where `h` is the height of the binary tree.
func ConstructFromInOrderPostOrder(inorder, postorder []int) *BinaryTree {
	var index = map[int]int{}
	for i := 0; i < len(inorder); i++ {
		index[inorder[i]] = i
	}

	var (
		postOrderIndex = len(postorder) - 1
		construct      func(postorder []int, left, right int) *BinaryTree
	)
	construct = func(postorder []int, left, right int) *BinaryTree {
		if left > right {
			return nil
		}

		rootValue := postorder[postOrderIndex]
		root := &BinaryTree{
			Value: rootValue,
		}
		postOrderIndex--

		root.Right = construct(postorder, index[postOrderIndex]+1, right)
		root.Left = construct(postorder, left, index[postOrderIndex]-1)

		return root
	}

	return construct(postorder, 0, len(postorder)-1)
}
