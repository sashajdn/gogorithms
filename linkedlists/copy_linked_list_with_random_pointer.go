package linkedlists

type RandomLinkedList struct {
	Value  int
	Next   *RandomLinkedList
	Random *RandomLinkedList
}

// CopyRandomList ...
//
// T -> O(n)
// S -> O(1)
func CopyRandomList(head *RandomLinkedList) *RandomLinkedList {
	if head == nil {
		return nil
	}

	// Weave linked lists.
	//
	// T -> O(n)
	// S -> O(1)
	var current = head
	for current != nil {
		next := current.Next

		copiedNode := &RandomLinkedList{
			Value: current.Value,
			Next:  next,
		}

		current.Next = copiedNode
		current = head
	}

	// Add random pointers to copied list.
	//
	// T -> O(n)
	// S -> O(1 )
	current = head
	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}

		current = current.Next.Next
	}

	// Unweave lists.
	//
	// T -> O(n)
	// S -> O(1)
	var (
		newHead          = head.Next
		original, copied = head, newHead
	)
	for original != nil {
		original.Next = original.Next.Next
		if copied.Next != nil {
			copied.Next = copied.Next.Next
		}

		original = original.Next
		copied = copied.Next
	}

	return copied
}
