package bst

import (
	"testing"
)

func TestInvert(t *testing.T) {
	t.Parallel()

	input := []int{8, 4, 9, 2, 5, 1, 6, 3, 7}
	expectedOutputArray := []int{7, 3, 6, 1, 5, 2, 9, 4, 8}

	bst := MinHeightBST(input)

	bst.Invert()
	inOrderTraversal := bst.InOrderTraversal([]int{})

	if !slicesEqual(expectedOutputArray, inOrderTraversal) {
		t.Fatalf("expected -> %v, got -> %v", expectedOutputArray, inOrderTraversal)
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
