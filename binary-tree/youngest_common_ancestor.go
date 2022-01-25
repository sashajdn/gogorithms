package binarytree

// AncestralTree ...
type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

// GetYoungestCommonAncestor given three inputs of an `AncestralTree`, where `top` is the top
// ancestor of the tree (has no ancestor) & the other inputs being descendants, this returns
// the youngest common ancestor.
// Here a descendant can be considered it's own ancestor.
//
// T -> O(d)
// S -> O(1)
func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	// Get tree depth of both: O(d)
	oneDepth, twoDepth := getTreeDepth(top, descendantOne), getTreeDepth(top, descendantTwo)

	// Recurse either until both same depth: O(d)
	var one, two = descendantOne, descendantTwo
	switch {
	case oneDepth > twoDepth:
		one = recurse(descendantOne, oneDepth-twoDepth)

		if one == two {
			return two
		}
	case twoDepth > oneDepth:
		two = recurse(descendantTwo, twoDepth-oneDepth)

		if one == two {
			return one
		}
	}

	// Recurse both in lockstep to find common ancestor: O(d)
	return findAncestor(top, one, two)
}

func findAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	if descendantOne == top || descendantTwo == top {
		return top
	}

	if descendantOne.Ancestor == descendantTwo.Ancestor {
		return descendantOne.Ancestor
	}

	return findAncestor(top, descendantOne.Ancestor, descendantTwo.Ancestor)
}

func recurse(node *AncestralTree, howMany int) *AncestralTree {
	if howMany == 0 {
		return node
	}
	if node.Ancestor == nil {
		return node
	}

	return recurse(node.Ancestor, howMany-1)
}

func getTreeDepth(top, d *AncestralTree) int {
	if d == top {
		return 0
	}

	return 1 + getTreeDepth(top, d.Ancestor)
}
