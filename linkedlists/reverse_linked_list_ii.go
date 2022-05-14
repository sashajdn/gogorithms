package linkedlists

// ReverseLinkedListII ...
//
// T -> O(n)
// S -> O(1)
func ReverseLinkedListII(head *LinkedList, left, right int) *LinkedList {
	if head == nil {
		return nil
	}

	var (
		follower        *LinkedList
		current, leader = head, head
		count           = 1
	)
	for count < left {
		count++
		follower, current, leader = current, current.Next, leader.Next
	}
	for count < right+1 {
		count++
		leader = leader.Next
	}

	reversedNode := kReverseLinkedList(leader, current, right-left)
	if follower == nil {
		return reversedNode
	}
	follower.Next = reversedNode

	return head
}

func kReverseLinkedList(previous, current *LinkedList, k int) *LinkedList {
	for current != nil && k >= 0 {
		next := current.Next
		current.Next = previous
		previous, current = current, next
		k--
	}

	return previous
}
