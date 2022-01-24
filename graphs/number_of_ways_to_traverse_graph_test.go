package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testFuncs = []func(width int, height int) int{
	NumberOfWaysToTraverseGraph_Dynamic,
	NumberOfWaysToTraverseGraph_Recursive,
	NumberOfWaysToTraverseGraph_Factorial,
}

func TestNumberOfWaysToTraverseAGraph(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		width, height  int
		expectedOutput int
	}{
		{
			name:           "2x1",
			width:          2,
			height:         1,
			expectedOutput: 1,
		},
		{
			name:           "1x2",
			width:          1,
			height:         2,
			expectedOutput: 1,
		},
		{
			name:           "3x4",
			width:          3,
			height:         4,
			expectedOutput: 10,
		},
		{
			name:           "2x3",
			width:          2,
			height:         3,
			expectedOutput: 3,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, tf := range testFuncs {
				res := tf(tt.width, tt.height)

				assert.Equal(t, tt.expectedOutput, res)
			}
		})
	}

}

func TestFactorial(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          int
		expectedOutput int
	}{
		{
			name:           "0",
			input:          0,
			expectedOutput: 1,
		},
		{
			name:           "1",
			input:          1,
			expectedOutput: 1,
		},
		{
			name:           "4",
			input:          4,
			expectedOutput: 24,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := factorial(tt.input)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
