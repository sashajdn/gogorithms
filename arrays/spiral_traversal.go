package arrays

// SpiralTraversal T -> O(), S -> O()
func SpiralTraversal(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}

	r := []int{}

	i, j := 0, 0
	startI, startJ := 0, 0
	endI, endJ := len(array), len(array[0])

	for {
		if startI == endI || startJ == endJ {
			break
		}

		for {
			if j < endJ-1 {
				r = append(r, array[i][j])
				j++

			} else {
				break
			}
		}

		for {
			if i < endI-1 {
				r = append(r, array[i][j])
				i++

			} else {
				break
			}
		}

		for {
			if j > startJ {
				r = append(r, array[i][j])
				j--
			} else {
				break
			}
		}

		for {
			if i > startI {
				r = append(r, array[i][j])
				i--
			} else {
				break
			}
		}

		startI++
		startJ++
		endJ--
		endI--

		i, j = startI, startJ
	}

	return r
}
