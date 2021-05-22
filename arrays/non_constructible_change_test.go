package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonConstructibleChange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		inputCoins     []int
		expectedOutput int
	}{
		{
			name:           "empty-coins",
			inputCoins:     []int{},
			expectedOutput: 1,
		},
		{
			name:           "example",
			inputCoins:     []int{1, 2, 5},
			expectedOutput: 4,
		},
		{
			name:           "example_2",
			inputCoins:     []int{5, 7, 1, 1, 2, 3, 22},
			expectedOutput: 20,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := NonConstructibleChange(tc.inputCoins)
			assert.Equal(t, tc.expectedOutput, res)
		})
	}
}
