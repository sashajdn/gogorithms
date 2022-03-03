package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKRadiusSubarrayAverage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		k              int
		expectedOutput []int
	}{
		{
			name:           "example_one",
			input:          []int{7, 4, 3, 9, 1, 8, 5, 2, 6},
			k:              3,
			expectedOutput: []int{-1, -1, -1, 5, 4, 4, -1, -1, -1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := KRadiusSubarrayAverage(tt.input, tt.k)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
