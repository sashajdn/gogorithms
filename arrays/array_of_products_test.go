package arrays

import (
	"testing"
)

func TestArrayOfProductsBasic(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{
			name:           "basic_pos_ints",
			input:          []int{1, 2, 3, 4},
			expectedOutput: []int{24, 12, 8, 6},
		},
		{
			name:           "pos_ints_with_0",
			input:          []int{1, 5, 0, 3},
			expectedOutput: []int{0, 0, 15, 0},
		},
		{
			name:           "pos_ints_with_two_zeros",
			input:          []int{100, 3, 7, 0, 5, 0, 9999},
			expectedOutput: []int{0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := arrayOfProductsBasic(tc.input)
			if !slicesEqual(result, tc.expectedOutput) {
				t.Fatalf("expected: %v, got: %v", tc.expectedOutput, result)
			}
		})
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
