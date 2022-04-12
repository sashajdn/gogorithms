package linkedlists

// ReverseKGroup ...
//
// T -> O(n)
// S -> O(1)
func ReverseKGroup(head *LinkedList, k int) *LinkedList {
	if head == nil {
		return nil
	}

	var (
		current, leader             = head, head
		newHead, kTail, reverseHead *LinkedList
	)
	for leader != nil {
		for i := 0; i < k; i++ {
			if leader == nil {
				if kTail != nil {
					kTail.Next = current
				}

				if newHead != nil {
					return newHead
				}

				return head
			}

			leader = leader.Next
		}

		reverseHead = reverseK(current, k)
		if newHead == nil {
			newHead = reverseHead
		}

		if kTail != nil {
			kTail.Next = reverseHead
		}

		kTail = current
		current = leader
	}

	if kTail != nil {
		kTail.Next = current
	}

	if newHead != nil {
		return newHead
	}

	return head
}

func reverseK(node *LinkedList, k int) *LinkedList {
	var (
		previous *LinkedList
		current  = node
	)
	for i := 0; i < k; i++ {
		next := current.Next
		current.Next = previous
		previous, current = current, next
	}

	return previous
}
