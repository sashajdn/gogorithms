package arrays

import (
	"testing"
)

func TestMoveElementToEnd(t *testing.T) {
	input := []int{2, 1, 2, 2, 2, 3, 4, 2}
	toMove := 2

	expected := []int{1, 3, 4, 2, 2, 2, 2, 2}
	result := MoveElementToEnd(input, toMove)

	if !arrayEqual(expected, result) {
		t.Fatalf("expected -> %v, got -> %v", expected, result)
	}
}
