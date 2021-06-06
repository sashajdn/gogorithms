package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunLengthEncoding(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		input         string
		expectedOuput string
	}{
		{
			name:          "algo-expert-example",
			input:         "AAAAAAAAAAAAABBCCCCDD",
			expectedOuput: "9A4A2B4C2D",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := RunLengthEncoding(tt.input)
			assert.Equal(t, tt.expectedOuput, res)
		})
	}
}
