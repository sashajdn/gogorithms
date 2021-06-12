package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDocument(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		chars          string
		doc            string
		expectedResult bool
	}{
		{
			name:           "empty",
			expectedResult: true,
		},
		{
			name:           "ae_example",
			chars:          "Bste!hetsi, ogEAxpelrt x ",
			doc:            "AlgoExpert is the Best!",
			expectedResult: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := GenerateDocument(tt.chars, tt.doc)
			assert.Equal(t, tt.expectedResult, res)
		})
	}
}
