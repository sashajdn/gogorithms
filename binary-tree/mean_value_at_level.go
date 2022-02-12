package binarytree

type Node struct {
	Value       int
	Left, Right *Node
}

// Find the average value at each level in a Binary Tree & return as an ordered array.
func AverageValueAtLevel(root *Node) []int {
	var (
		averageValues []int
		queue         = []*Node{root}
	)
	bfs(0, &averageValues, &queue)
	return averageValues
}

// T -> O(n) - where n is the number of nodes in the tree, since we have to visit every node
// S -> O(n) - since if tree is degenerate & tends to a linked list, we will have a recursive stack on n.
//             if tree perfect, then we have O(d), where d is the depth.
func bfs(level int, collection *[]int, queue *[]*Node) {
	var (
		sum int
		l   = len(*queue)
	)
	for i := 0; i < l; i++ {
		next := (*queue)[0]
		*queue = (*queue)[1:]

		var m = 1
		if i != 0 {
			m = i
		}

		sum *= m
		sum += next.Value
		sum /= i + 1

		if next.Left != nil {
			*queue = append(*queue, next.Left)
		}
		if next.Right != nil {
			*queue = append(*queue, next.Right)
		}

	}

	*collection = append(*collection, sum)

	if len(*queue) > 0 {
		bfs(level+1, collection, queue)
	}

	return
}
