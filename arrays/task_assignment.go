package arrays

import (
	"fmt"
	"sort"
)

func TaskAssignment(k int, tasks []int) [][]int {
	var assigments = make([][]int, k)

	var indexArray = make([]int, 0, len(tasks))
	for i := 0; i < len(tasks); i++ {
		indexArray = append(indexArray, i)
	}

	fmt.Println("Tasks: ", tasks, "Before: ", indexArray)

	sort.Slice(indexArray, func(i, j int) bool {
		fmt.Println("Sort: ", tasks[i], tasks[j])
		return tasks[i] <= tasks[j]
	})

	fmt.Println("Tasks: ", tasks, "Sorted: ", indexArray)

	for i := 0; i < k; i++ {
		switch {
		case i%2 == 0:
			for j := 0; j < k; j++ {
				idx := indexArray[i*k+j]
				assigments[j] = append(assigments[j], idx)
			}
		default:
			for j := 0; j < k; j++ {
				idx := indexArray[i*k+j]
				assigments[k-j-1] = append(assigments[k-j-1], idx)
			}
		}
	}

	return assigments
}
