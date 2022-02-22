package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRedundantConnection(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		edges          [][]int
		expectedOutput []int
	}{
		{
			name: "example_one",
			edges: [][]int{
				{1, 2},
				{1, 3},
				{2, 3},
			},
			expectedOutput: []int{2, 3},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := FindRedundantConnection(tt.edges)
			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
