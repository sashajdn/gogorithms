package linkedlists

// RemoveDuplicatesFromLinkedList O(T) -> O(n), O(S) -> O(1)
func RemoveDuplicatesFromLinkedList(linkedlist *LinkedList) *LinkedList {
	// We can assume the values of the linked list are in sorted order.
	head := linkedlist
	recurseAndRemove(head)
	return head
}

func recurseAndRemove(head *LinkedList) {
	if head.Next == nil {
		return
	}
	if head.Value == head.Next.Value {
		head.Next = head.Next.Next
		recurseAndRemove(head)
	}
	if head.Next == nil {
		return
	}
	recurseAndRemove(head.Next)
}
