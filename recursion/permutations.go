package recursion

// GetPermutations ...
//
// T -> O(n*n!)
// S -> O(n*n!)
func GetPermutations(array []int) [][]int {
	switch len(array) {
	case 0:
		return [][]int{}
	case 1:
		return [][]int{array}
	case 2:
		return [][]int{
			{array[0], array[1]},
			{array[1], array[0]},
		}
	}

	var perms [][]int
	for i := 0; i < len(array); i++ {
		h := array[i]

		var t []int
		switch {
		case i == 0:
			t = append(t, array[i+1:]...)
		case i == len(array)-1:
			t = append(t, array[:i]...)
		default:
			t = append(t, array[:i]...)
			t = append(t, array[i+1:]...)
		}

		for _, p := range GetPermutations(t) {
			perms = append(perms, append(p, h))
		}
	}
	return perms
}

// GetPermutations_Other ...
//
// T -> O(n!)
// S -> O(n)
func GetPermutations_Other(array []int) [][]int {
	var permutations [][]int
	switch len(array) {
	case 0, 1:
		return permutations
	case 2:
		permutations = append(permutations, []int{array[0], array[1]})
		permutations = append(permutations, []int{array[1], array[0]})
		return permutations
	}

	for i := 0; i < len(array); i++ {
		swap(array, 0, i)
		head, tail := array[0], array[1:]

		tailPermutations := GetPermutations_Other(tail)
		for _, permutation := range tailPermutations {
			permutations = append(permutations, append(permutation, head))
		}
	}

	swap(array, 0, len(array)-1)
	return permutations
}

// GetPermutations_Better ...
func GetPermutations_Better(array []int) [][]int {
	permutations := [][]int{}
	permutateHelper(0, array, &permutations)
	return permutations
}

func permutateHelper(i int, array []int, permutations *[][]int) {
	if i == len(array)-1 {
		newPermutation := make([]int, len(array))
		copy(newPermutation, array)
		*permutations = append(*permutations, newPermutation)
	}

	for j := i; j < len(array); j++ {
		swap(array, i, j) // swap before permutating.
		permutateHelper(i+1, array, permutations)
		swap(array, i, j) // swap back.
	}

}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
