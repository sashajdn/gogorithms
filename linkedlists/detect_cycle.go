package linkedlists

// DetectCycleTwoPointer ...
//
// T -> O(n) where `n` is the number of nodes in the linked list.
// S -> O(1)
func DetectCycleTwoPointer(head *LinkedList) bool {
	if head == nil || head.Next == nil {
		return false
	}

	p, pp := head, head.Next
	for p != nil && pp != nil {
		if p == pp {
			return true
		}

		if pp.Next == nil {
			return false
		}

		p, pp = p.Next, pp.Next.Next
	}

	return false
}

// DetectCycleHashSet ...
//
// T -> O(n) where `n` is the number of nodes in the linked list.
// S -> O(n)
func DetectCycleHashSet(head *LinkedList) bool {
	var (
		hs      = map[int]struct{}{}
		current = head
	)

	for current != nil {
		if _, ok := hs[current.Value]; ok {
			return true
		}

		current = current.Next
	}

	return false
}

// DetectCycleRecursive ...
//
// T -> O(n) where `n` is the number of nodes in the linked list.
// S -> O(n)
func DetectCycleRecursive(head *LinkedList) bool {
	var hs = map[int]struct{}{}

	var f func(*LinkedList) bool
	f = func(node *LinkedList) bool {
		if node == nil {
			return false
		}

		if _, ok := hs[node.Value]; ok {
			return true
		}

		return f(node.Next)
	}

	return f(head)
}
