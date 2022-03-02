package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContinuousSubArraySum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		k              int
		expectedOutput bool
	}{
		{
			name:           "singular_zero_value_array",
			input:          []int{0},
			k:              1,
			expectedOutput: false,
		},
		{
			name:           "example_one",
			input:          []int{23, 2, 4, 6, 6},
			k:              7,
			expectedOutput: true,
		},
		{
			name:           "example_two",
			input:          []int{23, 2, 6, 4, 7},
			k:              13,
			expectedOutput: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := ContinuousSubArraySum(tt.input, tt.k)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
