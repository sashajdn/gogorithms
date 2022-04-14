package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetMatrixToZero(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                 string
		inputMatrix          [][]int
		expectedOutputMatrix [][]int
	}{
		{
			name: "example_one",
			inputMatrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 6},
				{7, 8, 9, 10},
				{11, 12, 13, 14},
			},
			expectedOutputMatrix: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 8, 9, 0},
				{0, 12, 13, 0},
			},
		},
		{
			name: "example_two",
			inputMatrix: [][]int{
				{5, 1, 2, 0},
				{3, 4, 5, 6},
				{7, 0, 9, 10},
				{11, 12, 13, 14},
			},
			expectedOutputMatrix: [][]int{
				{0, 0, 0, 0},
				{3, 0, 5, 0},
				{0, 0, 0, 0},
				{11, 0, 13, 0},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := SetMatrixToZero(tt.inputMatrix)

			assert.Equal(t, tt.expectedOutputMatrix, res)
		})
	}
}
