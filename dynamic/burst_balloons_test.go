package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var checksBurstBalloons = []func(nums []int) int{
	BurstBalloonsBruteForce,
	BurstBalloonsDynamicBottomUp,
}

func TestBurstBalloons(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		expectedOutput int
	}{
		{
			name:           "example_one",
			input:          []int{3, 1, 5, 8},
			expectedOutput: 167,
		},
		{
			name:           "example_two",
			input:          []int{1, 5},
			expectedOutput: 10,
		},
		{
			name:           "example_three",
			input:          []int{3, 1, 10, 11, 5, 8},
			expectedOutput: 1622,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range checksBurstBalloons {
				res := c(tt.input)

				assert.Equal(t, tt.expectedOutput, res)
			}
		})
	}
}
