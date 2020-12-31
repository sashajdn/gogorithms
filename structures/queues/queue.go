package queues

type Queue interface {
	// Returns the last value in the Queue. Removes item.
	Pop() (interface{}, bool)
	// Enqueues a new item onto the Queue
	Enqueue(v interface{}) bool
	// Enqueues many items to the Queue
	EnqueueMany(a []interface{}) bool
	// Returns the size of the Queue
	Size() int
	// Returns the last items in the Queue, without removing it from said Queue
	Peek() (interface{}, bool)
}
