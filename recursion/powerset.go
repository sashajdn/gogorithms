package recursion

// Powerset : T -> O(n * 2^n), S -> O(n * 2^n)
func Powerset(array []int) [][]int {
	ps := make([]int, 0)
	ps = append(ps, make([]int, 0))
	for _, ele := range array {
		l := len(ps)
		for j := 0; j < l; j++ {
			ps = append(ps, append(append(make([]int, 0), ps[j]...), ele))
		}
	}
	return ps
}
