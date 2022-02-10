package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	checks = []func(array []int, order []int) []int{
		ThreeNumberSort,
		ThreeNumberSort_Linear,
	}
)

func TestThreeNumberSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		array, order   []int
		expectedOutput []int
	}{
		{
			name:           "example_one",
			array:          []int{1, 0, 0, -1, -1, 0, 1, 1},
			order:          []int{0, 1, -1},
			expectedOutput: []int{0, 0, 0, 1, 1, 1, -1, -1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range checks {
				res := c(tt.array, tt.order)

				assert.Equal(t, tt.expectedOutput, res)
			}

		})
	}
}
