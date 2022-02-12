package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageValueAtLevel(t *testing.T) {
	t.Parallel()

	exampleOneRoot := &Node{
		Value: 10,
		Left: &Node{
			Value: 5,
			Left: &Node{
				Value: 13,
			},
			Right: &Node{
				Value: 7,
				Left: &Node{
					Value: 6,
				},
			},
		},
		Right: &Node{
			Value: 5,
		},
	}

	tests := []struct {
		name           string
		root           *Node
		expectedOutput []int
	}{
		{
			name:           "example_one",
			root:           exampleOneRoot,
			expectedOutput: []int{10, 5, 10, 6},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := AverageValueAtLevel(tt.root)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
