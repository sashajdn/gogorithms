package dynamic

// MinDeleteOperations ...
//
// T -> O(n * m)
// S -> O(min(n, m))
func MinDeleteOperations(a, b string) int {
	if len(a) > len(b) {
		return MinDeleteOperations(b, a)
	}

	var (
		top    = make([]int, len(a)+1)
		bottom = make([]int, len(a)+1)
	)

	for j := len(b) - 1; j >= 0; j-- {
		for i := len(a) - 1; i >= 0; i-- {
			if a[i] == b[j] {
				top[i] = max(bottom[i+1]+1, max(top[i+1], bottom[i]))
				continue
			}

			top[i] = max(bottom[i], top[i+1])
		}

		bottom = top
		top = make([]int, len(a)+1)
	}

	return len(a) + len(b) - 2*bottom[0]
}
