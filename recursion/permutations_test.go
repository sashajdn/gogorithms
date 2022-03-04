package recursion

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var permutationsChecker = []func(array []int) [][]int{
	GetPermutations_Better,
	GetPermutations,
	GetPermutations_Other,
}

func TestGetPermutations(t *testing.T) {
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
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			sort.Slice(tt.expectedOutput, func(i, j int) bool {
				for k := 0; k < len(tt.expectedOutput[0]); k++ {
					if tt.expectedOutput[i][k] == tt.expectedOutput[j][k] {
						continue
					}

					return tt.expectedOutput[i][k] < tt.expectedOutput[j][k]
				}

				return false
			})

			for _, checker := range permutationsChecker {

				res := checker(tt.input)
				testSorter(&res)
				testSorter(&tt.expectedOutput)

				assert.Equal(t, tt.expectedOutput, res)
			}
		})
	}
}

func testSorter(s *[][]int) {
	sort.Slice(*s, func(i, j int) bool {
		for k := 0; k < len((*s)[0]); k++ {
			if (*s)[i][k] == (*s)[j][k] {
				continue
			}

			return (*s)[i][k] < (*s)[j][k]
		}

		return false
	})
}
