package linkedlists

// FindDuplicate ...
//
// T -> O(n)
// S -> O(1)
func FindDuplicate(array []int) int {
	var slow, fast = 0, 0
	for {
		slow = array[slow]
		fast = array[array[fast]]

		if slow == fast {
			break
		}
	}

	var slow2 = 0
	for {
		slow = array[slow]
		slow2 = array[slow2]

		if slow == slow2 {
			break
		}
	}

	return slow
}
