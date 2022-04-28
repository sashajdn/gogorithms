package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var longestDuplicateSubstringChecks = []func(s string) int{
	LongestDuplicateSubstring_HashMap,
	LongestDuplicateSubstring_TrieBinarySearch,
}

func TestLongestDuplicateSubstring(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name:           "empty",
			input:          "",
			expectedOutput: 0,
		},
		{
			name:           "no_duplicates",
			input:          "abcd",
			expectedOutput: 0,
		},
		{
			name:           "duplicates",
			input:          "banana",
			expectedOutput: 3,
		},
		{
			name:           "duplicates_1",
			input:          "abcdabababcd",
			expectedOutput: 4,
		},
		{
			name:           "duplicates_2",
			input:          "abjudkajxdkeabkjpl",
			expectedOutput: 2,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range longestDuplicateSubstringChecks {
				res := c(tt.input)

				assert.Equal(t, tt.expectedOutput, res)
			}
		})
	}
}
