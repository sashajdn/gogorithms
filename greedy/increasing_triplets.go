package greedy

func increasingTriplets(array []int) bool {
	if len(array) < 3 {
		return false
	}

	var sub = []int{array[0]}
	for i := 1; i < len(array); i++ {
		if array[i] > sub[len(sub)-1] {
			sub = append(sub, array[i])
			continue
		}

		for j := 0; j < len(sub); j++ {
			if array[i] <= sub[j] {
				sub[j] = array[i]
				break
			}
		}
	}

	return len(sub) >= 3
}
