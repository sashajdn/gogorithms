package linkedlists

import "container/heap"

// T -> O(n * log(k))
// S -> O(k)
func MergeKLists_Heap(lists []*LinkedList) *LinkedList {
	var (
		pointers = make([]*LinkedList, 0, len(lists))
		pq       = &PriorityQueue{}
	)

	// T -> O(k)
	// S -> O(k)
	for i, l := range lists {
		heap.Push(pq, &HeapItem{
			Node:  l,
			Index: i,
		})

		pointers = append(pointers, lists[i])
	}

	// T -> O(nlog(k))
	// S -> O(k)
	var sentinel = &LinkedList{}
	var current = sentinel
	for pq.Len() > 0 {
		next := heap.Pop(pq)
		heapItem := next.(*HeapItem)

		current.Next = heapItem.Node
		current = current.Next

		toPush := pointers[heapItem.Index]
		if toPush != nil {
			heap.Push(pq, &HeapItem{
				Node:  toPush,
				Index: heapItem.Index,
			})

			pointers[heapItem.Index] = pointers[heapItem.Index].Next
		}
	}

	return sentinel.Next
}

type HeapItem struct {
	Index int
	Node  *LinkedList
}

type PriorityQueue []*HeapItem

func (p PriorityQueue) Len() int           { return len(p) }
func (p PriorityQueue) Less(i, j int) bool { return p[i].Node.Value < p[j].Node.Value }
func (p PriorityQueue) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *PriorityQueue) Push(v interface{}) {
	hi := v.(*HeapItem)
	*p = append(*p, hi)
}
func (p *PriorityQueue) Pop() interface{} {
	v := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return v
}
