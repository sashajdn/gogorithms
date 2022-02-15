package heaps

type ContinuousMedianHandler struct {
	Median float64

	lesser *Heap
	higher *Heap
}

func NewContinuousMedianHandler() *ContinuousMedianHandler {
	// Write your code here.
	return &ContinuousMedianHandler{
		lesser: NewHeap(moreThan),
		higher: NewHeap(lessThan),
	}
}

func (handler *ContinuousMedianHandler) Insert(number int) {
	if handler.lesser.Length() < 1 {
		handler.lesser.Insert(number)
		return
	}

	switch {
	case number < handler.lesser.Peek():
		handler.lesser.Insert(number)
	default:
		handler.higher.Insert(number)
	}

	handler.rebalance()

	if handler.lesser.Length() == handler.higher.Length() {
		handler.Median = (float64(handler.higher.Peek()) + float64(handler.lesser.Peek())) / 2
		return
	}

	if handler.lesser.Length() > handler.higher.Length() {
		handler.Median = float64(handler.lesser.Peek())
		return
	}

	handler.Median = float64(handler.higher.Peek())
}

func (handler *ContinuousMedianHandler) GetMedian() float64 {
	return handler.Median
}

func (handler *ContinuousMedianHandler) rebalance() {
	larger, smaller := handler.getLargerSmaller()

	if larger.Length()-smaller.Length() < 2 {
		return
	}

	smaller.Insert(larger.Remove())
}

func (handler *ContinuousMedianHandler) getLargerSmaller() (*Heap, *Heap) {
	if handler.higher.Length() > handler.lesser.Length() {
		return handler.higher, handler.lesser
	}

	return handler.lesser, handler.higher
}

func NewHeap(comp func(a, b int) bool) *Heap
func lessThan(a, b int) bool { return a < b }
func moreThan(a, b int) bool { return a > b }

type Heap struct {
	values      []int
	comparision func(a, b int) bool
}

func (h *Heap) Length() int {
	return len(h.values)
}

func (h *Heap) Peek() int {
	if h.Length() < 1 {
		return -1
	}

	return h.values[0]
}

func (h *Heap) Remove() int {
	if h.Length() < 1 {
		return -1
	}

	v := h.Peek()
	h.swap(0, h.Length()-1)
	h.values = h.values[:h.Length()-1]
	h.siftDown(0, h.Length()-1)
	return v
}

func (h *Heap) Insert(value int) {
	h.values = append(h.values, value)
	h.siftUp()
}

func (h *Heap) siftUp() {
	currentIdx := h.Length() - 1
	for currentIdx > 0 {
		parentIdx := (currentIdx - 1) / 2
		if !h.comparision(parentIdx, currentIdx) {
			return
		}

		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
	}
}

func (h *Heap) siftDown(currentIdx, endIdx int) {
	if (2*currentIdx + 1) > h.Length()-1 {
		return
	}

	lidx, ridx := (2*currentIdx)+1, (2*currentIdx)+2
	compIndex := h.compIndexFromValue(lidx, ridx)
	if compIndex < 0 {
		return
	}

	if h.comparision(h.values[compIndex], h.values[currentIdx]) {
		h.swap(compIndex, currentIdx)
		h.siftDown(compIndex, h.Length()-1)
	}
}

func (h *Heap) compIndexFromValue(i, j int) int {
	if j > h.Length()-1 && i > h.Length()-1 {
		return -1
	}

	if j > h.Length()-1 {
		return i
	}

	lv, rv := h.values[i], h.values[j]
	if h.comparision(lv, rv) {
		return i
	}

	return j
}

func (h *Heap) swap(i, j int) {
	if i < h.Length() && j < h.Length() {
		h.values[i], h.values[j] = h.values[j], h.values[i]
	}
}
