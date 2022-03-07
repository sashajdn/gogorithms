package recursion

// Powerset ...
//
// T -> O(n * 2 ** n)
// S -> O(n * n**2)
func Powerset(array []int) [][]int {
	powerset := [][]int{{}}

	for _, v := range array {
		l := len(powerset)
		for i := 0; i < l; i++ {
			newSubset := append([]int{}, powerset[i]...)
			newSubset = append(newSubset, v)
			powerset = append(powerset, newSubset)
		}
	}

	return powerset
}

// Powerset_Recursive ...
// T -> O(n * 2 ** n)
// S -> O(k)
func Powerset_Recursive(array []int) [][]int {
	var output = [][]int{{}}
	powerset(array, &output)
	return output
}

func powerset(array []int, collector *[][]int) {
	if len(array) == 0 {
		return
	}

	h, t := array[0], array[1:]
	l := len(*collector)

	for i := 0; i < l; i++ {
		*collector = append(*collector, append((*collector)[i], h))
	}

	powerset(t, collector)
}
