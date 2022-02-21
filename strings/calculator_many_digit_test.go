package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateManyDigit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name:           "example_one",
			input:          "11 + 7 * 5",
			expectedOutput: 46,
		},
		{
			name:           "example_two",
			input:          "100 - 10 / 2    + 10 * 4 - 10",
			expectedOutput: 125,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := CalculateManyDigit_WithStack(tt.input)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
