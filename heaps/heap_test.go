package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntHeap(t *testing.T) {
	t.Parallel()

	var arrayOne = []int{6, 1, 5, 8, 9}
	var arrayTwo = []int{9, 1, 1, 8, 9}

	minHeap := BuildIntHeapFromArray(arrayOne, MinHeap)
	maxHeap := BuildIntHeapFromArray(arrayTwo, MaxHeap)

	// Test min heap.
	assert.Equal(t, 1, minHeap.Peek())
	assert.Equal(t, 1, minHeap.Peek())

	v := minHeap.Remove()
	assert.Equal(t, 1, v)
	assert.Equal(t, 5, minHeap.Peek())

	minHeap.Add(7)

	v = minHeap.Remove()
	assert.Equal(t, 5, v)
	assert.Equal(t, 6, minHeap.Peek())

	v = minHeap.Remove()
	assert.Equal(t, 6, v)
	assert.Equal(t, 7, minHeap.Peek())

	// Test max heap.
	assert.Equal(t, 9, maxHeap.Peek())
	assert.Equal(t, 9, maxHeap.Peek())

	v = maxHeap.Remove()
	assert.Equal(t, 9, v)
	assert.Equal(t, 9, maxHeap.Peek())

	v = maxHeap.Remove()
	assert.Equal(t, 9, v)
	assert.Equal(t, 8, maxHeap.Peek(), maxHeap.values)

	maxHeap.Add(99)
	assert.Equal(t, 99, maxHeap.Peek())
}
