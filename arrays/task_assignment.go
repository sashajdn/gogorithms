package arrays

import (
	"sort"
)

// TaskAssignment given an integer `k` representing a number of workers & an array
// of positive integers representing duration of tasks that must be completed by the workers.
// And that each worker must complete two unique tasks & can only work sequentially
// And that workers can work in parallel, returns the optimal set of tasks per worker to complete,
// such that the tasks are completed as fast as possible.
//
// T -> O(nlogn) where n is the number of tasks, since nlogn >> k
// S -> O(n) where n is the number of tasks.
func TaskAssignment(k int, tasks []int) [][]int {
	var assigments = make([][]int, k)

	var hm = make(map[int][]int)
	for i, v := range tasks {
		if _, ok := hm[v]; ok {
			hm[v] = append(hm[v], i)
			continue
		}

		hm[v] = []int{i}
	}

	// T -> O(nlogn)
	sort.Ints(tasks)

	// T -> O(k)
	for pass := 0; pass < len(tasks)/k; pass++ {
		for i := 0; i < k; i++ {

			var taskValue int
			switch {
			case pass%2 == 0:
				taskValue = tasks[(k*pass)+i]
			default:
				taskValue = tasks[(k*pass)+(k-i-1)]
			}

			bucket := hm[taskValue]

			var taskIdx int
			switch len(bucket) {
			case 1:
				taskIdx = bucket[0]
			default:
				taskIdx = bucket[0]
				hm[taskValue] = hm[taskValue][1:]
			}

			assigments[i] = append(assigments[i], taskIdx)
		}
	}

	return assigments
}
