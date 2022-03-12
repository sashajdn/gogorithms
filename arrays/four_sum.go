package arrays

// FourSum ...
//
// T -> Average case: O(n ** 2), Worst Case: O(n ** 3), where `n` is the number of items in the array.
//                    One the average case, where we don't have n > 1 pairs that have the same some,
//                    In the worst case, we could have n/2 pairs, giving us O(n ** 3).
// S -> O(n ** 2), since we need to store n**2 pairs in the hashmap.
func FourSum(array []int, target int) [][]int {
	var (
		sumPairs = map[int][][]int{}
		quads    = [][]int{}
	)
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			diff := target - (array[j] + array[i])

			if pairs, ok := sumPairs[diff]; ok {
				for _, pair := range pairs {
					quads = append(quads, append(pair, array[i], array[j]))
				}
			}
		}
		for k := 0; k < i; k++ {
			sum := array[k] + array[i]
			sumPairs[sum] = append(sumPairs[sum], []int{array[i], array[k]})
		}
	}
	return quads
}
