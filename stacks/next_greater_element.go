package stacks

// NextGreaterElement ...
//
// T -> O(n)
// S -> O(n)
func NextGreaterElement(array []int) []int {
	// Initialization.
	var output = make([]int, 0, len(array))
	for i := 0; i < len(array); i++ {
		output = append(output, -1)
	}

	var s stack
	for k := (2 * len(array)) - 1; k >= 0; k-- {
		i := k % len(array)

		for len(s) > 0 {
			if s.Peek() < array[i] {
				s.Pop()
				continue
			}

			output[i] = s.Peek()
			break
		}

		s.Add(array[i])
	}

	return output
}

type stack []int

func (s stack) Add(v int) {
	s = append(s, v)
}

func (s stack) Pop() int {
	v := s[len(s)-1]
	s = s[:len(s)-1]
	return v
}

func (s stack) Peek() int {
	return s[len(s)-1]
}
