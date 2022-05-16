package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var coinChangeChecks = []func(coins []int, amount int) int{
	CoinChange_BottomUp,
	CoinChange_TopDown2D,
	CoinChange_TopDown1D,
}

func TestCoinChange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		coins          []int
		amount         int
		expectedAmount int
	}{
		{
			name:           "example_one",
			coins:          []int{1, 2, 5},
			amount:         11,
			expectedAmount: 3,
		},
		{
			name:           "example_two",
			coins:          []int{2, 5, 6},
			amount:         14,
			expectedAmount: 3,
		},
		{
			name:           "example_three",
			coins:          []int{2, 3, 4, 5, 6},
			amount:         26,
			expectedAmount: 5,
		},
		{
			name:           "example_four",
			coins:          []int{2, 3, 4, 5, 6},
			amount:         69,
			expectedAmount: 12,
		},
		{
			name:           "example_five",
			coins:          []int{1, 2, 3, 4, 5, 6},
			amount:         69,
			expectedAmount: 12,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range coinChangeChecks {
				res := c(tt.coins, tt.amount)

				assert.Equal(t, tt.expectedAmount, res)
			}
		})
	}
}
