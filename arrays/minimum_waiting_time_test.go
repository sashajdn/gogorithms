package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumWaitingTime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		inputQueries   []int
		expectedOutput int
	}{
		{
			name:           "single-value",
			inputQueries:   []int{100},
			expectedOutput: 0,
		},
		{
			name:           "multi-value",
			inputQueries:   []int{3, 2, 1, 2, 6},
			expectedOutput: 17,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := MinimumWaitingTime(tc.inputQueries)
			assert.Equal(t, tc.expectedOutput, res)
		})
	}
}
