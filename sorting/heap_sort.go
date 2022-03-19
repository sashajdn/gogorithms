package sorting

// MaxHeap ...
type MaxHeap []int

// Heapify ...
//
// T -> O(nlog(n)) where `n` is the number of items in the array.
// S -> O(1)
func (h MaxHeap) Heapify() {
	rightMostChild := len(h) - 1
	rightMostParent := (rightMostChild - 1) / 2

	for k := rightMostParent; k >= 0; k-- {
		h.siftDown(k, rightMostChild)
	}
}

// Sort ...
// T -> O(nlog(n)) where `n` is the number of items in the array.
// S -> O(1)
func (h MaxHeap) Sort() {
	rightMostChild := len(h) - 1
	for k := rightMostChild; k >= 1; k-- {
		h.swap(0, k)
		h.siftDown(0, k-1)
	}
}

// T -> O(log(n)) where `n` is the number of items in the array.
// S -> O(1)
func (h MaxHeap) siftDown(startIndex, endIndex int) {
	var parentIndex = startIndex

	for (2*parentIndex)+1 <= endIndex {
		maxChildIndex := h.maxChildIndex(parentIndex, endIndex)
		if maxChildIndex == -1 {
			return
		}

		if h[maxChildIndex] <= h[parentIndex] {
			return
		}

		h.swap(parentIndex, maxChildIndex)
		parentIndex = maxChildIndex
	}
}

func (h MaxHeap) maxChildIndex(parentIndex, endIndex int) int {
	l, r := 2*parentIndex+1, 2*parentIndex+2
	if l > endIndex {
		return -1
	}

	if r > endIndex {
		return l
	}

	if h[l] >= h[r] {
		return l
	}

	return r
}

func (h MaxHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// HeapSort ...
//
// T -> O(nlog(n)) where `n` is the number of items in the array; note this is actually 2 * nlog(n) but we can ignore the constant factor.
// S -> O(1) as in place.
func HeapSort(array []int) []int {
	heap := MaxHeap(array)
	heap.Heapify()
	heap.Sort()
	return []int(heap)
}
