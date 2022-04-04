package graphs

// MaxProbability ...
//
// T -> O(v + e)
// S -> O(max(e, v))
func MaxProbability(numberOfNodes int, edges [][]int, successfulProbabilities []float64, start int, end int) float64 {
	var graph = map[int][][]int{}

	// T -> O(e)
	// S -> O(v)
	for i, edge := range edges {
		from, to := edge[0], edge[1]

		if _, ok := graph[from]; !ok {
			graph[from] = [][]int{}
		}
		graph[from] = append(graph[from], []int{to, i})

		if _, ok := graph[to]; !ok {
			graph[to] = [][]int{}
		}
		graph[to] = append(graph[to], []int{to, i})
	}

	// S -> O(v)
	var probabilities = make([]float64, numberOfNodes)
	probabilities[start] = 1

	// T -> O(v + e)
	// S -> O(e)
	var q = []int{start}
	for len(q) > 0 {
		var from int
		from, q = q[0], q[1:] // NOTE: this is a T -> O(1) operation only if we are using a double ended queue e.g linked list.

		for _, edge := range graph[from] {
			to, index := edge[0], edge[1]
			cost := successfulProbabilities[index]

			if cost*probabilities[from] > probabilities[to] {
				probabilities[to] = cost * probabilities[from]
				q = append(q, to)
			}
		}
	}

	return probabilities[end]
}
