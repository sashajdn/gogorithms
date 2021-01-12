package graphs

import "testing"

func TestBreadFirstSearch(t *testing.T) {
	t.Parallel()

	root := New("A")

	childA := New("B")
	childB := New("C")

	root.AddChild(childA)
	root.AddChild(childB)

	childAA := New("D")

	childA.AddChild(childAA)

	childBA := New("E")
	childBA.AddChild(childBA)

	childAAA := New("F")
	childAAB := New("G")

	childAA.AddChild(childAAA)
	childAA.AddChild(childAAB)

	expectedOutput := []string{"A", "B", "C", "D", "E", "F", "G"}
	res := root.BreadthFirstSearch([]string{})

	if !stringSlicesEqual(res, expectedOutput) {
		t.Fatalf("expected -> %v, got -> %v", expectedOutput, res)
	}
}

func stringSlicesEqual(a, b []string) bool {
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
