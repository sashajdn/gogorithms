package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicCalculator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedResult int
	}{
		{
			name:           "example_one",
			input:          "3 + 2 * 2",
			expectedResult: 7,
		},
		{
			name:           "example_two",
			input:          "3 + 10 * 6 * 4 - 3 + 1 * 10",
			expectedResult: 250,
		},
		{
			name:           "example_three",
			input:          "0-2147483647",
			expectedResult: -2147483647,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := BasicCalculator(tt.input)

			assert.Equal(t, tt.expectedResult, res)
		})
	}
}
