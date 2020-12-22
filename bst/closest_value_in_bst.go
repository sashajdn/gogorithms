package bst

// FindClosestValue T -> O(log(n)), S -> O(1)
func (tree *BST) FindClosestvalue(target int) int {
	if tree.value == target {
		return target
	}

	diff := abs(tree.value - target)

	if tree.left != nil {
		if target < tree.value {
			ldiff := abs(tree.left.value - target)

			if diff < ldiff {
				return tree.value
			}

			return tree.left.FindClosestvalue(target)
		}
	}

	if tree.right != nil {
		rdiff := abs(tree.right.value - target)

		if diff < rdiff {
			return tree.value
		}

		return tree.right.FindClosestvalue(target)
	}

	return tree.value
}

func abs(value int) int {
	if value > 0 {
		return value
	}
	return value * -1
}
