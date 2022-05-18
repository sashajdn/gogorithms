package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var maxProfitWithKTransactionsChecks = []func(prices []int, k int) int{
	MaxProfitWithKTransactions_BottomUp,
	MaxProfitWithKTransactions_TopDown2D,
}

func TestMaxProfitWithKTransactions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		prices         []int
		k              int
		expectedProfit int
	}{
		{
			name:           "example_zero",
			prices:         []int{5, 11, 3, 50, 60, 90},
			k:              1,
			expectedProfit: 87,
		},
		{
			name:           "example_one",
			prices:         []int{5, 11, 3, 50, 60, 90},
			k:              2,
			expectedProfit: 93,
		},
		{
			name:           "example_two",
			prices:         []int{5, 11, 3, 50, 60, 90, 2, 4, 100, 23, 17, 3},
			k:              4,
			expectedProfit: 191,
		},
		{
			name:           "example_three",
			prices:         []int{1, 25, 24, 23, 12, 36, 14, 40, 31, 41, 5},
			k:              4,
			expectedProfit: 84,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range maxProfitWithKTransactionsChecks {
				res := c(tt.prices, tt.k)

				assert.Equal(t, tt.expectedProfit, res)
			}
		})
	}
}
