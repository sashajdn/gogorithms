package arrays

func WaterArea(heights []int) int {
	if len(heights) < 3 {
		return 0
	}
	var (
		maxLeft  = heights[0]
		maxRight = heights[len(heights)-1]
		left     int
		right    = len(heights) - 1
		sum      int
	)

	for left <= right {
		if heights[left] < maxRight {
			if heights[left] > maxLeft {
				maxLeft = heights[left]
				left++
				continue
			}

			sum += (maxLeft - heights[left])
			left++
			continue
		}

		if heights[right] > maxRight {
			maxRight = heights[right]
			right--
			continue
		}

		sum += (maxRight - heights[right])
		right--
	}

	return sum
}
