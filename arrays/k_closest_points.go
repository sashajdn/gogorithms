package arrays

// KClosestPoints ...
//
// T -> Best: O(n), Worst: O(n ** 2), Avg: O(n)
// S -> O(1)
func KClosestPoints(points [][]int, k int) [][]int {
	if k >= len(points) {
		return points
	}

	return quickSelectPoints(points, 0, len(points)-1, k)
}

func quickSelectPoints(points [][]int, start, end, target int) [][]int {
	switch len(points) {
	case 0:
		return [][]int{}
	case 1:
		return points
	}

	pivot := start
	l, r := pivot+1, end

	for l <= r {
		ep, el, er := euclideanDistanceSquared(points[pivot]), euclideanDistanceSquared(points[l]), euclideanDistanceSquared(points[r])
		if el > ep && er < ep {
			swapPoints(points, l, r)
		}
		if el <= ep {
			l++
		}
		if er >= ep {
			r--
		}
	}
	swapPoints(points, pivot, r)

	switch {
	case r == target:
		return points[:target]
	case r < target:
		return quickSelectPoints(points, r+1, end, target)
	default:
		return quickSelectPoints(points, start, r-1, target)
	}
}

func euclideanDistanceSquared(point []int) int {
	x, y := point[0], point[1]
	return (x * x) + (y * y)
}

func swapPoints(array [][]int, i, j int) {
	array[i], array[j] = array[j], array[i]
}
