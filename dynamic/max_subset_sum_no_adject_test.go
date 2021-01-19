package dynamic

import "testing"

func TestMaxSubsetNoAdjacent(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		input          []int
		expectedOutput int
	}{
		{
			name:           "empty-array",
			input:          []int{},
			expectedOutput: 0,
		},
		{
			name:           "array-size-one",
			input:          []int{10},
			expectedOutput: 10,
		},
		{
			name:           "large-array",
			input:          []int{75, 105, 120, 75, 90, 135},
			expectedOutput: 330,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := MaxSubsetNoAdjacent(tc.input)
			assert(t, tc.expectedOutput, result)
		})
	}
}

func assertInt(t *testing.T, expected, result int) {
	if expected != result {
		t.Fatalf("expected -> %v, got -> %v", expected, result)
	}
}
