package binarytree

// VerticalOrderTraversal ...
//
// T -> O(n)
// S -> O(n)
func VerticalOrderTraversal(root *BinaryTree) [][]int {
	if root == nil {
		return [][]int{}
	}

	type queueItem struct {
		Node  *BinaryTree
		Width int
	}

	var (
		levels = map[int][]int{}
		queue  = []*queueItem{
			{
				Node:  root,
				Width: 0,
			},
		}
	)

	var leftWidth, rightWidth int
	for len(queue) > 0 {
		var next *queueItem
		next, queue = queue[0], queue[1:]

		if _, ok := levels[next.Width]; !ok {
			levels[next.Width] = []int{}
		}
		levels[next.Width] = append(levels[next.Width], next.Node.Value)

		if next.Node.Left != nil {
			queue = append(queue, &queueItem{
				Node:  next.Node.Left,
				Width: next.Width - 1,
			})
			leftWidth = min(leftWidth, next.Width-1)
		}

		if next.Node.Right != nil {
			queue = append(queue, &queueItem{
				Node:  next.Node.Right,
				Width: next.Width + 1,
			})
			rightWidth = max(rightWidth, next.Width+1)
		}
	}

	var output = make([][]int, 0, len(levels))
	for i := leftWidth; i <= rightWidth; i++ {
		output = append(output, levels[i])
	}

	return output
}
