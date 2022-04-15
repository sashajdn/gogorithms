package binarytree

// LevelOrderTraversal ...
//
// T -> O(n)
// S -> O(1)
func LevelOrderTraversal(root *BinaryTree) [][]int {
	var output = [][]int{}
	if root == nil {
		return output
	}

	var (
		queue = []*BinaryTree{root}
		level = 1
	)
	for len(queue) > 0 {
		l := len(queue)

		for i := 0; i < l; i++ {
			var next *BinaryTree
			next, queue = queue[0], queue[1:] // This is O(n) is golang; we should use a linked list or queue structure.

			if len(output) < level {
				output = append(output, []int{})
			}
			output[level-1] = append(output[level-1], next.Value)

			if next.Left != nil {
				queue = append(queue, next.Left)
			}
			if next.Right != nil {
				queue = append(queue, next.Right)
			}
		}

		level++
	}

	return output
}
