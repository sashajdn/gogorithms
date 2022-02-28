package backtracking

type Set map[int]struct{}

func (s *Set) Add(v int) bool {
	if _, ok := (*s)[v]; ok {
		return false
	}

	(*s)[v] = struct{}{}
	return true
}

func (s *Set) Delete(v int) bool {
	if _, ok := (*s)[v]; ok {
		delete(*s, v)
		return true
	}

	return false
}

func (s *Set) In(v int) bool {
	if _, ok := (*s)[v]; ok {
		return true
	}

	return false
}

// SolveNQueensSets ...
//
// T -> Upper Bound: O(n!), where `n` is the number of queens.
// S -> O(n), where n is the number of queens; this is the total number of recursive calls - since we recurse down `n` rows maximum.
func SolveNQueensSets(queens int) int {
	var (
		cols     Set = map[int]struct{}{}
		posDiags Set = map[int]struct{}{}
		negDiags Set = map[int]struct{}{}
	)

	var count int
	solveNQueensSets(&cols, &posDiags, &negDiags, queens, 0, &count)
	return count
}

func solveNQueensSets(cols, posDiags, negDiags *Set, n, row int, count *int) {
	if row == n {
		*count++
		return
	}

	for i := 0; i < n; i++ {
		posDiag, negDiag := i+row, i-row
		if cols.In(i) || posDiags.In(posDiag) || negDiags.In(negDiag) {
			continue
		}

		cols.Add(i)
		posDiags.Add(posDiag)
		negDiags.Add(negDiag)

		solveNQueensSets(cols, posDiags, negDiags, n, row+1, count)

		cols.Delete(i)
		posDiags.Delete(posDiag)
		negDiags.Delete(negDiag)
	}
}

// SolveNQueensRecursive ...
//
// T -> O(n**4)
// S -> O(1)
func SolveNQueensRecursive(queens int) [][][]int {
	var board = make([][]int, 0, queens)
	for i := 0; i < queens; i++ {
		board = append(board, make([]int, queens))
	}

	var outputs = [][][]int{}
	solveNQueens(board, queens, &outputs, 0)
	return outputs
}

func solveNQueens(board [][]int, queens int, outputs *[][][]int, row int) {
	if queens < 1 {
		boardCopy := copyBoard(board)
		*outputs = append(*outputs, boardCopy)
		return
	}

	for j := row; j < len(board); j++ {
		for i := 0; i < len(board); i++ {
			if board[j][i] == 1 {
				continue
			}

			if !isQueenPossible(board, i, j) {
				continue
			}

			board[j][i] = 1
			solveNQueens(board, queens-1, outputs, row+1)
			board[j][i] = 0
		}
	}
}

func isQueenPossible(board [][]int, x, y int) bool {
	for i := 0; i < len(board); i++ {
		if board[y][i] == 1 {
			return false
		}
	}

	for j := 0; j < len(board); j++ {
		if board[j][x] == 1 {
			return false
		}
	}

	for k := 1; k < len(board)-1; k++ {
		if x+k < len(board) && y+k < len(board) {
			if board[y+k][x+k] == 1 {
				return false
			}
		}

		if x-k >= 0 && y-k >= 0 {
			if board[y-k][x-k] == 1 {
				return false
			}
		}

		if x+k < len(board) && y-k >= 0 {
			if board[y-k][x+k] == 1 {
				return false
			}
		}

		if x-k >= 0 && y+k < len(board) {
			if board[y+k][x-k] == 1 {
				return false
			}
		}
	}

	return true
}

func copyBoard(board [][]int) [][]int {
	var boardCopy = make([][]int, 0, len(board))
	for j := 0; j < len(board); j++ {
		row := make([]int, len(board))
		for i := 0; i < len(board); i++ {
			row[i] = board[j][i]
		}
		boardCopy = append(boardCopy, row)
	}
	return boardCopy
}
