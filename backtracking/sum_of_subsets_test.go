package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfSubsets(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		target         int
		expectedOutput [][]int
	}{
		{
			name:   "example_one_all_positive",
			input:  []int{2, 1, 3, 5},
			target: 8,
			expectedOutput: [][]int{
				{3, 5},
				{2, 1, 5},
			},
		},
		{
			name:   "example_one_positive_and_negative",
			input:  []int{2, 1, 3, 5, 10, -10},
			target: 8,
			expectedOutput: [][]int{
				{3, 5},
				{3, 5, 10, -10},
				{2, 1, 5},
				{2, 1, 5, 10, -10},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := SumOfSubsets(tt.input, tt.target)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
