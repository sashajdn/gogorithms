package graphs

// CloneNode ...
type CloneNode struct {
	Value      int
	Neighbours []*CloneNode
}

// CloneGraphBFS ...
//
// T -> O(n) where `n` is the number of nodes in a graph.
// S -> O(max(ei)) where `ei` is the number of edges from a given vertex.
func CloneGraphBFS(node *CloneNode) *CloneNode {
	if node == nil {
		return nil
	}

	var (
		seen = map[int]*CloneNode{}
		q    = []*CloneNode{node}
	)

	seen[node.Value] = &CloneNode{
		Value: node.Value,
	}

	for len(q) > 0 {
		var next *CloneNode
		next, q = q[0], q[1:] // Note that this is an O(n) in go; we should use a linked list here for FIFO queue.

		for _, neighbour := range next.Neighbours {
			if _, ok := seen[neighbour.Value]; !ok {
				seen[neighbour.Value] = &CloneNode{
					Value: neighbour.Value,
				}

				q = append(q, neighbour)
			}

			seen[next.Value].Neighbours = append(seen[next.Value].Neighbours, seen[neighbour.Value])
		}
	}

	return seen[node.Value]
}

// CloneGraphDFS ...
//
// T -> O(v + e) where `v` is the number of nodes in the graph & `e` is the number of edges in the graph.
// S -> O(e) where `e` is the number of edges in the graph; due to the worst case recursive stack.
func CloneGraphDFS(node *CloneNode) *CloneNode {
	if node == nil {
		return nil
	}

	var visited = map[int]*CloneNode{
		node.Value: {
			Value: node.Value,
		},
	}

	clone(node, &visited)
	return visited[node.Value]
}

func clone(node *CloneNode, visited *map[int]*CloneNode) {
	if node == nil {
		return
	}

	for _, neighbour := range node.Neighbours {
		if _, ok := (*visited)[neighbour.Value]; !ok {
			(*visited)[neighbour.Value] = &CloneNode{
				Value: neighbour.Value,
			}

			clone(neighbour, visited)
		}

		(*visited)[node.Value].Neighbours = append((*visited)[node.Value].Neighbours, (*visited)[neighbour.Value])
	}
}
