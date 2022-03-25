package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingTriplets(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		expectedOutput bool
	}{
		{
			name:           "too_small",
			input:          []int{1},
			expectedOutput: false,
		},
		{
			name:           "not_valid",
			input:          []int{1, 0, 0, 0, -1, 4},
			expectedOutput: false,
		},
		{
			name:           "alternating",
			input:          []int{1, 2, 1, 2, 1, 2, 1, 2},
			expectedOutput: false,
		},
		{
			name:           "valid",
			input:          []int{99, 88, 77, 66, 75, 1, 0, 100, 2},
			expectedOutput: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := increasingTriplets(tt.input)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
