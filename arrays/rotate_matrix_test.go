package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateMatrix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          [][]int
		expectedOutput [][]int
	}{
		{
			name: "example_len_three",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expectedOutput: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "example_len_four",
			input: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
			expectedOutput: [][]int{
				{10, 11, 12, 1},
				{9, 16, 13, 2},
				{8, 15, 14, 3},
				{7, 6, 5, 4},
			},
		},
		{
			name: "example_len_four",
			input: [][]int{
				{1, 2, 3, 4, 5},
				{16, 17, 18, 19, 6},
				{15, 24, 25, 20, 7},
				{14, 23, 22, 21, 8},
				{13, 12, 11, 10, 9},
			},
			expectedOutput: [][]int{
				{13, 14, 15, 16, 1},
				{12, 23, 24, 17, 2},
				{11, 22, 25, 18, 3},
				{10, 21, 20, 19, 4},
				{9, 8, 7, 6, 5},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			RotateMatrix(tt.input)

			assert.Equal(t, tt.expectedOutput, tt.input)
		})
	}
}
