package graphs

// DepthFirstSearch
func (g *Graph) DepthFirstSearch(a []string) []string {
	a = append(a, g.id)
	for _, child := range g.children {
		a = child.DepthFirstSearch(a)
	}
	return a
}
