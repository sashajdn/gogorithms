package binarytree

// BuildTreePreOrderInOrder ...
//
// T -> O(n) where `n` is the number of nodes in the tree.
// S -> O(h) where `h` is the height of the tree.
func BuildTreePreOrderInOrder(preorder, inorder []int) *BinaryTree {
	var index = map[int]int{}
	for i := 0; i < len(inorder); i++ {
		index[inorder[i]] = i
	}

	var (
		construct     func(preorder []int, left, right int) *BinaryTree
		preOrderIndex int
	)
	construct = func(preorder []int, left, right int) *BinaryTree {
		if left > right {
			return nil
		}

		rootValue := preorder[preOrderIndex]
		preOrderIndex++

		root := &BinaryTree{
			Value: rootValue,
		}

		root.Left = construct(preorder, left, index[rootValue]-1)
		root.Right = construct(preorder, index[rootValue]+1, right)

		return root
	}

	return construct(preorder, 0, len(preorder)-1)
}
