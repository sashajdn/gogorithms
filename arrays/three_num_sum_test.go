package arrays

import (
	"testing"
)

func TestThreeNumberSum(t *testing.T) {
	input := []int{12, 3, 1, 2, -6, 5, -8, 6}
	target := 0
	expected := [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}}

	result := ThreeNumberSum(input, target)

	if !arraysEqual(result, expected) {
		t.Fatalf("expected -> %d, got -> %d", expected, result)
	}
}

func arraysEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, suba := range a {
		if !arrayEqual(suba, b[i]) {
			return false
		}
	}

	return true
}

func arrayEqual(a, b []int) bool {
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
