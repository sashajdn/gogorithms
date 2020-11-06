package recursion

// NthFib T -> O(n), S -> O(1)
func NthFib(n int) int {

	cache := []int{0, 1}

	counter := 2

	for counter <= n {
		nextFib := cache[0] + cache[1]
		cache = []int{cache[1], nextFib}
		counter++
	}

	if n > 0 {
		return cache[1]
	}

	return cache[0]
}
