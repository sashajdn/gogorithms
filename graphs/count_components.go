package graphs

type ComponentDSU struct {
	reps, ranks []int
}

func NewComponentDSU(size int) *ComponentDSU {
	var (
		ranks = make([]int, size)
		reps  = make([]int, size)
	)
	for i := 0; i < size; i++ {
		ranks[i] = 1
		reps[i] = i
	}

	return &ComponentDSU{
		ranks: ranks,
		reps:  reps,
	}
}

func (c *ComponentDSU) Find(group int) int {
	if c.reps[group] == group {
		return group
	}

	c.reps[group] = c.Find(c.reps[group])
	return c.reps[group]
}

func (c *ComponentDSU) Union(a, b int) int {
	ra, rb := c.Find(a), c.Find(b)
	if ra == rb {
		return 0
	}

	if c.ranks[ra] >= c.ranks[rb] {
		c.reps[rb] = c.reps[ra]
		c.ranks[ra] += c.ranks[rb]
		return 1
	}

	c.reps[ra] = c.reps[rb]
	c.ranks[rb] += c.ranks[ra]
	return 1
}

// CountComponents_DSU ...
//
// T -> O(E * alpha(n)) where `E` is the number of edges & `alpha` is the inverse ackermann function over the number of nodes.
// S -> O(V) where `V` is the number of vertices.
func CountComponents_DSU(n int, edges [][]int) int {
	var (
		dsu   = NewComponentDSU(n)
		count = n
	)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		count -= dsu.Union(from, to)
	}

	return count
}

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
