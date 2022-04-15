package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildSpiralMatrixs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          int
		expectedOutput [][]int
	}{
		{
			name:  "3",
			input: 3,
			expectedOutput: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name:  "4",
			input: 4,
			expectedOutput: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := BuildSpiralMatrix(tt.input)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
