package binarytree

// SumRootToLeafRecursive ...
//
// T -> O(n) where `n` is the number of nodes in the tree
// S -> O(H) where `H` is the height of the tree
func SumRootToLeafRecursive(root *TreeNode) int {
	var sum int
	sumRootToLeafRecursive(root, 0, &sum)
	return sum
}

func sumRootToLeafRecursive(node *TreeNode, previousNumber int, sum *int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*sum += (10 * previousNumber) + node.Val
		return
	}

	sumRootToLeafRecursive(node.Left, (10*previousNumber)+node.Val, sum)
	sumRootToLeafRecursive(node.Right, (10*previousNumber)+node.Val, sum)
}

type StackItem struct {
	Node  *TreeNode
	Value int
}

// SumRootToLeafStack ...
//
// T -> O(n) where `n` is the number of nodes in the tree
// S -> O(H) where `H` is the height of the tree
func SumRootToLeafStack(root *TreeNode) int {
	var (
		sum   int
		stack = []*StackItem{
			{
				Node:  root,
				Value: 0,
			},
		}
	)

	for len(stack) > 0 {
		var stackItem *StackItem
		stackItem, stack = stack[len(stack)-1], stack[:len(stack)-1]

		node, currentNumber := stackItem.Node, stackItem.Value
		if node != nil {
			currentNumber = (currentNumber * 10) + node.Val
			switch {
			case node.Left == nil && node.Right == nil:
				sum += currentNumber
			default:
				stack = append(stack, &StackItem{
					Node:  node.Left,
					Value: currentNumber,
				})

				stack = append(stack, &StackItem{
					Node:  node.Right,
					Value: currentNumber,
				})
			}
		}
	}

	return sum
}

// SumRootToLeafMorrisTraversal ...
//
// T -> O(n)
// S -> O(1)
func SumRootToLeafMorrisTraversal(root *TreeNode) int {
	var sum int

	var morris func(root *TreeNode)
	morris = func(root *TreeNode) {
		var currentNumber int

		for root != nil {
			if root.Left == nil {
				currentNumber = currentNumber*10 + root.Val
				if root.Right == nil {
					sum += currentNumber
				}
				root = root.Right
				continue
			}

			var currentSteps = 1
			rightMost := root.Left
			for rightMost.Right != nil && rightMost.Right == root {
				rightMost = rightMost.Right
				currentSteps++
			}

			if rightMost == nil {
				currentNumber = currentNumber*10 + root.Val
				rightMost.Right = root
				root = root.Left
				continue
			}

			if root.Left == nil {
				sum += currentNumber
			}

			for i := 0; i < currentSteps; i++ {
				currentNumber /= 10
			}

			rightMost.Right = nil
			root = root.Right
		}
	}

	morris(root)
	return sum
}
