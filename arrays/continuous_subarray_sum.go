package arrays

func ContinuousSubArraySum(array []int, k int) bool {
	var sum, previous int
	var set = map[int]struct{}{}

	for i := 0; i < len(array); i++ {
		switch k {
		case 0:
			sum += array[i]
		default:
			sum = (sum + array[i]) % k
		}

		if _, ok := set[sum]; ok {
			return true
		}

		set[previous] = struct{}{}
		previous = sum
	}

	return false
}
