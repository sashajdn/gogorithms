package heaps

type HeapType int

const (
	MinHeap HeapType = iota + 1
	MaxHeap
)

func (h HeapType) Comparision() Comparision {
	switch h {
	case MinHeap:
		return func(a, b int) bool { return a < b }
	default:
		return func(a, b int) bool { return a > b }
	}
}

func BuildIntHeapFromArray(array []int, heapType HeapType) *IntHeap {
	h := &IntHeap{
		comp:   heapType.Comparision(),
		values: array,
	}

	h.buildIntHeap()

	return h
}

type Comparision func(a, b int) bool

type IntHeap struct {
	comp   Comparision
	values []int
}

func (h *IntHeap) Add(value int) {
	h.values = append(h.values, value)
	h.siftUp()
}

func (h *IntHeap) Remove() int {
	v := h.values[0]
	h.swap(0, len(h.values)-1)
	h.values = h.values[:len(h.values)-1]
	h.siftDown(0, len(h.values)-1)
	return v
}

func (h *IntHeap) Peek() int {
	if len(h.values) < 1 {
		return -1
	}
	return h.values[0]
}

func (h *IntHeap) buildIntHeap() {
	for k := (len(h.values) - 2) / 2; k >= 0; k-- {
		h.siftDown(k, len(h.values)-1)
	}
}

func (h *IntHeap) siftUp() {
	currentIdx := len(h.values) - 1

	for currentIdx > 0 {
		parentIdx := (currentIdx - 1) / 2

		if h.comp(h.values[currentIdx], h.values[parentIdx]) {
			h.swap(currentIdx, parentIdx)
			currentIdx = parentIdx
			continue
		}

		return
	}
}

func (h *IntHeap) siftDown(currentIdx, endIdx int) {
	if currentIdx >= endIdx {
		return
	}

	lci, rci := (2*currentIdx)+1, (2*currentIdx)+2

	minIdx := h.minIdxFromValue(h.values, lci, rci)
	if minIdx < 0 {
		return
	}

	if !h.comp(h.values[minIdx], h.values[currentIdx]) {
		return
	}

	h.swap(minIdx, currentIdx)
	h.siftDown(minIdx, endIdx)
}

func (h *IntHeap) swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *IntHeap) minIdxFromValue(array []int, l, r int) int {
	if l >= len(array) {
		return -1
	}

	if r >= len(array) {
		return l
	}

	if h.comp(array[l], array[r]) {
		return l
	}

	return r
}
