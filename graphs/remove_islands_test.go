package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveIslands(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		inputMatrix    [][]int
		expectedOutput [][]int
	}{
		{
			name: "example-one",
			inputMatrix: [][]int{
				{1, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 1},
				{0, 0, 1, 0, 1, 0},
				{1, 1, 0, 0, 1, 0},
				{1, 0, 1, 1, 0, 0},
				{1, 0, 0, 0, 0, 1},
			},
			expectedOutput: [][]int{
				{1, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 1},
				{0, 0, 0, 0, 1, 0},
				{1, 1, 0, 0, 1, 0},
				{1, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 1},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := RemoveIslands(tt.inputMatrix)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
