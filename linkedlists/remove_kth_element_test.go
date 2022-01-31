package linkedlists

import "testing"

func TestRemoveKthNodeFromEnd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                string
		inputArray          []int
		k                   int
		expectedOutputArray []int
	}{
		{
			name:                "seq-array-k-middle",
			inputArray:          []int{1, 2, 3, 4, 5, 6, 7},
			k:                   4,
			expectedOutputArray: []int{1, 2, 4, 5, 6, 7},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			head := New(tc.inputArray[0])
			FromArray(tc.inputArray[1:], head)

			RemoveKthNodeFromEnd(head, tc.k)
			res := head.ToArray()

			if !func(a, b []int) bool {
				if len(a) != len(b) {
					return false
				}
				for i, v := range a {
					switch {
					case v != b[i]:
						return false
					}
				}
				return true
			}(tc.expectedOutputArray, res) {
				t.Fatalf("expected -> %v, got -> %v", tc.expectedOutputArray, res)
			}
		})
	}
}
