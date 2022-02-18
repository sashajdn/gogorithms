package linkedlists

// Reverse ...
//
// T -> O(n) where n is the length of the linked list.
// S -> O(1)
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
