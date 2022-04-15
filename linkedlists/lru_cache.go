package linkedlists

type LRUCacheNode struct {
	Key, Value     int
	Previous, Next *LRUCacheNode
}

type LRUCache struct {
	Head, Tail *LRUCacheNode
	reference  map[int]*LRUCacheNode
	capacity   int
}

func (l *LRUCache) isOverCapacity() bool {
	return len(l.reference) > l.capacity
}

// Get ...
//
// T -> O(1)
// S -> O(1)
func (l *LRUCache) Get(key int) int {
	node, ok := l.reference[key]
	if !ok {
		return -1
	}

	if node == l.Head {
		return node.Value
	}

	previous, next := node.Previous, node.Next

	if previous != nil {
		previous.Next = next
	}
	if next != nil {
		next.Previous = previous
	}

	if node == l.Tail {
		l.Tail = l.Tail.Previous
	}

	node.Next = l.Head
	l.Head.Previous = node
	l.Head = node

	return node.Value
}

// Put ...
//
// T -> O(1)
// S -> O(1)
func (l *LRUCache) Put(key, value int) {
	if node, ok := l.reference[key]; ok {
		node.Value = value
		l.Get(key) // Force reorder.
		return
	}

	node := &LRUCacheNode{
		Key:   key,
		Value: value,
		Next:  l.Head,
	}
	l.reference[key] = node

	if l.isOverCapacity() {
		if l.Tail.Previous != nil {
			l.Tail.Previous.Next = nil
		}
		l.Tail.Previous = nil
		l.Tail = l.Tail.Previous

		delete(l.reference, key)
	}

	if l.Head != nil {
		l.Head.Previous = node
	}
	l.Head = node

	if l.Tail == nil {
		l.Tail = node
	}
}

// NewLRUCache ...
//
// T -> O(1)
// S -> O(1)
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		reference: make(map[int]*LRUCacheNode),
		capacity:  capacity,
	}
}
