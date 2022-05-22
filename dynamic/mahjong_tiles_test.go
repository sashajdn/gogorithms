package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mahjongChecks = []func(tiles string) bool{
	MahjongTiles_BottomUp,
}

func TestMahjongTiles(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		inputTiles     string
		expectedOutput bool
	}{
		{
			name:           "example_one",
			inputTiles:     "11123",
			expectedOutput: true,
		},
		{
			name:           "example_two",
			inputTiles:     "11111",
			expectedOutput: true,
		},
		{
			name:           "example_three",
			inputTiles:     "1223",
			expectedOutput: false,
		},
		{
			name:           "example_four",
			inputTiles:     "11116933121",
			expectedOutput: false,
		},
		{
			name:           "example_five",
			inputTiles:     "11134577",
			expectedOutput: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range mahjongChecks {
				res := c(tt.inputTiles)

				assert.Equal(t, tt.expectedOutput, res)
			}
		})
	}
}
