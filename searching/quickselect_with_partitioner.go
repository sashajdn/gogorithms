package searching

import (
	"math/rand"
	"time"
)

type QuickSelectType int

const (
	QuickSelectWithPartitionerFirstElement QuickSelectType = iota + 1
	QuickSelectWithPartitionerRandom
	QuickSelectWithPartitionerMedian
)

func (q QuickSelectType) Select(array []int, start, end, target int) int {
	pivot := q.selectPivot(array, start, end)
	l, r := pivot+1, end

	for l <= r {
		if array[l] > array[pivot] && array[r] < array[pivot] {
			q.swap(array, l, r)
		}

		if array[l] <= array[pivot] {
			l++
		}
		if array[r] >= array[pivot] {
			r--
		}
	}
	q.swap(array, pivot, r)

	if target == r {
		return array[r]
	}

	switch {
	case r < target:
		return q.Select(array, r+1, end, target)
	default:
		return q.Select(array, start, r-1, target)
	}
}

func (q QuickSelectType) selectPivot(array []int, start, end int) int {
	switch q {
	case QuickSelectWithPartitionerFirstElement:
		return start
	case QuickSelectWithPartitionerMedian:
		if end-start < 3 {
			return start
		}

		if (end-start)%2 == 0 {
			return ((end - start) / 2) - 1
		}

		return (end - start) / 2
	case QuickSelectWithPartitionerRandom:
		rand.Seed(time.Now().UnixNano())
		p := start + rand.Intn((end-start)+1)
		q.swap(array, p, start)
		return start
	}
	return start
}

func (q QuickSelectType) swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

// QuickSelectWithPartitioner ...
func QuickSelectWithPartitioner(array []int, k int, parititioner QuickSelectType) int {
	return parititioner.Select(array, 0, len(array)-1, k-1)
}
