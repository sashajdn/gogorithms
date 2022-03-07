package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPathBinaryMatrix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          [][]int
		expectedOutput int
	}{
		{},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := ShortestPathBinaryMatrix(tt.input)

			assert.Equal(t, tt.expectedOutput, res)

		})
	}
}
