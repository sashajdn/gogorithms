package stacks

import (
	"math"
)

// LargestRectangleUnderSkyline ...
//
// T -> O(n ** 2)
// S -> O(1)
func LargestRectangleUnderSkyline_BruteForce(buildings []int) int {
	var maxArea = math.MinInt
	for i, building := range buildings {
		var count int
		for j := i; j >= 0; j-- {
			if buildings[j] >= buildings[i] {
				count++
			}
		}

		for k := i + 1; k < len(buildings); k++ {
			if buildings[k] >= buildings[i] {
				count++
			}
		}

		maxArea = max(maxArea, count*building)
	}

	return maxArea
}

// LargestRectangleUnderSkyline_Stack  ...
//
// T -> O(n)
// S -> O(n)
func LargestRectangleUnderSkyline_Stack(buildings []int) int {
	var (
		stack             []int
		maxSoFar          int
		extendedBuildings = append(buildings, 0)
	)

	for i, buildingHeight := range extendedBuildings {
		for len(stack) != 0 && buildings[stack[len(stack)-1]] >= buildingHeight {
			var rightBound int
			rightBound, stack = stack[len(stack)-1], stack[:len(stack)-1]

			rightBoundHeight := buildings[rightBound]

			var width = i
			if len(stack) != 0 {
				width = i - stack[len(stack)-1] - 1
			}

			maxSoFar = max(maxSoFar, width*rightBoundHeight)
		}

		stack = append(stack, i)
	}

	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
