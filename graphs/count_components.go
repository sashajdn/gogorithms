package graphs

// CountComponents_DFS ...
//
// T -> O(v + e) // where v is the number of vertices & e is the number of edges duoe to DFS.
// S -> O(n) // where n is the number of nodes
func CountComponents_DFS(n int, edges [][]int) int {
	// T: O(n)
	// S: O(v + e)
	adj := map[int][]int{}
	for _, edge := range edges {
		fro, to := edge[0], edge[1]

		_, ok := adj[fro]
		switch {
		case ok:
			adj[fro] = append(adj[fro], to)
		default:
			adj[fro] = []int{to}
		}

		_, ok = adj[to]
		switch {
		case ok:
			adj[to] = append(adj[to], fro)
		default:
			adj[to] = []int{fro}
		}
	}

	var (
		connectedComponents int
		visited             = make([]bool, n)
	)
	// T -> O(n)
	// S -> O(n)
	for i := 0; i < len(adj); i++ {
		count := countComponents_DFS(i, adj, &visited)
		if count > 0 {
			connectedComponents++

		}
	}

	// T -> O(n)
	// S -> O(1)
	for _, v := range visited {
		if !v {
			connectedComponents++
		}
	}

	return connectedComponents
}

func countComponents_DFS(node int, nodes map[int][]int, visited *[]bool) int {
	if (*visited)[node] {
		return 0
	}

	(*visited)[node] = true

	var count = 1
	for _, neighbour := range nodes[node] {
		count += countComponents_DFS(neighbour, nodes, visited)
	}

	return count
}
