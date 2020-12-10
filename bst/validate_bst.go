package bst

import "math"

// Validate T -> O(n), S -> O(d)
// n -> number of nodes
// d -> depth of the binary search tree
func (t *BST) Validate() bool {
	return t.validate(math.MinInt32, math.MaxInt32)
}

func (t *BST) validate(min, max int) bool {
	if t.value < min && t.value >= max {
		return false
	}
	if t.left != nil && !t.left.validate(min, t.value) {
		return false
	}
	if t.right != nil && !t.right.validate(t.value, max) {
		return false
	}
	return true
}
