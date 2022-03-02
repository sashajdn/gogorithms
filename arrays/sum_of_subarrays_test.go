package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sumOfSubArraysFuncs = []func([]int, int) int{
	SumOfSubArraysBruteForce,
	SumOfSubArraysCumSum,
	SumOfSubArraysWithoutSpace,
	SumOfSubArraysHashMap,
}

func TestSumOfSubArrays(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		target         int
		expectedOutput int
	}{
		{
			name:           "empty_array",
			input:          []int{},
			expectedOutput: 0,
		},
		{
			name:           "example_one",
			input:          []int{1, 1, 1},
			target:         2,
			expectedOutput: 2,
		},
		{
			name:           "example_two",
			input:          []int{4, 5, 3, 2, 1, 9},
			target:         9,
			expectedOutput: 2,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for _, f := range sumOfSubArraysFuncs {
				res := f(tt.input, tt.target)
				assert.Equal(t, tt.expectedOutput, res)
			}
		})
	}
}
