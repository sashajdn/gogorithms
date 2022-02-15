package heaps

// Heap ...
type Heap interface {
	Peek() int
	Insert(value int)
	Pop() int
}
