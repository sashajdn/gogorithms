package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var interweavingStringsChecks = []func(s1, s2, s3 string) bool{
	InterweavingStrings_Recursive,
	InterweavingStrings_TopDown,
	InterweavingStrings_BottomUp2D,
	InterweavingStrings_BottomUp1D,
}

func TestInterweavingStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		s1, s2, s3     string
		expectedOutput bool
	}{
		{
			name:           "empty",
			s1:             "",
			s2:             "",
			s3:             "",
			expectedOutput: true,
		},
		{
			name:           "invalid_one",
			s1:             "this is a long string",
			s2:             "short string",
			s3:             "blah diff size",
			expectedOutput: false,
		},
		{
			name:           "valid_one",
			s1:             "aabcc",
			s2:             "dbbca",
			s3:             "aadbbcbcac",
			expectedOutput: true,
		},
		{
			name:           "invalid_two",
			s1:             "aabcc",
			s2:             "dbbca",
			s3:             "aadbbbaccc",
			expectedOutput: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range interweavingStringsChecks {
				res := c(tt.s1, tt.s2, tt.s3)

				assert.Equal(t, tt.expectedOutput, res)
			}
		})
	}
}
