package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxRussianDolls(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          [][]int
		expectedOutput int
	}{
		{
			name:           "empty",
			expectedOutput: 0,
		},
		{
			name: "all_the_same",
			input: [][]int{
				{1, 1},
				{1, 1},
				{1, 1},
				{1, 1},
				{1, 1},
				{1, 1},
			},
			expectedOutput: 1,
		},
		{
			name: "increasing_example_one",
			input: [][]int{
				{5, 4},
				{6, 4},
				{6, 7},
				{2, 3},
			},
			expectedOutput: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := MaxRussianDolls(tt.input)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
