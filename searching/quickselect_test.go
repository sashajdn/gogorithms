package searching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSelect(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          []int
		k              int
		expectedOutput int
	}{
		{
			name:           "example_one",
			input:          []int{7, 2, 4, 15, 9, 5},
			k:              2,
			expectedOutput: 4,
		},
		{
			name:           "example_two",
			input:          []int{7, 2, 4, 15, 9, 5},
			k:              5,
			expectedOutput: 9,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := Quickselect(tt.input, tt.k)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}
