package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSunsetViews(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		buildings      []int
		direction      string
		expectedOutput []int
	}{
		{
			name:           "example-one",
			buildings:      []int{3, 5, 4, 4, 3, 1, 3, 2},
			direction:      "WEST",
			expectedOutput: []int{0, 1},
		},
		{
			name:           "example-two",
			buildings:      []int{20, 2, 3, 1, 5, 6, 9, 1, 9, 9, 11, 10, 9, 12, 8},
			direction:      "EAST",
			expectedOutput: []int{0, 13, 14},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := SunsetViews(tt.buildings, tt.direction)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
