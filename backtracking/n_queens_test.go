package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveNQueens_Sets(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		queens         int
		expectedOutput int
	}{
		{
			name:           "example_four_queens",
			queens:         4,
			expectedOutput: 2,
		},
		{
			name:           "example_six_queens",
			queens:         6,
			expectedOutput: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := SolveNQueensSets(tt.queens)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}

}

func TestSolveNQueens_Recursive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		queens         int
		expectedOutput [][][]int
	}{
		{
			name:   "example_four_queens",
			queens: 4,
			expectedOutput: [][][]int{
				{
					{0, 1, 0, 0},
					{0, 0, 0, 1},
					{1, 0, 0, 0},
					{0, 0, 1, 0},
				},
				{
					{0, 0, 1, 0},
					{1, 0, 0, 0},
					{0, 0, 0, 1},
					{0, 1, 0, 0},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := SolveNQueensRecursive(tt.queens)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
