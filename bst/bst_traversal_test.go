package bst

import "testing"

func TestBSTTraversal(t *testing.T) {
	t.Parallel()

	// Root
	bst := &BST{
		value: 10,
	}

	// Left
	bst.left = &BST{
		value: 5,
	}
	bst.left.left = &BST{
		value: 2,
	}
	bst.left.left.left = &BST{
		value: 1,
	}

	bst.left.right = &BST{
		value: 5,
	}

	// Right
	bst.right = &BST{
		value: 15,
	}
	bst.right.right = &BST{
		value: 22,
	}

	tests := []struct {
		name           string
		expectedOutput []int
		traversalFunc  func(array []int) []int
	}{
		{
			name:           "in-order-traversal",
			expectedOutput: []int{1, 2, 5, 5, 10, 15, 22},
			traversalFunc:  bst.InOrderTraversal,
		},
		{
			name:           "pre-order-traversal",
			expectedOutput: []int{10, 5, 2, 1, 5, 15, 22},
			traversalFunc:  bst.PreOrderTraversal,
		},
		{
			name:           "post-order-traversal",
			expectedOutput: []int{1, 2, 5, 5, 22, 15, 10},
			traversalFunc:  bst.PostOrderTraversal,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := tc.traversalFunc([]int{})
			if !isSliceEqual(res, tc.expectedOutput) {
				t.Fatalf("expected -> %v, got -> %v", tc.expectedOutput, res)

			}
		})
	}
}

func isSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
