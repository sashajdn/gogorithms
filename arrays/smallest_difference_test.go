package arrays

import (
	"testing"
)

func TestSmallestDifference(t *testing.T) {

	a := []int{-1, 5, 10, 20, 28, 3}
	b := []int{26, 134, 135, 15, 17}

	expected := []int{28, 26}
	result := SmallestDifference(a, b)

	if !arrayEqualSM(result, expected) {
		t.Fatalf("expected -> %v, got -> %v", expected, result)
	}
}

func arrayEqualSM(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
