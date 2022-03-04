package tries

// FindWords ...
//
// T -> O
func FindWords(board [][]byte, words []string) []string {
	var positions = map[rune][][]int{}
	for j := 0; j < len(board); j++ {
		for i := 0; i < len(board[0]); i++ {
			if _, ok := positions[rune(board[j][i])]; !ok {
				positions[rune(board[j][i])] = [][]int{}
			}

			positions[rune(board[j][i])] = append(positions[rune(board[j][i])], []int{i, j})
		}
	}

	var validWords []string
	for _, word := range words {
		pp, ok := positions[rune(word[0])]
		if !ok {
			continue
		}

		var seen = make([][]bool, 0, len(board))
		for i := 0; i < len(board); i++ {
			seen = append(seen, make([]bool, len(board[0])))
		}

		for _, position := range pp {
			if dfs(board, position[0], position[1], &seen, word, 1) {
				validWords = append(validWords, word)
				break
			}
		}
	}

	return validWords
}

func dfs(board [][]byte, i, j int, seen *[][]bool, word string, index int) bool {
	if index == len(word) {
		return true
	}

	(*seen)[j][i] = true
	char := rune(word[index])

	nn := neighbours(board, i, j, char, seen)
	for _, n := range nn {
		x, y := n[0], n[1]
		if dfs(board, x, y, seen, word, index+1) {
			return true
		}
	}

	return false
}

func neighbours(board [][]byte, i, j int, char rune, seen *[][]bool) [][]int {
	var nn [][]int
	if j > 0 && !(*seen)[j-1][i] && rune(board[j-1][i]) == char {
		nn = append(nn, []int{i, j - 1})
		(*seen)[j-1][i] = true
	}

	if i > 0 && !(*seen)[j][i-1] && rune(board[j][i-1]) == char {
		nn = append(nn, []int{i - 1, j})
		(*seen)[j][i-1] = true
	}

	if j < len(board)-1 && !(*seen)[j+1][i] && rune(board[j+1][i]) == char {
		nn = append(nn, []int{i, j + 1})
		(*seen)[j+1][i] = true
	}

	if i < len(board[0])-1 && !(*seen)[j][i+1] && rune(board[j][i+1]) == char {
		nn = append(nn, []int{i + 1, j})
		(*seen)[j][i+1] = true
	}

	return nn
}
