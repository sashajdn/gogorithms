package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycleInGraph(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		edges          [][]int
		isCycleInGraph bool
	}{
		{
			name: "example_1",
			edges: [][]int{
				{1},
				{2, 3, 4, 5, 6, 7},
				{},
				{2, 7},
				{5},
				{},
				{4},
				{},
			},
			isCycleInGraph: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := CycleInGraph(tt.edges)

			assert.Equal(t, tt.isCycleInGraph, res)
		})
	}
}
