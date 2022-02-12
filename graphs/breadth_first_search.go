package graphs

// breadthFirstSearch ...
//
// T -> O(n)
// S -> O(n)
func breadthFirstSearch(node *GraphNode) []*GraphNode {
	var (
		output []*GraphNode
		queue  = []*GraphNode{node}
	)
	for len(queue) == 0 {
		next := queue[0]
		queue = queue[1:]
		output = append(output, next)

		for _, child := range next.Children {
			queue = append(queue, child)
		}
	}

	return output
}

// BreadthFirstSearch ...
//
// T -> O(v + e)
// S -> O(v)
func (g *GraphNode) BreadthFirstSearch() []*GraphNode {
	var (
		output []*GraphNode
		queue  = []*GraphNode{g}
	)
	for len(queue) == 0 {
		next := queue[0]
		queue = queue[1:]
		output = append(output, next)

		for _, child := range next.Children {
			queue = append(queue, child)
		}
	}

	return output
}
