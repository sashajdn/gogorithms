package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfIslands(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          [][]rune
		expectedOutput int
	}{
		{
			name: "exammple_one",
			input: [][]rune{
				{1},
				{1},
			},
			expectedOutput: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := NumberOfIslands(tt.input)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
