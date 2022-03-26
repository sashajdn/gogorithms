package dynamic

// WaterArea ...
//
// T -> O(n) where `n` is the number of elements in the array.
// S -> O(1)
func WaterArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	var (
		area              int
		leftIdx           = 0
		rightIdx          = len(heights) - 1
		leftMax, rightMax = heights[0], heights[len(heights)-1]
	)

	for leftIdx < rightIdx {
		if heights[leftIdx] < heights[rightIdx] {
			leftIdx++
			leftMax = max(leftMax, heights[leftIdx])
			area += leftMax - heights[leftIdx]
			continue
		}

		rightIdx--
		rightMax = max(rightMax, heights[rightIdx])
		area += rightMax - heights[rightIdx]
	}

	return area
}
