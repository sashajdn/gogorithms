package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPermutation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{
			name:           "len_two_array",
			input:          []int{1, 2},
			expectedOutput: []int{2, 1},
		},
		{
			name:           "len_four_array_one",
			input:          []int{1, 3, 4, 2},
			expectedOutput: []int{1, 4, 2, 3},
		},
		{
			name:           "len_four_array_two",
			input:          []int{2, 4, 3, 1},
			expectedOutput: []int{3, 1, 2, 4},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			NextPermutation(tt.input)

			assert.Equal(t, tt.expectedOutput, tt.input)
		})
	}
}
