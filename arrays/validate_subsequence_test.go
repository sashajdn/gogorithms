package arrays

import (
	"testing"
)

func TestValidateSubsequence(t *testing.T) {
	arr := []int{1, 3}
	seq := []int{7, 5}
	expected := false

	if ValidateSubsequence(arr, seq) != expected {
		t.Fatal()
	}

	arr = []int{9, 3, 7, 5, 1, 9}
	seq = []int{7, 5, 1}
	expected = true

	if ValidateSubsequence(arr, seq) != expected {
		t.Fatal()
	}
}
