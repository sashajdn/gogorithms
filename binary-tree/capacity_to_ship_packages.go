package binarytree

// ShipWithinDays ...
//
// T -> O(nlog(n))
// S -> O(1)
func ShipWithinDays(weights []int, days int) int {
	var l, r = 0, 0
	for _, weight := range weights {
		l = max(l, weight)
		r += weight
	}

	for l < r {
		mid, need, current := (l+r)/2, 1, 0

		for _, weight := range weights {
			if current+weight > mid {
				need++
				current = 0
			}

			current += weight
		}

		if need > days {
			l = mid + 1
			continue
		}

		r = mid

	}
	return l
}
