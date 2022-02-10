package recursion

// Powerset ...
//
// T -> O(n * n**2)
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
