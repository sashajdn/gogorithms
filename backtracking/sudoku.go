package backtracking

// SolveSudoku ...
//
// T -> O(1) since we have a well defined board, if the board size was unbounded, this would become O(n ** 2)
// S -> O(1) since we don't have to create no auxilary space & any recursive stack is well defined - as long as the board is.
func SolveSudoku(board [][]int) {
	for j := 0; j < len(board); j++ {
		for i := 0; i < len(board[0]); i++ {
			if board[j][i] != 0 {
				continue
			}

			for n := 1; n < 10; n++ {
				if !isPossible(board, j, i, n) {
					continue
				}

				board[j][i] = n
				SolveSudoku(board)
				board[j][i] = 0
			}
			return
		}
	}
}

func isPossible(board [][]int, x, y, n int) bool {
	for j := 0; j < len(board); j++ {
		if board[j][x] == n {
			return false
		}
	}

	for i := 0; i < len(board[0]); i++ {
		if board[y][i] == n {
			return false
		}
	}

	m, n := (x/3)*3, (y/3)*3
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			col, row := n+j, m+i

			if board[col][row] == n {
				return false
			}
		}
	}

	//
	return true
}
