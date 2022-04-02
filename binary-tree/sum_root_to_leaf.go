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
	if root == nil {
		return 0
	}

	var (
		sum           int
		currentNumber int
	)

	for root != nil {
		if root.Left != nil {
			var (
				predecessor = root.Left
				steps       = 1
			)

			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
				steps++
			}

			switch {
			case predecessor.Right == nil:
				currentNumber = (currentNumber * 10) + root.Val
				predecessor.Right = root
				root = root.Left
			default:
				if predecessor.Left == nil {
					sum += currentNumber
				}

				for i := 0; i < steps; i++ {
					currentNumber /= 10
				}

				predecessor.Right = nil
				root = root.Right
			}

			continue
		}

		currentNumber = currentNumber*10 + root.Val
		if root.Right == nil {
			sum += currentNumber
		}
		root = root.Right
	}

	return sum
}
