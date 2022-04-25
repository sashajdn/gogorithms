package binarytree

import "math"

// MaxLevelSum ...
//
// T -> O(n) where `n` is the number of nodes in the tree.
// S -> O(h)
func MaxLevelSum(root *BinaryTree) int {
	var (
		maxLevel     int
		maxSum       = math.MinInt
		queue        = []*BinaryTree{root}
		currentLevel = 1
	)
	for len(queue) > 0 {
		var (
			queueSize  = len(queue)
			currentSum int
		)
		for i := 0; i < queueSize; i++ {
			var next *BinaryTree
			next, queue = queue[0], queue[1:]

			currentSum += next.Value

			if next.Left != nil {
				queue = append(queue, next.Left)
			}
			if next.Right != nil {
				queue = append(queue, next.Right)
			}
		}

		if currentSum > maxSum {
			maxLevel, maxSum = currentLevel, currentSum
		}

		currentLevel++
	}

	return maxLevel
}
