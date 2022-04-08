package priority

import (
	"container/heap"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	array := []int{5, 3, 3, 1, 6, 7, 8, 3, 2, 5, 3}
	pq := &PriorityQueue{}

	for _, v := range array {
		heap.Push(pq, v)
	}

	var collector []int
	for pq.Len() > 0 {
		v := heap.Pop(pq)
		vi := v.(int)
		collector = append(collector, vi)
	}

	sort.Ints(array)

	assert.Equal(t, array, collector)
}
