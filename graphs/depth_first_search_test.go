package graphs

import (
	"testing"
)

func TestDepthFirstSearch(t *testing.T) {

	g6 := Graph{
		id: "f",
	}

	g5 := Graph{
		id:       "e",
		children: []*Graph{&g6},
	}

	g4 := Graph{
		id: "d",
	}

	g3 := Graph{
		id:       "c",
		children: []*Graph{&g4, &g5},
	}

	g2 := Graph{
		id: "b",
	}

	g1 := Graph{
		id:       "a",
		children: []*Graph{&g2, &g3},
	}

	expected := []string{"a", "b", "c", "d", "e", "f"}
	a := g1.DepthFirstSearch([]string{})

	if !arraysEqual(expected, a) {
		t.Fatalf("expected -> %v, got -> %v", expected, a)
	}
}

func arraysEqual(a, b []string) bool {
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
