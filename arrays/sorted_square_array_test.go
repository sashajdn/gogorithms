package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquareArray(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{
			name:           "basic_array",
			input:          []int{1, 2, 3, 4, 5},
			expectedOutput: []int{1, 4, 9, 16, 25},
		},
		{
			name:           "empty_array",
			input:          []int{},
			expectedOutput: []int{},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := SortedSquareArray(tc.input)
			assert.Equal(t, tc.expectedOutput, res)
		})
	}
}
