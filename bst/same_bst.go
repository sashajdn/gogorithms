package bst

import "math"

// SameBsts ...
// T -> O(n**2)
// S -> O(n**2)
func SameBsts(arrayOne, arrayTwo []int) bool {
	if len(arrayOne) != len(arrayTwo) {
		return false
	}

	if len(arrayOne) == 0 {
		return true
	}

	if arrayOne[0] != arrayTwo[0] {
		return false
	}

	lstOne, rstOne := buildLeftSubtree(arrayOne), buildRightSubtree(arrayOne)
	lstTwo, rstTwo := buildLeftSubtree(arrayTwo), buildRightSubtree(arrayTwo)

	return SameBsts(lstOne, lstTwo) && SameBsts(rstOne, rstTwo)
}

func buildLeftSubtree(array []int) []int {
	var lst []int
	h, t := array[0], array[1:]
	for _, v := range t {
		if v < h {
			lst = append(lst, v)
		}
	}

	return lst
}

func buildRightSubtree(array []int) []int {
	var rst []int
	h, t := array[0], array[1:]
	for _, v := range t {
		if v >= h {
			rst = append(rst, v)
		}
	}
	return rst
}

// SameBsts_Better ...
//
// T -> O(n ** 2).
// S -> O(d)
func SameBsts_Better(arrayOne, arrayTwo []int) bool {
	return sameBSTHelper(arrayOne, arrayTwo, 0, 0, math.MaxInt, math.MinInt)
}

func sameBSTHelper(arrayOne, arrayTwo []int, indexOne, indexTwo, minBound, maxBound int) bool {
	if indexOne == -1 || indexTwo == -1 {
		return indexOne == indexTwo
	}

	if arrayOne[indexOne] != arrayTwo[indexTwo] {
		return false
	}

	leftIndexOne := getIdxOfFirstSmaller(arrayOne, indexOne, minBound)
	leftIndexTwo := getIdxOfFirstSmaller(arrayTwo, indexTwo, minBound)
	rightIndexOne := getIdxOfFirstSmaller(arrayOne, indexOne, maxBound)
	rightIndexTwo := getIdxOfFirstSmaller(arrayTwo, indexTwo, maxBound)

	currentValue := arrayOne[indexOne]

	return sameBSTHelper(arrayOne, arrayTwo, leftIndexOne, leftIndexTwo, minBound, currentValue) && sameBSTHelper(arrayOne, arrayTwo, rightIndexOne, rightIndexTwo, currentValue, maxBound)
}

func getIdxOfFirstSmaller(array []int, index, bound int) int {
	for i := index; i < len(array); i++ {
		if array[i] < array[index] && array[i] >= bound {
			return i
		}
	}
	return -1
}

func getIdxForFirstGreater(array []int, index, bound int) int {
	for j := index; j < len(array); j++ {
		if array[j] >= array[index] && array[j] < bound {
			return j
		}
	}
	return -1
}
