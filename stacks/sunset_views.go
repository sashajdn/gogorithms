package stacks

import (
	"math"
	"strings"
)

func SunsetViews(buildings []int, direction string) []int {
	var (
		buildingsWithViews = []int{}
		maxSoFar           = math.MinInt
	)
	switch strings.ToUpper(direction) {
	case "WEST":
		for i := 0; i < len(buildings); i++ {
			if buildings[i] > maxSoFar {
				maxSoFar = buildings[i]
				buildingsWithViews = append(buildingsWithViews, i)
			}
		}
	case "EAST":
		var stack []int
		for i := len(buildings) - 1; i >= 0; i-- {
			if buildings[i] > maxSoFar {
				maxSoFar = buildings[i]
				stack = append(stack, i)
			}
		}

		for k := len(stack) - 1; k >= 0; k-- {
			buildingsWithViews = append(buildingsWithViews, stack[k])
		}
	}

	return buildingsWithViews
}
