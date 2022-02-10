package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerset(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		expectedOutput [][]int
	}{
		{
			name:  "example_one",
			input: []int{1, 2, 3},
			expectedOutput: [][]int{
				{},
				{1},
				{2},
				{3},
				{1, 2},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := Powerset(tt.input)

			testSorter(&res)
			testSorter(&tt.expectedOutput)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
