package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesFromLinkedList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		nodes        []*LinkedList
		expectedList []int
	}{
		{
			name: "with-duplicates",
			nodes: []*LinkedList{
				{
					Value: 1,
				},
				{
					Value: 1,
				},
				{
					Value: 1,
				},
				{
					Value: 4,
				},
			},
			expectedList: []int{1, 4},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Build linkedlist from a list of nodes.
			root := tt.nodes[0]
			buildLinkedList(root, tt.nodes[1:])
			// Remove duplicates.
			head := RemoveDuplicatesFromLinkedList(root)
			// Construct array to test against.
			a := constructArray(root, []int{})

			assert.Equal(t, root, head)
			assert.Equal(t, tt.expectedList, a)
		})
	}
}

func buildLinkedList(head *LinkedList, nodes []*LinkedList) {
	if len(nodes) == 0 {
		return
	}
	head.Next = nodes[0]
	buildLinkedList(head.Next, nodes[1:])
}

func constructArray(head *LinkedList, a []int) []int {
	if head == nil {
		return a
	}
	a = append(a, head.Value)
	return constructArray(head.Next, a)
}
