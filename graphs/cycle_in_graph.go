package graphs

// CycleInGraph given a slice of slices of integers representing edges in a unweighted, directed graph,
// with at least one node. Returns a boolean representing if the graph contains a cycle.
// A cycle is defined as any number of vertices, including just one vertex, that are
// connected in a close chain.
//
// T -> O(v+e)
// S -> O(v)
func CycleInGraph(edges [][]int) bool {
	var (
		visited          = make([]bool, len(edges))
		currentlyInStack = make([]bool, len(edges))
	)
	for i := range edges {
		if visited[i] {
			continue
		}

		if cycleFound := depthfs(edges, i, visited, currentlyInStack); cycleFound {
			return true
		}

	}

	return false
}

func depthfs(edges [][]int, node int, visited []bool, currentlyInStack []bool) bool {
	visited[node] = true
	currentlyInStack[node] = true

	for _, e := range edges[node] {
		switch {
		case !visited[e]:
			if cycleFound := depthfs(edges, e, visited, currentlyInStack); cycleFound {
				return true
			}
		case currentlyInStack[e]:
			return true
		}
	}

	currentlyInStack[node] = false
	return false
}
