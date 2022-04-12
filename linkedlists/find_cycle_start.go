package linkedlists

// FindCycleStart ...
//
// T -> O(n)
// S -> O(1)
func FindCycleStart(head *LinkedList) *LinkedList {
	if head == nil {
		return nil
	}

	var slow, fast = head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	var slow2 = head
	for slow != slow2 {
		slow, slow2 = slow.Next, slow2.Next
	}

	return slow
}
