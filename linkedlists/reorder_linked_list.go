package linkedlists

// ReOrderLinkedList ...
//
// T -> O(n)
// S -> O(1)
func ReOrderLinkedList(head *LinkedList) {
	if head == nil || head.Next == nil {
		return
	}

	var (
		prev       *LinkedList
		slow, fast = head, head
	)
	for fast != nil && fast.Next != nil {
		prev = slow
		slow, fast = slow.Next, fast.Next.Next
	}
	prev.Next = nil

	var (
		previous *LinkedList
		current  = slow
	)
	for current != nil {
		next := current.Next
		current.Next = previous
		previous, current = current, next
	}

	var (
		sent          = &LinkedList{}
		first, second = head, previous
	)
	for second != nil {
		sent.Next = first
		first = first.Next
		sent = sent.Next

		sent.Next = second
		second = second.Next
		sent = sent.Next
	}
}
