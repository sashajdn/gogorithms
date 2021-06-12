package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryConverter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{
			name:           "a",
			input:          "10011011",
			expectedOutput: 155,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := BinaryConverter(tt.input)
			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
