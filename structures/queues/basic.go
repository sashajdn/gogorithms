package queues

import (
	"math"
	"sync"
)

// BasicQueue a very basic FIFO queue with the option of a max length. Thread Safe.
type BasicQueue struct {
	q         []interface{}
	maxLength int
	sync.Mutex
}

// NewBasicQueue is a BasicQueue factory method
// a maxLength <= 0 results in an "infite" size queue (the infite size is the max int32 in go)
func NewBasicQueue(maxLength int) *BasicQueue {
	ml := maxLength
	if maxLength <= 0 {
		ml = math.MaxInt32
	}
	return &BasicQueue{
		maxLength: ml,
	}
}

// Pop removes & returns the last item in the queue
func (q *BasicQueue) Pop() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()
	if len(q.q) == 0 {
		return nil, false
	}
	r := q.q[0]
	q.q = q.q[1:]
	return r, true
}

// Enqueue appends item `v` to the end of the queue.
func (q *BasicQueue) Enqueue(v interface{}) bool {
	q.Lock()
	defer q.Unlock()
	if len(q.q) >= q.maxLength {
		return false
	}
	q.q = append(q.q, v)
	return true
}

// EnqueueMany appends many items `a` to the end of the queue in the sequence of which there are given
// Returns an integer that represents how many items were enqueued.
func (q *BasicQueue) EnqueueMany(a []interface{}) int {
	if len(a) == 0 {
		return 0
	}
	if !q.Enqueue(a[0]) {
		return 0
	}
	return 1 + q.EnqueueMany(a[1:])
}

// Size returns how many items are currently in the queue
func (q *BasicQueue) Size() int {
	q.Lock()
	defer q.Unlock()
	return len(q.q)
}

// Peek returns the last item of the queue without removing
func (q *BasicQueue) Peek() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()
	if len(q.q) == 0 {
		return nil, false
	}
	return q.q[0], true
}
