package stacks

// Stack : struct
type Stack struct {
	stack []interface{}
}

// NewStack : Factory method
func NewStack() *Stack {
	return &Stack{}
}

// Pop : pops the last value off the stack
func (s *Stack) Pop() interface{} {
	v := s.Peek()
	s.stack = s.stack[:len(s.stack) - 1]
	return v
}

// Peek : peeks at the last value in the stack
func (s *Stack) Peek() interface{} {
	return s.stack[len(s.stack) - 1]
}

// Push : pushes a new ``value`` onto the stack
func (s *Stack) Push(values ...interface{}) {
	s.stack = append(s.stack, values...)
}

// BalancedBrackets : checks if string contains balanced brackets
func BalancedBrackets(s string) bool {
	stack := NewStack()
	bh := newBracketHandler()

	for _, char := range s{
		if bh.isOpenBracket(char) {
			stack.Push(char)
		}

		if bh.isClosedBracket(char) {
			expected := stack.Pop()
			if expected != bh.closed2Open[char] {
				return false
			}	
		}
	}
	return true
}

type bracketHandler struct {
	closed2Open map[rune]rune
}

func newBracketHandler() *bracketHandler{
	closed2Open := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	return &bracketHandler{closed2Open}
}

func  (b *bracketHandler) isOpenBracket(r rune) bool{
	for _, c := range b.closed2Open {
		if c == r {
			return true
		}
	}
	return false
}

func (b *bracketHandler) isClosedBracket(r rune) bool {
	for o := range b.closed2Open {
		if o == r {
			return true
		}
	}
	return false
}
