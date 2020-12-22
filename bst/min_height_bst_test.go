package bst

import (
	"testing"
)

func TestMinHeightBST(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                string
		inputArray          []int
		expectedOutputArray []int
	}{
		{
			name:       "non-empty-array",
			inputArray: generateInputArray(5),
		},
		{
			name:       "large-non-empty-array",
			inputArray: generateInputArray(1000),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			bst := MinHeightBST(tc.inputArray)

			if !isBinaryTreeEqualToArray(bst, tc.inputArray) {
				t.Fatalf("expected -> %v, got -> %v", tc.inputArray, bst.InOrderTraversal([]int{}))
			}
		})
	}
}

func isBinaryTreeEqualToArray(t *BST, a []int) bool {
	b := t.InOrderTraversal([]int{})
	if len(b) != len(a) {
		return false
	}
	for i, v := range b {
		if v != a[i] {
			return false
		}
	}
	return true
}

func generateInputArray(howManyInts int) []int {
	r := []int{}
	for i := 0; i < howManyInts; i++ {
		r = append(r, i)
	}
	return r
}
