package sets

type DSU interface {
	BuildFromArray(array []int)
	UnionSets(a, b int)
	FindRepresentative(value int) int
}

type IntDSU struct {
	representatives []int
	ranks           []int
}

// NewIntDSU ...
//
// T -> O(n) where n is the size of the array.
// S -> O(n)
func NewIntDSU(array []int) *IntDSU {
	var (
		r = make([]int, 0, len(array))
		s = make([]int, 0, len(array))
	)

	for i := 0; i < len(array); i++ {
		r = append(r, i)
		s = append(s, 1)
	}

	return &IntDSU{
		representatives: r,
		ranks:           s,
	}
}

// Find ...
//
// T -> Best: O(1) due to path compression, Avg: O(1), Worst: O(n ** 2)
// S -> O(1)
func (i *IntDSU) Find(value int) int {
	if value == i.representatives[value] {
		return value // we have already performed the union.
	}

	// Path compression.
	// [0, 0, 3, 0] -> value 2
	// rep(2) -> 3

	i.representatives[value] = i.Find(i.representatives[value])
	return i.representatives[value]
}

// Union ...
//
// Same as `Find`
func (i *IntDSU) Union(a, b int) {
	ra, rb := i.Find(a), i.Find(b)
	if ra == rb {
		// We've already unioned.
		return
	}

	switch {
	case i.ranks[ra] >= i.ranks[rb]:
		i.representatives[rb] = i.representatives[ra]
		i.ranks[ra] += i.ranks[rb]
	default:
		i.representatives[ra] = i.representatives[rb]
		i.ranks[rb] += i.ranks[ra]
	}
}
