package arrays

import "testing"

func TestIsMonotonicArray(t *testing.T) {

	input := []int{-5, -6, -7, -8}
	expected := true
	result := IsMonotonicArray(input)

	if expected != result {
		t.Fatal()
	}

	input = []int{1, 4, 6, 7, 9}
	expected = true
	result = IsMonotonicArray(input)

	if expected != result {
		t.Fatal()
	}

	input = []int{-5, 5, -6, 10, -11}
	expected = false
	result = IsMonotonicArray(input)

	if expected != result {
		t.Fatal()
	}

	input = []int{0}
	expected = true
	result = IsMonotonicArray(input)

	if expected != result {
		t.Fatal()
	}

	input = []int{-1, -5, -10, -1100, -900, -1101, -1102, -9001}
	expected = false
	result = IsMonotonicArray(input)

	if expected != result {
		t.Fatal()
	}
}
