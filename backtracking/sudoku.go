package backtracking

// SolveSudoku ...
//
// T -> O(1)
// S -> O(1)
func SolveSudoku(board [][]int) {
	for col := 0; col < len(board); col++ {
		for row := 0; row < len(board); row++ {
			if board[col][row] != 0 {
				continue
			}

			for i := 1; i < 10; i++ {
				if !isPossible(board, col, row, i) {
					continue
				}

				board[col][row] = i
				SolveSudoku(board)
				board[col][row] = 0
			}

			return
		}
	}
}

func isPossible(grid [][]int, y, x, n int) bool {
	for i := 0; i < 9; i++ {
		if grid[y][i] == n || grid[x][i] == n {
			return false
		}
	}

	x0 := (x / 3) * 3
	y0 := (y / 3) * 3

	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			if grid[y0+j][x0+i] == n {
				return false
			}
		}
	}

	return true
}
