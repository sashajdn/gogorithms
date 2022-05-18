package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var decodeWaysCheck = []func(s string) int{
	DecodeWays_BottomUp,
	DecodeWays_TopDown,
	DecodeWays_TopDownTwoPointer,
}

func TestDecodeWays(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                 string
		input                string
		expectedNubmerOfWays int
	}{
		{
			name:                 "example_one",
			input:                "226",
			expectedNubmerOfWays: 3,
		},
		{
			name:                 "example_two",
			input:                "1223582",
			expectedNubmerOfWays: 5,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range decodeWaysCheck {
				res := c(tt.input)

				assert.Equal(t, tt.expectedNubmerOfWays, res)
			}
		})
	}
}
