package sorting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var quickSortCheckers = []func([]int){
	QuickSort_FirstValue,
}

func TestQuickSort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{
			name:           "example_one",
			input:          []int{5, 7, 3, 8, 2, 3},
			expectedOutput: []int{2, 3, 3, 5, 7, 8},
		},
		{
			name:           "example_two",
			input:          []int{2, 1},
			expectedOutput: []int{1, 2},
		},
		{
			name:           "example_three",
			input:          []int{4, 1, 5, 0, -9, -3, -3, 9, 3, -4, -9, 8, 1, -3, -7, -4, -9, -1, -7, -2, -7, 4},
			expectedOutput: []int{-9, -9, -9, -7, -7, -7, -4, -4, -3, -3, -3, -2, -1, 0, 1, 1, 3, 4, 4, 5, 8, 9},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range quickSortCheckers {
				c(tt.input)
				assert.Equal(t, tt.expectedOutput, tt.input)

				fmt.Println("Res: ", tt.input, tt.expectedOutput)
			}
		})
	}
}
