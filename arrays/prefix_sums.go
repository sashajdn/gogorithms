package arrays

import (
	"math/rand"
	"time"
)

type WeightedRand struct {
	prefixSums []int
	totalSum   int
}

func NewWeightedRand(weights []int) *WeightedRand {
	var (
		totalSum   int
		prefixSums []int
	)
	for _, w := range weights {
		totalSum += w
		prefixSums = append(prefixSums, totalSum)
	}

	return &WeightedRand{
		prefixSums: prefixSums,
		totalSum:   totalSum,
	}
}

func (w *WeightedRand) PickIndex() int {
	rand.Seed(time.Now().UnixNano())
	target := int(float64(w.totalSum)*rand.Float64()) + 1
	return w.binarySearch(0, len(w.prefixSums), target)
}

func (w *WeightedRand) binarySearch(l, h, target int) int {
	if l >= h {
		return l
	}

	mid := l + (h-l)/2

	switch {
	case target > w.prefixSums[mid]:
		return w.binarySearch(mid+1, h, target)
	default:
		return w.binarySearch(l, mid, target)
	}
}
