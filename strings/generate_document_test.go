package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDocument(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		characters     string
		document       string
		expectedOutput bool
	}{
		{
			name:           "empty_document_no_chars",
			characters:     "",
			document:       "",
			expectedOutput: true,
		},
		{
			name:           "empty_document_with_chars",
			characters:     "aaa",
			document:       "",
			expectedOutput: true,
		},
		{
			name:           "fewer_chars_than_doc_length",
			characters:     "ddd",
			document:       "dddd",
			expectedOutput: false,
		},
		{
			name:           "correct_chars",
			characters:     "aabbcc",
			document:       "baccba",
			expectedOutput: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := GenerateDocument(tt.characters, tt.document)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
