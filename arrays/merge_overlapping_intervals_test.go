package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeOverlappingIntervals(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                    string
		inputIntervals          [][]int
		expectedOutputIntervals [][]int
	}{
		{
			name:                    "empty_intervals",
			inputIntervals:          [][]int{},
			expectedOutputIntervals: [][]int{},
		},
		{
			name: "intervals",
			inputIntervals: [][]int{
				{
					1, 2,
				},
				{
					3, 5,
				},
				{
					4, 7,
				},
				{
					6, 8,
				},
				{
					9, 10,
				},
			},
			expectedOutputIntervals: [][]int{
				{
					1, 2,
				},
				{
					3, 8,
				},
				{
					9, 10,
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := MergeOverlappingIntervals(tc.inputIntervals)
			assert.Equal(t, tc.expectedOutputIntervals, res)
		})
	}
}
