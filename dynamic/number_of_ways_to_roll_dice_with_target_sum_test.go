package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var numberOfWaysToRollDiceCheck = []func(d, f, target int) int{
	NumberOfWaysToRollDice_TopDown,
	NumberOfWaysToRollDice_BottomUp,
}

func TestNumberOfWaysToRollDice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		d, f, target   int
		expectedOutput int
	}{
		{
			name:           "example_one",
			d:              1,
			f:              6,
			target:         3,
			expectedOutput: 1,
		},
		{
			name:           "example_two",
			d:              2,
			f:              6,
			target:         7,
			expectedOutput: 6,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range numberOfWaysToRollDiceCheck {
				res := c(tt.d, tt.f, tt.target)

				assert.Equal(t, tt.expectedOutput, res)
			}
		})
	}
}
