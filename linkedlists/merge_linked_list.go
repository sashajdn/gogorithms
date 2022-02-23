package linkedlists

// MergeLinkedLists_Iterative ...
//
// T: O(a + b) where a is the length of the first linked list passed; and b the second.
// S: O(1)
func MergeLinkedLists_Iterative(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	if headOne == nil && headTwo == nil {
		return nil
	}

	var sentinel = &LinkedList{}
	a, b, prev := minNode(headOne, headTwo)
	sentinel.Next = prev

	for a != nil || b != nil {
		if a == nil {
			prev.Next = b
			b = b.Next
			prev = prev.Next
			continue
		}

		if b == nil {
			prev.Next = a
			a = a.Next
			prev = prev.Next
			continue
		}

		if a.Value >= b.Value {
			prev.Next = b
			b = b.Next
			prev = prev.Next
			continue
		}

		prev.Next = a
		a = a.Next
		prev = prev.Next
	}

	return sentinel.Next
}

// MergeLinkedLists_Recursive ...
//
// T -> O(a + b) where a is the length of the first linked list passed & b the second.
// S -> O(a + b) due to the recursive stack.
func MergeLinkedLists_Recursive(headOne, headTwo *LinkedList) *LinkedList {
	if headOne == nil && headTwo == nil {
		return nil
	}

	var sentinel = &LinkedList{}
	mergeLists(sentinel, headOne, headTwo)
	return sentinel.Next
}

func mergeLists(previous, a, b *LinkedList) {
	if a == nil && b == nil {
		return
	}

	if a == nil {
		previous.Next = b
		mergeLists(previous.Next, a, b.Next)
		return
	}

	if b == nil {
		previous.Next = a
		mergeLists(previous.Next, a.Next, b)
		return
	}

	if a.Value < b.Value {
		previous.Next = a
		mergeLists(previous.Next, a.Next, b)
		return
	}

	previous.Next = b
	mergeLists(previous.Next, a, b.Next)
}

func minNode(a, b *LinkedList) (*LinkedList, *LinkedList, *LinkedList) {
	if a == nil && b == nil {
		return nil, nil, nil
	}
	if a == nil {
		return nil, b.Next, b
	}
	if b == nil {
		return a.Next, nil, a
	}

	if a.Value < b.Value {
		return a.Next, b, a
	}

	return a, b.Next, b
}
