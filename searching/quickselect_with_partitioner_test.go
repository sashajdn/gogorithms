package searching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortWithPartitioner(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		k              int
		partitioner    QuickSelectType
		expectedResult int
	}{
		{
			name:           "example_one_first_elem",
			input:          []int{5, 3, 7, 8, 9},
			k:              2, // 2nd min element.
			partitioner:    QuickSelectWithPartitionerFirstElement,
			expectedResult: 5,
		},
		{
			name:           "example_two_first_elem",
			input:          []int{5, 3, 7, 8, 9, 88},
			k:              4, // 4th min element.
			partitioner:    QuickSelectWithPartitionerFirstElement,
			expectedResult: 8,
		},
		{
			name:           "example_one_median",
			input:          []int{5, 3, 7, 8, 9, 88},
			k:              4, // 4th min element.
			partitioner:    QuickSelectWithPartitionerMedian,
			expectedResult: 8,
		},
		{
			name:           "example_two_median",
			input:          []int{5, 3},
			k:              2, // 2nd min element.
			partitioner:    QuickSelectWithPartitionerMedian,
			expectedResult: 5,
		},
		{
			name:           "example_one_random",
			input:          []int{10, 14, 6, 2, 4, 8, 5, 1, 3},
			k:              5, // 2nd min element.
			partitioner:    QuickSelectWithPartitionerRandom,
			expectedResult: 5,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := QuickSelectWithPartitioner(tt.input, tt.k, tt.partitioner)

			assert.Equal(t, tt.expectedResult, res)
		})
	}
}
