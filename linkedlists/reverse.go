package linkedlists

// Reverse ...
//
// T -> O(n) where n is the length of the linked list.
// S -> O(n)
func Reverse(head *LinkedList) *LinkedList {
	if head == nil {
		return nil
	}

	return reverse(head, nil)
}

func reverse(curr, prev *LinkedList) *LinkedList {
	if curr.Next == nil {
		curr.Next = prev
		return curr
	}

	next := curr.Next
	curr.Next = prev
	return reverse(next, curr)
}

// ReverseIterative ...
// T -> O(n) where `n` is the length of the linked list.
// S -> O(1)
func ReverseIterative(head *LinkedList) *LinkedList {
	var (
		previous *LinkedList
		current  = head
	)
	for current != nil {
		next := current.Next
		current.Next = previous
		previous, current = current, next
	}
	return previous
}
