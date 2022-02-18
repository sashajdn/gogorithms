package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		target         int
		expectedOutput bool
	}{
		{
			name:           "example_one",
			input:          []int{1, 4, 6, 7, 8, 9, 23},
			target:         7,
			expectedOutput: true,
		},
		{
			name:           "example_two",
			input:          []int{1, 4, 6, 7, 8, 9, 23},
			target:         23,
			expectedOutput: true,
		},
		{
			name:           "example_three",
			input:          []int{1, 4, 6, 7, 8, 9, 23},
			target:         24,
			expectedOutput: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := binarySearch(tt.input, tt.target)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
