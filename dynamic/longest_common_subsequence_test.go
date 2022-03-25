package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var longestCommonSubsequenceFs = []func(s1, s2 string) int{
	LongestCommonSubSequenceRecursive,
	LongestCommonSubSequenceRecursiveWithMemo,
	LongestCommonSubSequenceDynamicBottomUp,
	LongestCommonSubSequenceDynamicBottomUpOptimized,
}

func TestLongestCommonSubSequence(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		s1, s2         string
		expectedOutput int
	}{
		{
			name:           "empty_string",
			expectedOutput: 0,
		},
		{
			name:           "non_emtpy_strings_example_1",
			s1:             "abcde",
			s2:             "ace",
			expectedOutput: 3,
		},
		{
			name:           "non_emtpy_strings_example_2",
			s1:             "azbaaxxzziks",
			s2:             "akszzaais",
			expectedOutput: 6,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			for _, f := range longestCommonSubsequenceFs {
				res := f(tt.s1, tt.s2)
				assert.Equal(t, tt.expectedOutput, res)
			}
		})
	}
}
