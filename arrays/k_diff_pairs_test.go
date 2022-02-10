package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKDiffPairs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		nums           []int
		k              int
		expectedOutput int
	}{
		{
			name:           "example_one",
			nums:           []int{3, 1, 4, 1, 5},
			k:              2,
			expectedOutput: 2,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := findPairs(tt.nums, tt.k)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}

}
