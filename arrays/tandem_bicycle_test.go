package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTandemBicycle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		red, blue      []int
		fastest        bool
		expectedOutput int
	}{
		{
			name:           "fastest",
			red:            []int{1, 2, 3, 4},
			blue:           []int{1, 2, 3, 4},
			fastest:        true,
			expectedOutput: 14,
		},
		{
			name:           "not_fastest",
			red:            []int{1, 2, 3, 4},
			blue:           []int{1, 2, 3, 4},
			fastest:        false,
			expectedOutput: 10,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := TandemBicycle(tt.red, tt.blue, tt.fastest)
			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
