package graphs

import "math"

// DirectedNode ...
type DirectedNode int

// DirectedEdge ...
type DirectedEdge struct {
	To, From DirectedNode
	Cost     int
}

// DirectedGraph ...
type DirectedGraph []*DirectedEdge

// BellmanFordGraph finds the shortest distance from a node to all other nodes; and sets nodes that are part of
// a negative cycle as minus infinity.
//
// T -> O(v * e)
// S -> O(v)
func BellmanFordGraph(startNode int, graph DirectedGraph) []int {
	var nodesSet = map[DirectedNode]struct{}{}
	for _, edge := range graph {
		if _, ok := nodesSet[edge.To]; !ok {
			nodesSet[edge.To] = struct{}{}
		}
		if _, ok := nodesSet[edge.From]; !ok {
			nodesSet[edge.From] = struct{}{}
		}
	}

	var distances = make([]int, len(nodesSet))
	for i := 0; i < len(distances); i++ {
		distances[i] = math.MaxInt
	}
	distances[startNode] = 0

	for i := 0; i < len(nodesSet); i++ {
		for _, edge := range graph {
			if distances[edge.From] == math.MaxInt {
				distances[edge.To] = edge.Cost
				continue
			}

			distances[edge.To] = min(distances[edge.To], distances[edge.From]+edge.Cost)
		}
	}

	for i := 0; i < len(nodesSet); i++ {
		for _, edge := range graph {
			if distances[edge.From]+edge.Cost < distances[edge.To] {
				distances[edge.To] = math.MinInt
			}
		}
	}

	return distances
}
