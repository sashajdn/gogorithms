package bst

import (
	"sort"
	"testing"
)

func TestBranchSums(t *testing.T) {
	// build binary tree
	lbst2 := BinaryTree{
		Value: 2,
	}

	rbst2 := BinaryTree{
		Value: 6,
	}

	lbst1 := BinaryTree{
		Value: 5,
		Left:  &lbst2,
		Right: &rbst2,
	}

	rlbst2 := BinaryTree{
		Value: 8,
	}

	rbst1 := BinaryTree{
		Value: 11,
		Left:  &rlbst2,
	}

	lbst0 := BinaryTree{
		Value: 10,
		Left:  &lbst1,
		Right: &rbst1,
	}

	rlbst1 := BinaryTree{
		Value: 19,
	}

	rrbst1 := BinaryTree{
		Value: 25,
	}

	rbst0 := BinaryTree{
		Value: 24,
		Left:  &rlbst1,
		Right: &rrbst1,
	}

	bt := BinaryTree{
		Value: 13,
		Left:  &lbst0,
		Right: &rbst0,
	}

	// compare
	expected := []int{30, 34, 42, 56, 62}
	result := BranchSums(&bt)

	if !arraysEqual(expected, result) {
		t.Fatalf("expected -> %v got -> %v", expected, result)
	}

}

func arraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i, val := range a {
		if b[i] != val {
			return false
		}
	}
	return true
}
