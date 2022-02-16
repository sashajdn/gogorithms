package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountComponents(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		edges          [][]int
		n              int
		expectedOutput int
	}{
		{
			name: "example_one",
			edges: [][]int{
				{0, 1},
				{0, 2},
				{1, 2},
			},
			n:              4,
			expectedOutput: 2,
		},
		{
			name: "example_two",
			edges: [][]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{3, 4},
			},
			n:              5,
			expectedOutput: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := CountComponents_DFS(tt.n, tt.edges)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
