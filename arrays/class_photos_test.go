package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClassPhotos(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		red            []int
		blue           []int
		expectedResult bool
	}{
		{
			name:           "single_values",
			red:            []int{1},
			blue:           []int{100},
			expectedResult: true,
		},
		{
			name:           "multi-value-true",
			red:            []int{6, 4, 2},
			blue:           []int{1, 3, 5},
			expectedResult: true,
		},
		{
			name:           "multi-value-false",
			red:            []int{6, 4, 2},
			blue:           []int{1, 7, 3},
			expectedResult: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := ClassPhotos(tt.red, tt.blue)
			assert.Equal(t, tt.expectedResult, res)
		})
	}
}
