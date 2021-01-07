package arrays

import "testing"

func TestFirstDuplicateValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		expectedOutput int
	}{
		{
			name:           "empty-array",
			input:          []int{},
			expectedOutput: -1,
		},
		{
			name:           "len-5",
			input:          []int{1, 2, 3, 4, 1},
			expectedOutput: 1,
		},
		{
			name:           "len-10",
			input:          []int{5, 3, 2, 7, 8, 4, 1, 8, 9, 1},
			expectedOutput: 8,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := FirstDuplicateValue(tc.input)
			if result != tc.expectedOutput {
				t.Fatalf("expected -> %v, got -> %v", tc.expectedOutput, result)
			}
		})
	}
}
