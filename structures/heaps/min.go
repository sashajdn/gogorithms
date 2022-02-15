package heaps

// MinHeap ...
type MinHeap []int

// NewMinHeap ...
//
// T -> O(n)
// S -> O(1)
func NewMinHeap(array []int) *MinHeap {
	h := MinHeap(array)
	h.BuildHeap(array)
	return &h
}

// Peek ...
//
// T -> O(1)
// S -> O(1)
func (h *MinHeap) Peek() int {
	if len(*h) < 1 {
		return -1
	}

	return (*h)[0]
}

// Insert ...
//
// T -> O(log(n))
// S -> O(1)
func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	h.siftUp()
}

// Remove ...
//
// T -> O(log(n))
// S -> O(1)
func (h *MinHeap) Remove() int {
	v := h.Peek()

	h.swap(0, len(*h)-1)
	*h = (*h)[:len(*h)-1]

	h.siftDown(0, len(*h)-1)

	return v
}

// BuildHeap ...
//
// T -> O(n)
// S -> O(1)
func (h *MinHeap) BuildHeap() {
	lastIdx := (len(*h) - 2) / 2
	for i := lastIdx; i >= 0; i-- {
		h.siftDown(i, len(*h)-1)
	}
}

// siftUp ...
// T -> O(logn)
// S -> O(1)
func (h *MinHeap) siftUp() {
	currentIdx := len(*h) - 1

	for currentIdx >= 0 {
		parentIdx := (currentIdx - 1) / 2
		if (*h)[currentIdx] < (*h)[parentIdx] {
			h.swap(currentIdx, parentIdx)
			currentIdx = parentIdx
		}

		return
	}
}

// siftDown ...
// T -> O(logn)
// S -> O(1)
func (h *MinHeap) siftDown(currentIndex, endIndex int) {
	var (
		lidx, ridx int
		length     = len(*h)
	)

	if (2*currentIndex)+1 >= length {
		return
	}

	if (2*currentIndex)+1 < length {
		lidx = (2 * currentIndex) + 1
	}

	if (2*currentIndex)+2 < length {
		ridx = (2 * currentIndex) + 2
	}

	minLR := h.minIndexFromValues(lidx, ridx)

	if (*h)[currentIndex] < minLR {
		h.swap(currentIndex, minLR)
		h.siftDown(minLR, endIndex)
	}
}

func (h *MinHeap) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) minIndexFromValues(lidx, ridx int) int {
	// Since we check right to left & already return before this gets called if lidx not in bounds.
	if ridx == 0 {
		return lidx
	}

	if (*h)[lidx] < (*h)[ridx] {
		return lidx
	}

	return ridx
}
