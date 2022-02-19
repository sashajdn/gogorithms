package arrays

import (
	"testing"
)

var twoSumCheckers = []func([]int, int) []int{
	TwoNumberSum_HashmapOnePass,
	TwoNumberSum_Sorted,
}

func TestTwoNumberSum(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}
	target := 8
	expected := []int{3, 5}

	for _, c := range twoSumCheckers {
		res := c(arr, target)

		if !equals(res, expected) {
			t.Fatal()
		}
	}

}

func equals(a, b []int) bool {
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
