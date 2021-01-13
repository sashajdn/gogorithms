package recursion

import (
	"testing"
)

func TestGetPermutations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		expectedOutput [][]int
	}{
		{
			name:           "empty-input",
			input:          []int{},
			expectedOutput: [][]int{},
		},
		{
			name:  "3-input",
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
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := GetPermutations(tc.input)
			if !slicesEqual(tc.expectedOutput, res) {
				t.Fatalf("expected -> %v, got -> %v", tc.expectedOutput, res)
			}
		})
	}
}

func slicesEqual(a, b [][]int) bool {
	if len(a) == 0 {
		return true
	}
	if len(a) == 1 {
		return sliceIn(a[0], b)
	}
	if sliceIn(a[0], b) && !sliceIn(a[0], a[1:]) {
		return slicesEqual(a[1:], b)
	}
	return false
}

func sliceIn(a []int, b [][]int) bool {
	if len(b) == 0 {
		return false
	}
	if sliceEqual(a, b[0]) {
		return true
	}
	return sliceIn(a, b[1:])
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	if a[0] != b[0] {
		return false
	}
	return sliceEqual(a[1:], b[1:])
}
