package graphs

// SurroundedRegions ...
//
// T -> O(w * h)
// S -> O(w * h) due to recursive functions, worst case all elements in board are on recursive call stack.
func SurroundedRegions(board [][]rune) {
	// T -> O(w * h)
	// S -> O(w * h)
	for j := 0; j < len(board); j++ {
		for i := 0; i < len(board[0]); i++ {
			if board[j][i] != '0' {
				continue
			}

			if isConnectedToBorder(board, j, i) {
				continue
			}

			colour(board, j, i)
		}
	}

	// T -> O(w * h)
	// S -> O(1)
	for j := 0; j < len(board); j++ {
		for i := 0; i < len(board[0]); i++ {
			switch board[j][i] {
			case 'C':
				board[j][i] = 'O'
			case '#':
				board[j][i] = 'X'
			}
		}
	}
}

func isConnectedToBorder(board [][]rune, j, i int) bool {
	if board[j][i] != 'O' {
		return false
	}
	board[j][i] = 'C'

	var directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var isConnected bool
	if j == 0 || i == 0 || j == len(board)-1 || i == len(board[0]) {
		isConnected = true
	}
	for _, direction := range directions {
		dj, di := j+direction[0], i+direction[i]

		if dj < 0 || dj > len(board)-1 {
			continue
		}
		if di < 0 || di > len(board[0])-1 {
			continue
		}

		if isConnectedToBorder(board, dj, di) {
			isConnected = true
		}
	}

	return isConnected
}

func colour(board [][]rune, j, i int) {
	if board[j][i] != 'C' {
		return
	}
	board[j][i] = '#'

	var directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for _, direction := range directions {
		dj, di := j+direction[0], i+direction[i]

		if dj < 0 || dj > len(board)-1 {
			continue
		}
		if di < 0 || di > len(board[0])-1 {
			continue
		}

		colour(board, dj, di)
	}
}
