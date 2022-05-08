package graphs

import "math"

// WallsAndGates ...
//
// T -> O(n * m) where `n` is the number of columns & `m` is the number of rows.
// S -> O(n * m) since we at most can insert `n * m` elements into the queue; for instance when every single
//               element is a gate.
func WallsAndGates(rooms [][]int) {
	var directions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var queue [][]int
	for j := 0; j < len(rooms); j++ {
		for i := 0; i < len(rooms[0]); i++ {
			if rooms[j][i] == 0 {
				queue = append(queue, []int{j, i})
			}
		}
	}

	var currentLevel int
	for len(queue) > 0 {
		l := len(queue)

		for i := 0; i < l; i++ {
			var next []int
			next, queue = queue[0], queue[1:]

			j, i := next[0], next[1]

			if rooms[j][i] < 0 || rooms[j][i] == math.MaxInt {
				continue
			}

			for _, direction := range directions {
				dj, di := j+direction[0], i+direction[1]

				if dj < 0 || dj > len(rooms)-1 {
					continue
				}
				if di < 0 || di > len(rooms[0])-1 {
					continue
				}

				if rooms[dj][di] < math.MaxInt {
					continue
				}

				rooms[dj][di] = currentLevel + 1
				queue = append(queue, []int{dj, di})
			}
		}

		currentLevel++
	}
}
