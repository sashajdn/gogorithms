package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var knapsack01Checks = []func(items [][]int, capacity int) ([][]int, int){
	KnapsackO1_TopDown,
	KnapsackO1_BottomUp,
}

func TestKnapsack01(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		items          [][]int
		capacity       int
		expectedItems  [][]int
		expectedProfit int
	}{
		{
			name: "example_one",
			items: [][]int{
				{100, 100},
				{2, 2},
			},
			capacity:       1,
			expectedItems:  [][]int{},
			expectedProfit: 0,
		},
		{
			name: "example_two",
			items: [][]int{
				{100, 100},
				{50, 2},
				{51, 98},
			},
			capacity: 100,
			expectedItems: [][]int{
				{50, 2},
				{51, 98},
			},
			expectedProfit: 101,
		},
		{
			name: "example_three",
			items: [][]int{
				{1, 2},
				{4, 3},
				{5, 6},
				{6, 7},
			},
			capacity: 10,
			expectedItems: [][]int{
				{4, 3},
				{6, 7},
			},
			expectedProfit: 10,
		},
		{
			name: "example_four",
			items: [][]int{
				{11, 16},
				{4, 3},
				{10, 5},
				{5, 6},
				{6, 7},
			},
			capacity: 22,
			expectedItems: [][]int{
				{4, 3},
				{10, 5},
				{5, 6},
				{6, 7},
			},
			expectedProfit: 25,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, c := range knapsack01Checks {
				includedItems, profit := c(tt.items, tt.capacity)

				assert.Equal(t, tt.expectedItems, includedItems)
				assert.Equal(t, tt.expectedProfit, profit)
			}
		})
	}
}
