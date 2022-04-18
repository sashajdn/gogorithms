package backtracking

// WordSearch ...
//
// T -> O(N * 3^L) where `N` is the number of cells in the grid & `L` that of the word.
// S -> O(L) due to the recursive stack.
func WordSearch(board [][]byte, word string) bool {
	var charSet = map[rune]struct{}{}
	for j := 0; j < len(board); j++ {
		for i := 0; i < len(board[0]); i++ {
			char := rune(board[j][i])

			if _, ok := charSet[char]; !ok {
				charSet[char] = struct{}{}
			}
		}
	}

	var visited = map[int]struct{}{}
	for j := 0; j < len(board); j++ {
		for i := 0; i < len(board[0]); i++ {
			if search(j, i, board, charSet, visited, word) {
				return true
			}
		}
	}

	return false
}

func search(j, i int, board [][]byte, charSet map[rune]struct{}, visited map[int]struct{}, word string) bool {
	switch len(word) {
	case 0:
		return true
	case 1:
		if rune(word[0]) == rune(board[j][i]) {
			return true
		}

		return false
	}

	id := j*len(board[0]) + i
	if _, ok := visited[id]; ok {
		return false
	}
	visited[id] = struct{}{}
	defer delete(visited, id)

	char, tail := rune(word[0]), word[1:]
	if _, ok := charSet[char]; !ok {
		return false
	}

	if char != rune(board[j][i]) {
		return false
	}

	for _, neighbour := range fetchWordSearchNeighbours(j, i, board) {
		dj, di := neighbour[0], neighbour[1]

		neighbourID := (dj * len(board[0])) + di
		if _, ok := visited[neighbourID]; ok {
			continue
		}

		if search(dj, di, board, charSet, visited, tail) {
			return true
		}
	}

	return false
}

func fetchWordSearchNeighbours(j, i int, board [][]byte) [][]int {
	var (
		neighbours = [][]int{}
		directions = [][]int{
			{0, 1},
			{0, -1},
			{1, 0},
			{-1, 0},
		}
	)
	for _, direction := range directions {
		dj, di := j+direction[0], i+direction[1]

		if dj < 0 || dj > len(board)-1 {
			continue
		}
		if di < 0 || di > len(board[0])-1 {
			continue
		}

		neighbours = append(neighbours, []int{dj, di})
	}

	return neighbours
}
