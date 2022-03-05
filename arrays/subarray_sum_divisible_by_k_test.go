package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubArrayDivisibleByK(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		k              int
		expectedOutput int
	}{
		{
			name:           "empty",
			input:          []int{},
			expectedOutput: 0,
		},
		{
			name:           "example_one",
			input:          []int{4, 5, 0, -2, -3, 1},
			k:              5,
			expectedOutput: 7,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := SubArrayDivisibleByK(tt.input, tt.k)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}

}
