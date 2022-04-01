package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestDivisor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		threshold      int
		expectedOutput int
	}{
		{
			name:           "empty_numbers",
			input:          []int{},
			threshold:      10,
			expectedOutput: 0,
		},
		{
			name:           "example_one",
			input:          []int{1, 2, 5, 9},
			threshold:      6,
			expectedOutput: 5,
		},
		{
			name:           "example_two",
			input:          []int{44, 22, 33, 11, 1},
			threshold:      5,
			expectedOutput: 44,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := FindSmallestDivisor(tt.input, tt.threshold)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
