package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sameBSTCheckers = []func(arrayOne, arrayTwo []int) bool{
		SameBsts,
		SameBsts_Better,
	}
)

func TestSameBST(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		arrayOne, arrayTwo []int
		expectedOutput     bool
	}{
		{
			name:           "example_one",
			arrayOne:       []int{10, 15, 8, 12, 94, 81, 5, 2, 11},
			arrayTwo:       []int{10, 8, 5, 15, 2, 12, 11, 94, 81},
			expectedOutput: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, check := range sameBSTCheckers {
				res := check(tt.arrayOne, tt.arrayTwo)

				assert.Equal(t, tt.expectedOutput, res)
			}
		})
	}
}
