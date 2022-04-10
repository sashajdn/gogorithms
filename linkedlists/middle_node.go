package linkedlists

// MiddleNode ...
//
// T -> O(n)
// S -> O(1)
func MiddleNode(head *LinkedList) *LinkedList {
	var leader, follower = head, head
	for leader != nil && leader.Next != nil {
		leader, follower = leader.Next.Next, follower.Next
	}
	return follower
}
