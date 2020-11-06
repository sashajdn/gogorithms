package recursion

import (
	"testing"
)

func TestNthFib(t *testing.T) {
	n := 0
	expected := 0
	result := NthFib(n)

	if result != expected {
		t.Fatalf("expected -> %d, result -> %d", expected, result)
	}

	n = 1
	expected = 1
	result = NthFib(n)

	if result != expected {
		t.Fatalf("expected -> %d, result -> %d", expected, result)
	}

	n = 2
	expected = 1
	result = NthFib(n)

	if result != expected {
		t.Fatalf("expected -> %d, result -> %d", expected, result)
	}

	n = 10
	expected = 55
	result = NthFib(n)

	if result != expected {
		t.Fatalf("expected -> %d, result -> %d", expected, result)
	}

	n = 12
	expected = 144
	result = NthFib(n)

	if result != expected {
		t.Fatalf("expected -> %d, result -> %d", expected, result)
	}

	n = 25
	expected = 75025
	result = NthFib(n)

	if result != expected {
		t.Fatalf("expected -> %d, result -> %d", expected, result)
	}
}
