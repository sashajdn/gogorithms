package queues

import "testing"

func assert(t *testing.T, expected interface{}, got interface{}) {
	if expected != got {
		t.Fatalf("expected -> %v, got -> %v", expected, got)
	}
}

func assertSlice(t *testing.T, expected []interface{}, got []interface{}, equalFunc func(a, b []interface{}) bool) {
	if !equalFunc(expected, got) {
		t.Fatalf("expected -> %v, got -> %v", expected, got)
	}
}

func slicesEqual(a, b []interface{}) bool {
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
