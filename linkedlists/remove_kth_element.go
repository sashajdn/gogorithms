package linkedlists

// RemoveKthNodeFromEnd removes the kth element from the end of a singly linked list.
// It is done in place.
//
// T -> O(n) where n the length of the linked list.
// S -> O(1)
func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	if k == 0 {
		return
	}

	// O(n) - we iterate on the linked list until the `Next` value is empty.
	switch l := length(head); {
	case k >= l:
		head.Value = head.Next.Value
		head.Next = head.Next.Next
	default:
		// O(n) - we recurse on the list n-k times.
		kthPrevNode := recurse(head, l-k-1)
		kthPrevNode.Next = kthPrevNode.Next.Next
	}

}

func length(head *LinkedList) int {
	if head == nil {
		return 0
	}

	return 1 + length(head.Next)
}

func recurse(head *LinkedList, i int) *LinkedList {
	if i == 0 {
		return head
	}

	if head == nil {
		return head
	}

	return recurse(head.Next, i-1)
}
