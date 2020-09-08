package stacks

import "fmt"

type MinMaxStack struct {
	min []int
	max []int
	stack []int
}

func New() *MinMaxStack {
	return &MinMaxStack{}
}

func (m *MinMaxStack) Pop() (int, error) {
	v, err := m.Peek()

	if err != nil {
		return 0,  err
	}

	m.min = m.min[:m.length() - 1]
	m.max = m.max[:m.length() - 1]
	m.stack = m.stack[:m.length() - 1]
	return v, nil
}

func (m *MinMaxStack) Peek() (int, error) {
	if len(m.stack) == 0 {
		return 0, fmt.Errorf("stack empty")
	}
 	return m.stack[m.length() - 1], nil
}

func (m *MinMaxStack) Push(value int) {
	if len(m.stack) == 0 {
		m.min = append(m.min, value)
		m.max = append(m.max, value)
	} else {
		m.min = append(m.min, Min(value, m.min[m.length() - 1]))
		m.max = append(m.max, Max(value, m.max[m.length() - 1]))
	}
	m.stack = append(m.stack, value)
}

func (m *MinMaxStack) GetMin() (int, error) {
	if len(m.stack) == 0 {
		return 0, fmt.Errorf("stack empty")
	}
	return m.min[m.length() - 1], nil
}

func (m *MinMaxStack) GetMax() (int, error) {
	if len(m.stack) == 0 {
		return 0, fmt.Errorf("stack empty")
	}
	return m.max[m.length() - 1], nil
}

func (m *MinMaxStack) length() int {
	return len(m.stack)
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
