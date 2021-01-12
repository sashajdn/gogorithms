package graphs

import "sync"

// Graph
type Graph struct {
	id       string
	children []*Graph
	sync.Locker
}

func (g *Graph) AddChild(c *Graph) {
	g.Lock()
	defer g.Unlock()
	g.children = append(g.children, c)
}

func New(id string) *Graph {
	return &Graph{
		id: id,
	}
}
