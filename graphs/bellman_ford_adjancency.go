package graphs

import "math"

// BellmanFordAdjacency is a single source shortest path algorithm, it finds the shortest path from a source node to all others.
// It can be used even when we have negative cycle weights.
//
// T -> O(e * v)
// S -> O()
func BellmanFordAdjacency(start int, edges [][][]int) []int {
	// T -> O(v)
	var distances = make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		distances[i] = math.MaxInt
	}
	distances[start] = 0

	// T -> O(ve)
	for i := 0; i < len(edges); i++ {
		for node, outboundEdges := range edges {
			for _, edge := range outboundEdges {
				destinationNode, destinationCost := edge[0], edge[1]

				if distances[node] == math.MaxInt {
					distances[destinationNode] = destinationCost
					continue
				}

				distances[destinationNode] = min(distances[destinationNode], destinationCost+distances[node])
			}
		}
	}

	// T -> O(ve)
	//
	// Run a second time to detect which nodes are part of negative cycle.
	// A negative cycle has occurred if we find can find a better path than the optimal solution.
	// We do this `v` times to propagate the negative cycle to all nodes; if we find a node that is part of a negative cycle
	// then we set to minus infinity - in this case the minimum possible number in golang.
	for i := 0; i < len(edges); i++ {
		for node, outboundEdges := range edges {
			for _, edge := range outboundEdges {
				destinationNode, destinationCost := edge[0], edge[1]

				// To account for integer overflow if we sum math.MaxInt + math.MaxInt
				if distances[node] == math.MaxInt && distances[destinationNode] == math.MaxInt {
					continue
				}

				if distances[node]+destinationCost < distances[destinationNode] {
					distances[destinationNode] = math.MinInt
				}
			}
		}
	}

	return distances
}
