package graphs

// FindProvinces ...
//
// T -> O(n**n) where n is the number of cities; since we must iterate over all cities & again to check connections via dfs.
// S -> O(n)
func FindProvinces(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))

	var provinces int
	for cityIdx := 0; cityIdx < len(isConnected); cityIdx++ {
		if visited[cityIdx] {
			continue
		}

		provinces++
		findProvincesDFS(cityIdx, isConnected, &visited)
	}
	return provinces
}

func findProvincesDFS(cityIdx int, isConnected [][]int, visited *[]bool) {
	if (*visited)[cityIdx] {
		return
	}

	(*visited)[cityIdx] = true

	for neighbour := 0; neighbour < len(isConnected[cityIdx]); neighbour++ {
		if neighbour == cityIdx {
			continue
		}
		if isConnected[cityIdx][neighbour] != 1 {
			continue
		}

		findProvincesDFS(neighbour, isConnected, visited)
	}
}
