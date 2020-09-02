package sorting

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	array := []int{5, 4, 3, 7, 2, 1, 6}
	sortedArray := MergeSort(array)

	for idx, val := range sortedArray{
		if val != idx + 1 {
			t.Fatal()
		}
	}
}

func TestMerge(t *testing.T) {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 4, 6, 8}
	arr := make([]int, 0)

	sortedToTest := merge(arr, a, b)

	for idx, val := range sortedToTest {
		if val != idx + 1 {
			t.Fatal()
		}
	}
}
