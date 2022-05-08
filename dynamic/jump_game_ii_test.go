package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var checks = []func(nums []int) int{
	JumpGameII_Dynamic,
	JumpGameII_Greedy,
}

func TestJumpGameII(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		expectedOutput int
	}{
		{
			name:           "example_one",
			input:          []int{2, 3, 1, 1, 4},
			expectedOutput: 2,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range checks {
				res := c(tt.input)

				assert.Equal(t, tt.expectedOutput, res)
			}
		})
	}
}
