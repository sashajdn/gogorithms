package graphs

// Graph
type Graph struct {
	id       string
	children []*Graph
}

// DepthFirstSearch
func (g *Graph) DepthFirstSearch(a []string) []string {
	a = append(a, g.id)
	for _, child := range g.children {
		a = child.DepthFirstSearch(a)
	}
	return a
}
