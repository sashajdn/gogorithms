package stacks

import (
	"testing"
)

func TestBalancedBrackets(t *testing.T) {
	s := "((("
	expected := false
	res := BalancedBrackets(s)
	if res != expected {
		// t.Fatalf("expected: %v, got %v", expected, res)
	}

	s = "[[({})]]{}([])"
	expected = true
	res = BalancedBrackets(s)
	if res != expected {
		t.Fatalf("expected: %v, got %v", expected, res)
	}

	s = "hello (there)! How {do[we(do)]}"
	res = BalancedBrackets(s)
	if res != expected {
		t.Fatalf("expected: %v, got %v", expected, res)
	}
}
