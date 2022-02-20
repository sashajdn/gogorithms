package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinRemoveToMakeValidParentheses(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "example_one",
			input:          "abcd",
			expectedOutput: "abcd",
		},
		{
			name:           "example_two",
			input:          "))abcd",
			expectedOutput: "abcd",
		},
		{
			name:           "example_two",
			input:          "))abcd(",
			expectedOutput: "abcd",
		},
		{
			name:           "example_two",
			input:          "))a(b))cd(",
			expectedOutput: "a(b)cd",
		},
		{
			name:           "example_two",
			input:          "))a(((b)c)((((d)))",
			expectedOutput: "a(((b)c)((d)))",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := MinRemoveToMakeValidParentheses(tt.input)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
