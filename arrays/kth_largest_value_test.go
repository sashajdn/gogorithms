package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargestValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		k              int
		expectedOutput int
	}{
		{
			name:           "example_one",
			input:          []int{6, 3, 7, 8},
			k:              3,
			expectedOutput: 6,
		},
		{
			name:           "example_two",
			input:          []int{3},
			k:              1,
			expectedOutput: 3,
		},
		{
			name:           "example_three",
			input:          []int{3, 6, 3, 7, 8, 3, 5, 6, 1},
			k:              2,
			expectedOutput: 7,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := findKthLargest(tt.input, tt.k)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
