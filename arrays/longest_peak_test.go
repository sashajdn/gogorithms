package arrays

import "testing"

func TestLongestPeak(t *testing.T) {
	array := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	expected := 6
	result := LongestPeak(array)

	if expected != result {
		t.Fatalf("expected -> %d, got -> %d", expected, result)
	}

	array = []int{1, 2, 4}
	expected = 0
	result = LongestPeak(array)

	if expected != result {
		t.Fatalf("expected -> %d, got -> %d", expected, result)
	}

	array = []int{-1, -2, -4, -3}
	expected = 0
	result = LongestPeak(array)

	if expected != result {
		t.Fatalf("expected -> %d, got -> %d", expected, result)
	}

	array = []int{1, 2, 3, 3, 2, 1}
	expected = 0
	result = LongestPeak(array)

	if expected != result {
		t.Fatalf("expected -> %d, got -> %d", expected, result)
	}

	array = []int{1, 2, 3, 2, 1, 1}
	expected = 5
	result = LongestPeak(array)

	if expected != result {
		t.Fatalf("expected -> %d, got -> %d", expected, result)
	}

	array = []int{0, 1}
	expected = 0
	result = LongestPeak(array)

	if expected != result {
		t.Fatalf("expected -> %d, got -> %d", expected, result)
	}
}
