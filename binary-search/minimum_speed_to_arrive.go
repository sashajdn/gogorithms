package binarysearch

import "math"

// MinSpeedOnTime ...
//
// T -> O(n log n)
// S -> O(1)
func MinSpeedOnTime(distances []int, hour float64) int {
	if len(distances) > int(math.Ceil(hour)) {
		return -1
	}

	var left, right = 0, math.MaxInt
	for left < right {
		mid := (left + right) / 2

		if isSpeedFeasible(distances, mid, hour) {
			right = mid
			continue
		}

		left = mid + 1
	}

	return left
}

func isSpeedFeasible(distances []int, speed int, hour float64) bool {
	switch len(distances) {
	case 0:
		return false
	case 1:
		return (float64(distances[0]) / float64(speed)) <= hour
	}

	var hoursTaken float64
	for _, d := range distances[:len(distances)-1] {
		hoursTaken += math.Ceil(float64(d) / float64(speed))
	}

	hoursTaken += (float64(distances[len(distances)-1]) / float64(speed))

	if hoursTaken <= hour {
		return true
	}

	return false
}
