package bst

// Invert inverts binary search tree in place
// O(T) -> O(n), O(S) -> O(logn)
func (t *BST) Invert() {
	if t.left == nil && t.right == nil {
		return
	}
	t.left, t.right = t.right, t.left
	if t.left != nil {
		t.left.Invert()
	}
	if t.right != nil {
		t.right.Invert()
	}
}
