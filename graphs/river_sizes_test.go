package graphs

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRiverSizes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		rivers             [][]int
		expectedRiverSizes []int
	}{
		{
			name: "full-matrix",
			rivers: [][]int{
				{1, 0, 0, 1, 0},
				{1, 0, 1, 0, 0},
				{0, 0, 1, 0, 1},
				{1, 0, 1, 0, 1},
				{1, 0, 1, 1, 0},
			},
			expectedRiverSizes: []int{1, 2, 2, 2, 5},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := RiverSizes(tt.rivers)
			assertEqual(t, tt.expectedRiverSizes, res)
		})
	}
}

func assertEqual(t *testing.T, expected, got []int) {
	sort.Ints(expected)
	sort.Ints(got)
	if len(expected) != len(got) {
		t.Fatalf("Failed: expected %v, got %v", expected, got)
	}

	assert.Equal(t, expected, got)
}
