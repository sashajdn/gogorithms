package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskAssignment(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		k              int
		tasks          []int
		expectedOutput [][]int
	}{
		{
			name:  "example_one",
			k:     3,
			tasks: []int{1, 3, 5, 3, 1, 4},
			expectedOutput: [][]int{
				{0, 2},
				{4, 5},
				{1, 3},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := TaskAssignment(tt.k, tt.tasks)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}

}
