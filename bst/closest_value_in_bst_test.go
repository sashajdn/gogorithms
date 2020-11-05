package bst

import (
	"testing"
)

func TestClosestValueInBst(t *testing.T) {
	lbst2 := BST{
		value: 2,
	}

	rbst2 := BST{
		value: 6,
	}

	lbst1 := BST{
		value: 5,
		left:  &lbst2,
		right: &rbst2,
	}

	rlbst2 := BST{
		value: 8,
	}

	rbst1 := BST{
		value: 11,
		left:  &rlbst2,
	}

	lbst0 := BST{
		value: 10,
		left:  &lbst1,
		right: &rbst1,
	}

	rlbst1 := BST{
		value: 19,
	}

	rrbst1 := BST{
		value: 25,
	}

	rbst0 := BST{
		value: 24,
		left:  &rlbst1,
		right: &rrbst1,
	}

	bst := BST{
		value: 13,
		left:  &lbst0,
		right: &rbst0,
	}

	target := 5
	found := bst.FindClosestvalue(target)

	if found != target {
		t.Errorf("expected %d got %d", target, found)
	}

	target = 19
	found = bst.FindClosestvalue(target)

	if found != target {
		t.Errorf("expected %d got %d", target, found)
	}

	target = 9
	found = bst.FindClosestvalue(target)

	diff := abs(target - found)

	if diff != 1 {
		t.Errorf("expected %d got %d", 1, diff)
	}

	target = 27
	found = bst.FindClosestvalue(target)

	diff = abs(target - found)

	if diff != 2 {
		t.Errorf("expected %d got %d", 2, diff)
	}
}
