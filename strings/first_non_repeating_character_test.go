package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstNonRepeatingCharacter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name:           "null_string",
			input:          "",
			expectedOutput: -1,
		},
		{
			name:           "single_char_string",
			input:          "a",
			expectedOutput: 0,
		},
		{
			name:           "same_char_only",
			input:          "aa",
			expectedOutput: -1,
		},
		{
			name:           "non_repeating",
			input:          "ahaidxidh",
			expectedOutput: 5,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := FirstNonRepeatingCharacter(tt.input)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
