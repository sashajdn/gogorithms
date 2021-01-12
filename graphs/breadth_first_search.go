package graphs

// BreadthFirstSearch : time -> O(v + e), space O(v)
func (g *Graph) BreadthFirstSearch(array []string) []string {
	q := []*Graph{g}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		array = append(array, current.id)
		q = append(q, current.children...)
	}
	return array
}
