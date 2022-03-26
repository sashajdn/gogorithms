package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSumIncreasingSubsequence(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                          string
		inputArray                    []int
		expectedMaxSum                int
		expectedMaxIncreasingSequence []int
	}{
		{},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			maxSum, maxIncreasingSequence := MaxSumIncreasingSubsequence(tt.inputArray)

			assert.Equal(t, tt.expectedMaxSum, maxSum)
			assert.Equal(t, tt.expectedMaxIncreasingSequence, maxIncreasingSequence)
		})
	}
}
