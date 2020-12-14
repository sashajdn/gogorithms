package bst

func (t *BST) InOrderTraversal(array []int) []int {
	if t.left != nil {
		array = t.left.InOrderTraversal(array)
	}
	array = append(array, t.value)
	if t.right != nil {
		array = t.right.InOrderTraversal(array)
	}
	return array
}

func (t *BST) PreOrderTraversal(array []int) []int {
	array = append(array, t.value)
	if t.left != nil {
		array = t.left.PreOrderTraversal(array)
	}
	if t.right != nil {
		array = t.right.PreOrderTraversal(array)
	}
	return array

}

func (t *BST) PostOrderTraversal(array []int) []int {
	if t.left != nil {
		array = t.left.PostOrderTraversal(array)
	}
	if t.right != nil {
		array = t.right.PostOrderTraversal(array)
	}
	array = append(array, t.value)
	return array
}
