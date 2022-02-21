package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexEqualsValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		expectedOutput int
	}{
		{
			name:           "empty_array",
			input:          []int{},
			expectedOutput: -1,
		},
		{
			name:           "singular_array",
			input:          []int{0},
			expectedOutput: 0,
		},
		{
			name:           "singular_incorrect_array",
			input:          []int{10},
			expectedOutput: -1,
		},
		{
			name:           "correct_array",
			input:          []int{8, 2, 6, 3},
			expectedOutput: 3,
		},
		{
			name:           "multiple_correct_array",
			input:          []int{0, 1, 6, 3},
			expectedOutput: 0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := IndexEqualsValue(tt.input)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
