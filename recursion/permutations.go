package recursion

func GetPermutations(array []int) [][]int {
	if len(array) == 0 {
		return [][]int{}
	}
	if len(array) == 2 {
		head, tail := array[0], array[1]
		return [][]int{
			{head, tail},
			{tail, head},
		}
	}
	arr := [][]int{}
	for i := 0; i < len(array); i++ {
		head, tails := []int{array[0]}, GetPermutations(array[1:])
		for _, tail := range tails {
			permutation := append(head, tail...)
			arr = append(arr, permutation)
		}
		array = func(a []int) []int {
			array = append(array, array[0])
			array = array[1:]
			return array
		}(array)
	}
	return arr
}
