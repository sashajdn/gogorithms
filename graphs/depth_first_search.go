package graphs

// depthFirstSearch ...
//
// T -> O(v + e) where v is the number of vertices and is the number of edges.
// S -> O(v)
func depthFirstSearch(node *GraphNode) []*GraphNode {
	var output []*GraphNode
	dfsHelper(node, &output)
	return output
}

func dfsHelper(node *GraphNode, collection *[]*GraphNode) {
	*collection = append(*collection, node)
	for _, child := range node.Children {
		dfsHelper(child, collection)
	}
}

// DepthFirstSearch ...
//
// T -> O(v + e) where v is the number of vertices and is the number of edges.
// S -> O(v)
func (g *GraphNode) DepthFirstSearch() []*GraphNode {
	var output = []*GraphNode{g}
	for _, child := range g.Children {
		output = append(output, child.DepthFirstSearch()...)
	}

	return output
}
