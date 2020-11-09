package arrays

import "testing"

func TestSpiralTraversal(t *testing.T) {

	array := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	result := SpiralTraversal(array)

	if !arrayEqual(expected, result) {
		t.Fatalf("expected -> %v, got %v", expected, result)
	}

	array = [][]int{
		{1, 2, 3, 4},
		{8, 7, 6, 5},
	}

	expected = []int{1, 2, 3, 4, 5, 6, 7, 8}
	result = SpiralTraversal(array)

	if !arrayEqual(expected, result) {
		t.Fatalf("expected -> %v, got %v", expected, result)
	}
}
