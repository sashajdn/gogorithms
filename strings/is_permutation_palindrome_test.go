package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPermutationPalindrome(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedOutput bool
	}{
		{
			name:           "example_one",
			input:          "aab",
			expectedOutput: true,
		},
		{
			name:           "example_two",
			input:          "aabc",
			expectedOutput: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := isPermutationPalindrome(tt.input)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
