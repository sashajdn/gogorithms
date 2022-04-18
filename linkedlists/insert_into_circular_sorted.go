package linkedlists

// InsertIntoSortedCircular ...
//
// T -> O(n) where `n` is the number of elements in the linkedlist
// S -> O(1)
func InsertIntoSortedCircular(head *LinkedList, value int) *LinkedList {
	node := &LinkedList{
		Value: value,
	}

	if head == nil {
		node.Next = node
		return node
	}

	var current, next = head, head.Next
	defer func() {
		current.Next = node
		node.Next = next
	}()

	for {
		switch {
		case current == next:
			return head
		case current.Value > next.Value && (value > current.Value || value < next.Value):
			return head
		case value == next.Value:
			return head
		case value > current.Value && value < next.Value:
			return head
		case next == head:
			return head
		default:
			current, next = current.Next, next.Next
		}
	}
}
