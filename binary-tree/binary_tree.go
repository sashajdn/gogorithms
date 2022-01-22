package binarytree

// BinaryTree ...
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// BinaryTreeParent ...
type BinaryTreeParent struct {
	Value  int
	Left   *BinaryTreeParent
	Right  *BinaryTreeParent
	Parent *BinaryTreeParent
}
