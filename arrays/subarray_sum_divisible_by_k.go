package arrays

// SubArrayDivisibleByK ...
//
// T -> O(n) where `n` is the number of items in the array; since we iterate through each of the items once, we have a time complexity
//           of order `n`.
// S -> O(min(k, n)) where `n` is the number of items in the array; here we create a hashmap of remainders we will be at most size `n`.
func SubArrayDivisibleByK(array []int, k int) int {
	if len(array) == 0 {
		return 0
	}
	if k == 0 {
		return 0
	}

	var (
		cumSum, count int
		frequencies   = make([]int, k)
	)

	frequencies[0] = 1

	for _, v := range array {
		cumSum = (v + cumSum)
		remainder := cumSum % k

		if remainder < 0 {
			remainder += k
		}

		count += frequencies[remainder]
		frequencies[remainder]++
	}

	return count
}
