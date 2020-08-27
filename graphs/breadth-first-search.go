package graphs

// Node : graph node
type Node struct {
	Name string
	Children []*Node
}

// BreadthFirstSearch : time -> O(v + e), space O(v + e)
func (n *Node )BreadthFirstSearch(array []string) []string {
	q := []*Node{n}
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		array = append(array, q.Name)
		for _, child := range q.Children {
			q = append(q, child)
		}
	}
	return array
}
