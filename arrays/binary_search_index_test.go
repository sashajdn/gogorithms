package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var binarySearchIndexFuncs = []func([]int, int) int{
	BinarySearchIndexIterative,
	BinarySearchIndexRecursive,
}

func TestBinarySearchIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		target         int
		expectedOutput int
	}{
		{
			name:           "example_one",
			input:          []int{1, 4, 6, 8},
			target:         5,
			expectedOutput: 2,
		},
		{
			name:           "example_two",
			input:          []int{1, 4, 6, 8},
			target:         7,
			expectedOutput: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range binarySearchIndexFuncs {
				res := c(tt.input, tt.target)
				assert.Equal(t, tt.expectedOutput, res)
			}

		})
	}
}
