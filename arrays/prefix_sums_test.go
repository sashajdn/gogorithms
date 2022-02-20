package arrays

import (
	"testing"
)

func TestPrefixSums(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		weights []int
	}{
		{
			name:    "example_one",
			weights: []int{2, 3, 10, 4, 8},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			wr := NewWeightedRand(tt.weights)

			var counter = map[int]int{}
			for i := 0; i < 1000; i++ {
				v := wr.PickIndex()
				counter[v]++
			}
		})
	}

}
