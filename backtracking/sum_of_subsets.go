package backtracking

func SumOfSubsets(array []int, target int) [][]int {
	var totalSum int
	for _, v := range array {
		totalSum += v
	}

	var output [][]int
	sumOfSubsets(array, 0, target, 0, totalSum, []int{}, &output)
	return output
}

func sumOfSubsets(array []int, index, target, sum, leaving int, collection []int, output *[][]int) {
	if sum == target && len(array) > 1 && len(collection) > 0 && array[index-1] == collection[len(collection)-1] {
		*output = append(*output, collection)
	}

	if index >= len(array) {
		return
	}

	sumOfSubsets(array, index+1, target, sum, leaving, collection, output)

	collection = append(collection, array[index])
	sumOfSubsets(array, index+1, target, sum+array[index], leaving-array[index], collection, output)
}
