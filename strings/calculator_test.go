package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name:           "example_one",
			input:          "3+2*2",
			expectedOutput: 7,
		},
		{
			name:           "example_two",
			input:          "3+ 2* 2 / 2",
			expectedOutput: 5,
		},
		{
			name:           "example_three",
			input:          "3 + 2",
			expectedOutput: 5,
		},
		{
			name:           "example_four",
			input:          "10 * 10 / 2 + 3 - 10 / 5",
			expectedOutput: 51,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := Calculate(tt.input)

			assert.Equal(t, tt.expectedOutput, res)

		})
	}
}
