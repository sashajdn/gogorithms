package arrays

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKClosestPoints(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		points         [][]int
		k              int
		expectedOutput [][]int
	}{
		{
			name: "example_one",
			points: [][]int{
				{10, 10},
				{1, 3},
				{-2, 2},
				{9, 9},
				{7, 3},
			},
			k: 2,
			expectedOutput: [][]int{
				{1, 3},
				{-2, 2},
			},
		},
		{
			name: "example_two",
			points: [][]int{
				{1, 4},
				{1, 3},
				{-2, 2},
				{9, 9},
				{7, 9},
				{8, 3},
				{10, 10},
			},
			k: 4,
			expectedOutput: [][]int{
				{1, 3},
				{-2, 2},
				{1, 4},
				{8, 3},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := KClosestPoints(tt.points, tt.k)

			sorter(tt.expectedOutput)
			sorter(res)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}

func sorter(array [][]int) {
	sort.Slice(array, func(i, j int) bool {
		if array[i][0] == array[j][0] {
			return array[i][1] < array[j][1]
		}

		return array[i][0] < array[j][0]
	})
}
