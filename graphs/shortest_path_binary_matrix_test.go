package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPathBinaryMatrix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          [][]int
		expectedOutput int
	}{
		{
			name: "example_one_successful",
			input: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			expectedOutput: 4,
		},
		{
			name: "example_two_failure",
			input: [][]int{
				{1, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			expectedOutput: -1,
		},
		{
			name: "example_four_failure",
			input: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 1, 1},
			},
			expectedOutput: -1,
		},
		{
			name: "example_three_multiple_paths",
			input: [][]int{
				{0, 0, 0, 0, 0},
				{1, 0, 1, 1, 0},
				{1, 0, 0, 1, 0},
				{1, 0, 1, 0, 0},
				{1, 1, 0, 0, 0},
			},
			expectedOutput: 5,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := ShortestPathBinaryMatrix(tt.input)

			assert.Equal(t, tt.expectedOutput, res)

		})
	}
}
