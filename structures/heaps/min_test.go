package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	t.Parallel()

	input := []int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41}
	expectedHeap := []int{-5, 2, 6, 7, 8, 8, 24, 391, 24, 56, 12, 24, 48, 41}

	// Build Array.
	h := BuildHeapFromArray(input)
	assert.Equal(t, expectedHeap, h.d)

	// Peek.
	assert.Equal(t, expectedHeap[0], h.Peek())

	// Remove.
	v := h.Remove()
	assert.Equal(t, expectedHeap[0], v)
	assert.Equal(t, expectedHeap[1], h.Peek())

	// Insert.
	h.Insert(-100)
	assert.Equal(t, expectedHeap[0], -100)
}
