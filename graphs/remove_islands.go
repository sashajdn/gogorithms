package graphs

// RemoveIslands given a two-dimensional matrix, containing only `0` and `1` representing a two
// toned image where an island is defined as any number of `1`s that are horizontally or vertically
// adjacent & don't touch the border of the image, returns the matrix with the islands removed &
// replaced with zeros.
// May mutate the original matrix.
//
// T -> O(...)
// S -> O(...)
func RemoveIslands(matrix [][]int) [][]int {
	var visited = make([][]bool, 0, len(matrix))
	for k := 0; k < len(matrix); k++ {
		visited = append(visited, make([]bool, len(matrix[0])))
	}

	for j := 0; j < len(matrix); j++ {
		for i := 0; i < len(matrix[0]); i++ {
			if matrix[j][i] == 0 {
				continue
			}

			if visited[j][i] {
				continue
			}

			var islands = [][]int{}
			dfs(matrix, i, j, visited, &islands)

			var isIsland = true
			for _, island := range islands {
				x, y := island[0], island[1]

				if x == 0 || x == len(matrix[0])-1 {
					isIsland = false
					break
				}
				if y == 0 || y == len(matrix)-1 {
					isIsland = false
					break
				}
			}

			if !isIsland {
				continue
			}

			for _, island := range islands {
				x, y := island[0], island[1]
				matrix[y][x] = 0
			}
		}
	}

	return matrix
}

func neighbours(matrix [][]int, i, j int) [][]int {
	var n [][]int

	// Rows.
	if i > 0 {
		n = append(n, []int{i - 1, j})
	}
	if i < len(matrix[0])-1 {
		n = append(n, []int{i + 1, j})
	}

	// Columns.
	if j > 0 {
		n = append(n, []int{i, j - 1})
	}
	if j < len(matrix)-1 {
		n = append(n, []int{i, j + 1})
	}

	return n
}

func dfs(matrix [][]int, i, j int, visited [][]bool, islands *[][]int) {
	if visited[j][i] {
		return
	}

	visited[j][i] = true

	if matrix[j][i] == 0 {
		return
	}

	*islands = append(*islands, []int{i, j})

	nn := neighbours(matrix, i, j)

	for _, n := range nn {
		x, y := n[0], n[1]
		dfs(matrix, x, y, visited, islands)
	}
}
