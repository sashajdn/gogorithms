package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var slidingWindowMaximumChecks = []func(nums []int, k int) []int{
	// SlidingWindowMaximum_BruteForce,
	// SlidingWindowMaximum_Deque,
	SlidingWindowMaximum_Dynamic,
}

func TestSlidingWindowMaximum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		k              int
		expectedOutput []int
	}{
		{
			name:           "empty",
			input:          []int{},
			k:              10,
			expectedOutput: []int{},
		},
		{
			name:           "non-empty",
			input:          []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:              3,
			expectedOutput: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:           "non_empty_2",
			input:          []int{10, 3, 4, 9, 1, 3, 8, 10, 11, 15, 12, 1, 100},
			k:              3,
			expectedOutput: []int{10, 9, 9, 9, 8, 10, 11, 15, 15, 15, 100},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range slidingWindowMaximumChecks {
				res := c(tt.input, tt.k)
				assert.Equal(t, tt.expectedOutput, res)
			}
		})
	}
}
