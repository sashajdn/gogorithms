package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mergeKLinkedListsFs = []func(lists []*LinkedList) *LinkedList{
	MergeKLists_DivideAndConquer,
	MergeKLists_Heap,
}

func TestMergeKLinkedList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                    string
		adjacentInput           [][]int
		expectedAdjancentOutput []int
	}{
		{
			name: "example_even",
			adjacentInput: [][]int{
				{1, 3, 5, 6, 8, 14},
				{2, 4, 9},
				{7, 10, 15},
				{11, 12, 13, 16},
			},
			expectedAdjancentOutput: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
		{
			name: "example_odd",
			adjacentInput: [][]int{
				{1, 3, 5},
				{2, 4, 9},
				{7, 10, 15},
				{11, 12, 13, 16},
				{6, 8, 14},
			},
			expectedAdjancentOutput: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, f := range mergeKLinkedListsFs {
				input := arraysToLinkedLists(tt.adjacentInput)
				output := f(input)

				adjacentOutput := linkedListToArray(output)

				assert.Equal(t, tt.expectedAdjancentOutput, adjacentOutput)
			}
		})
	}
}

func arrayToLinkedList(array []int) *LinkedList {
	var (
		sent    = &LinkedList{}
		current = sent
	)

	for _, a := range array {
		current.Next = &LinkedList{
			Value: a,
		}
		current = current.Next
	}

	return sent.Next
}

func arraysToLinkedLists(arrays [][]int) []*LinkedList {
	var output = make([]*LinkedList, 0, len(arrays))
	for _, arr := range arrays {
		ll := arrayToLinkedList(arr)

		output = append(output, ll)
	}

	return output
}
