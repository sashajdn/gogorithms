package arrays

import (
	"fmt"
	"sort"
)

// ThreeNumberSort ...
// T -> O(nlogn)
// S -> O(1)
func ThreeNumberSort(array []int, order []int) []int {
	ax := indexSort{
		a:     &array,
		order: &order,
	}

	fmt.Println(array, *ax.a)

	sort.Sort(ax)

	return *ax.a
}

func indexFromValue(array []int, value int) int {
	for i, a := range array {
		if a == value {
			return i
		}
	}

	// We shouldn't never hit this point.
	return 0
}

type indexSort struct {
	a     *[]int
	order *[]int
}

func (i indexSort) Len() int { return len(*i.a) }

func (i indexSort) Swap(n, m int) {
	(*i.a)[n], (*i.a)[m] = (*i.a)[m], (*i.a)[n]
}

func (i indexSort) Less(n, m int) bool {
	x, y := indexFromValue(*i.order, (*i.a)[n]), indexFromValue(*i.order, (*i.a)[m])

	return x < y
}

// ThreeNumberSort_Linear ...
//
// T -> O(n)
// S -> O(1)
func ThreeNumberSort_Linear(array []int, order []int) []int {
	var x, y, z int

	for _, v := range array {
		switch v {
		case order[0]:
			x++
		case order[1]:
			y++
		case order[2]:
			z++
		}
	}

	for i := 0; i < x; i++ {
		array[i] = order[0]
	}

	for i := x; i < x+y; i++ {
		array[i] = order[1]
	}

	for i := x + y; i < x+y+z; i++ {
		array[i] = order[2]
	}

	return array
}
