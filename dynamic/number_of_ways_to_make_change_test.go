package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfWaysToMakeChange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		n              int
		denoms         []int
		expectedOutput int
	}{
		{
			name:           "zero_n",
			n:              0,
			expectedOutput: 1,
		},
		{
			name:           "empty_denoms",
			n:              10,
			denoms:         []int{},
			expectedOutput: 0,
		},
		{
			name:           "non_empty_denoms",
			n:              6,
			denoms:         []int{1, 5},
			expectedOutput: 2,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := NumberOfWaysToMakeChange(tt.n, tt.denoms)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
