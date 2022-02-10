package arrays

func findPairs(nums []int, k int) int {
	var hm = make(map[int]struct{}, 0)

	for _, n := range nums {
		if _, ok := hm[n-k]; ok {
			continue
		}
		hm[n-k] = struct{}{}
	}

	var count int
	for _, n := range nums {
		if _, ok := hm[n]; ok {
			count++
			delete(hm, n)
		}
	}

	return count
}
