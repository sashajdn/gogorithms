package dynamic

// JumpGameIII ...
//
// T -> O(n) where `n` is the number of items in the array.
// S -> O(n) since we are recursing with DFS, in the worst case we can have `n` recursive calls on the stack.
func JumpGameIII(arr []int, start int) bool {
	switch arr[start] {
	case 0:
		return true
	case -1:
		return false
	}

	var currentJump = arr[start]
	if start-currentJump < 0 && start+currentJump >= len(arr) {
		return false
	}

	if start-currentJump < 0 {
		return JumpGameIII(arr, start+currentJump)
	}

	if start+currentJump >= len(arr) {
		return JumpGameIII(arr, start-currentJump)
	}

	return JumpGameIII(arr, start-currentJump) || JumpGameIII(arr, start+currentJump)
}
